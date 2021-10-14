kafka-producer-perf-test.sh --producer.config client-broker-ssl.properties --producer-props bootstrap.servers="zookeeper01:9093,zookeeper02:9093,zookeeper03:9093" --topic sample --num-records 2000 --throughput 100 --record-size 256



- Cấu hình client cho kafka
```
ssl.truststore.location=/usr/local/kafka/exchange/ssl-ca-public/public-client.truststore.jks
ssl.truststore.password=truststore-secretanh99em2400379-public
ssl.keystore.location=/usr/local/kafka/exchange/ssl-ca-public/public-client.keystore.jks
ssl.keystore.password=keystore-secretanh99em2400379-public
ssl.key.password=keystore-secretanh99em2400379-public

security.protocol=SSL
ssl.enabled.protocols=TLSv1.2
````

- Khởi tạo file client cho zookeeper
```
vi client-zookeeper-ssl.properties

zookeeper.ssl.client.enable=true
zookeeper.clientCnxnSocket=org.apache.zookeeper.ClientCnxnSocketNetty
zookeeper.ssl.keystore.location=/usr/local/server.keystore.jks
zookeeper.ssl.keystore.password=keystore-secretanh99em2400379
zookeeper.ssl.truststore.location=/usr/local/server.truststore.jks
zookeeper.ssl.truststore.password=truststore-secretanh99em2400379
zookeeper.set.acl=true

security.protocol=SSL
ssl.enabled.protocols=TLSv1.2
````


- Nếu thêm dòng này vào kafka thì sẽ tự động init super admin cho lần đầu kết nối tới kafka
```
super.users=User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNOi,C=VI
```

- Thực hiện khởi tạo một số acl cho người dùng. Cái này là liên kết giữa zk và kafka client
````


./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01.site2:2182 --add --operation ClusterAction  --force  --cluster --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNOi,C=VI 



````

- Sau đó khởi tạo một acl cho kafka khi thực hiện đồng bộ giữa các node mới nhau. Sử dụng ssl CName được gen ra. Lưu ý ssl giữa zk và brober và broker với broker sẽ khác.
```
./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01:2182 --add --operation ClusterAction --force  --cluster --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNoi,C=VI

<!-- role for mirror marker -->

./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01.site2:2182 --add --operation All --topic '*'  --cluster --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNoi,C=VI --group '*'

./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01:2182 --add --operation Describe --topic '*'  --cluster --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNoi,C=VI

./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01:2182 --add --operation DescribeConfigs --topic '*' --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNoi,C=VI

./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01:2182 --add --operation DescribeConfigs --topic '*' --allow-principal User:CN=hn01.central.kafka,OU=CXVIEW,O=org,L=HN,ST=HaNoi,C=VI

./kafka-topics.sh --command-config client-broker-ssl.properties --bootstrap-server kafka01:9093 --create --replication-factor 1 --partitions 4 --topic sample2
```



## prod

- After install kafka and zookeeper 

```


```

- Clean
```
rm -rf /usr/local/public-broker.truststore.jks /usr/local/public-broker.keystore.jks /tmp/cert-* /usr/local/kafka/exchange/ssl-ca-public/ /usr/local//kafka/exchange/ssl-client-dir-public

```



```
scp /usr/local/public-broker.truststore.jks  root@zookeeper02:/usr/local/public-broker.truststore.jks
scp /usr/local/public-broker.truststore.jks  root@zookeeper03:/usr/local/public-broker.truststore.jks
scp /usr/local/public-broker.keystore.jks   root@zookeeper03:/usr/local/public-broker.keystore.jks
scp /usr/local/public-broker.keystore.jks   root@zookeeper02:/usr/local/public-broker.keystore.jks
```



## Export  Certificate Authority (CA) certificate:

```
keytool -export -alias caroot -file caRoot.der -keystore /usr/local/public-broker.keystore.jks

openssl x509 -inform der -in caRoot.der -out caRoot.pem


```


## export client cert

```
keytool -export -alias localhost -file certificate.der -keystore /usr/local/kafka/exchange/ssl-ca-public/public-client.keystore.jks

openssl x509 -inform der -in certificate.der -out certificate.pem

```

## export client key
```
keytool -importkeystore -srckeystore /usr/local/kafka/exchange/ssl-ca-public/public-client.keystore.jks -destkeystore keystore.p12 -deststoretype PKCS12

openssl pkcs12 -in keystore.p12  -nodes -nocerts -out certificate.key


## https://serverfault.com/questions/715827/how-to-generate-key-and-crt-file-from-jks-file-for-httpd-apache-server
```