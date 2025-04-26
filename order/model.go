package main

import "github.com/google/uuid"

type Order struct {
	ID            string   `json:"id"`
	CustomerID    string   `json:"customer_id"`
	LineItems     []string `json:"line_items"`
	Price         float64  `json:"price"`
	PaymentStatus string   `json:"payment_status"`
}

type NewOrderCommand struct {
	CustomerID string   `json:"customer_id"`
	LineItems  []string `json:"line_items"`
}

func NewOrder(customerID string, lineItems []string) Order {
	return Order{
		ID:            uuid.New().String(),
		CustomerID:    customerID,
		LineItems:     lineItems,
		Price:         20.0,
		PaymentStatus: "UNPAID",
	}
}

type OrderCreatedEvent struct {
	ID            string   `json:"id"`
	CustomerID    string   `json:"customer_id"`
	LineItems     []string `json:"line_items"`
	Price         float64  `json:"price"`
	PaymentStatus string   `json:"payment_status"`
}
