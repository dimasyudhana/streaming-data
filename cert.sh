# Generate the Kafka private key
openssl genpkey -algorithm RSA -out service.key -pkeyopt rsa_keygen_bits:2048

# Generate the Certificate Signing Request (CSR)
openssl req -new -key service.key -out service.csr -subj "/C=US/ST=California/L=SanFrancisco/O=YourOrg/OU=Kafka/CN=localhost"

# Create a self-signed certificate
openssl x509 -req -in service.csr -signkey service.key -out service.cert -days 365

# Generate the CA private key
openssl genpkey -algorithm RSA -out ca.key -pkeyopt rsa_keygen_bits:2048

# Generate the CA certificate
openssl req -new -x509 -key ca.key -out ca.pem -days 365 -subj "/C=US/ST=California/L=SanFrancisco/O=YourOrg/OU=KafkaCA/CN=localhost"

# Create the keystore with the server certificate
openssl pkcs12 -export -in service.cert -inkey service.key -out service.p12 -name kafka -CAfile ca.pem -caname root

# Convert the PKCS12 keystore into a Java Keystore
keytool -importkeystore -srckeystore service.p12 -srcstoretype PKCS12 -destkeystore service.keystore.jks -deststoretype JKS

# Create the truststore
keytool -import -file ca.pem -alias CARoot -keystore service.truststore.jks -storepass secret123
