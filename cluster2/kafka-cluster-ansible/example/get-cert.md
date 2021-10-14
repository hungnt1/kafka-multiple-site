

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