
```
./kafka-acls.sh --zk-tls-config-file client-zookeeper-ssl.properties --authorizer-properties zookeeper.connect=zookeeper01:2182 --add --operation Create --topic '*' --allow-principal User:*
```