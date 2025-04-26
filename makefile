generate_kafka_cert:
	keytool -genkey -keyalg RSA -keystore service.keystore.jks -storepass changeit -validity 365 -keysize 2048
	keytool -export -alias localhost -file service.truststore.jks -keystore service.keystore.jks -storepass changeit
