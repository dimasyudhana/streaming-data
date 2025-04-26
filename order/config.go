package main

var DefaultConfig = map[string]interface{}{
	"name":       "order_service",
	"port":       50050,
	"log_format": "json",
	"log_level":  "debug",
	"tracer":     "no-op",
	"version":    "0.0.1",
	"event": map[string]interface{}{
		"sender": map[string]interface{}{
			"type": "kafka",
			"config": map[string]interface{}{
				"schema":        "kafka://",
				"kafka_brokers": "localhost:9092",
				"kafka_cert":    "service.cert",
				"kafka_key":     "service.key",
				"kafka_pem":     "ca.pem",
				"kafka_auth":    "PLAINTEXT", // local
			},
		},
		"event_config": map[string]interface{}{
			"event_map": map[string]string{
				"order_created": "order-created",
				"order_updated": "order-updated",
			},
		},
	},
}
