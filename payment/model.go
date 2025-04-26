package main

import "github.com/google/uuid"

type Payment struct {
	ID      string
	OrderID string
	Value   float64
}

type Order struct {
	OrderID string
	Price   float64
}

type OrderCreated struct {
	ID         string   `json:"id"`
	CustomerID string   `json:"customer_id"`
	LineItems  []string `json:"line_items"`
	Price      float64  `json:"price"`
}

type NewPaymentCommand struct {
	OrderID       string  `json:"order_id"`
	Value         float64 `json:"value"`
	CreditCardNum string  `json:"credit_card_number"`
	CreditCardCVC string  `json:"credit_card_cvc"`
}

func NewPayment(orderID string, value float64) Payment {
	return Payment{
		ID:      uuid.New().String(),
		OrderID: orderID,
		Value:   value,
	}
}
