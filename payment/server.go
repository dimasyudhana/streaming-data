package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"stream/commonlib"
)

func Run() {
	conf, err := commonlib.Load(DefaultConfig, "")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	em, err := commonlib.New(context.Background(), conf.Get("event"))
	if err != nil {
		log.Fatal("Failed to initialize event emitter:", err)
	}

	paymentService := NewPaymentService(em)

	if _, err := commonlib.Consumer("order-created", "order-group", "PLAINTEXT", "", "", "", func(m *commonlib.Msg) error {

		var wrapper struct {
			Data json.RawMessage `json:"data"`
		}

		if err := json.Unmarshal(m.Data, &wrapper); err != nil {
			log.Printf("[ERROR] Failed to unmarshal wrapper: %v", err)
			return err
		}

		var order OrderCreated
		if err := json.Unmarshal(wrapper.Data, &order); err != nil {
			log.Printf("[ERROR] Failed to unmarshal order data: %v", err)
			return err
		}

		paymentService.AddOrder(Order{
			OrderID: order.ID,
			Price:   float64(order.Price),
		})

		return nil

	}); err != nil {
		log.Fatalf("Failed to subscribe to Kafka topic: %v", err)
	}

	log.Println("Consumer subscribed and listening for new orders...")

	paymentHandler := NewServer(paymentService)

	commonlib.ConsoleGreet(conf.GetString("name"), conf.GetString("version"), "", conf.GetInt("port"))
	log.Fatal(http.ListenAndServe(":50051", paymentHandler.Router))
}
