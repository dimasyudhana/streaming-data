package main

import (
	"context"
	"fmt"
	"log"
	"stream/commonlib"
)

type PaymentService struct {
	orders map[string]Order
	data   map[string]Payment
	event  *commonlib.Emitter
}

func NewPaymentService(em *commonlib.Emitter) *PaymentService {
	return &PaymentService{
		orders: make(map[string]Order, 0),
		data:   make(map[string]Payment, 0),
		event:  em,
	}
}

func (p *PaymentService) MakePayment(command NewPaymentCommand) (string, error) {
	if order, ok := p.orders[command.OrderID]; ok {
		if command.Value < order.Price {
			return "", fmt.Errorf("balance is not enough")
		}

		payment := NewPayment(order.OrderID, command.Value)
		p.data[payment.ID] = payment

		event := map[string]interface{}{
			"ID":      payment.ID,
			"OrderID": payment.OrderID,
			"Value":   float32(payment.Value),
		}

		err := p.event.Publish(context.Background(), "order_created", payment.ID, event, nil)
		if err != nil {
			return "", fmt.Errorf("failed to publish new order event: %v", err)
		}

		log.Printf("PaymentCreated. Published: %v", payment)

		return payment.ID, nil
	}
	return "", fmt.Errorf("Order not found. Payment is not created")
}

func (p *PaymentService) publishPaymentCreated(newPayment Payment) error {

	return nil
}

func (p *PaymentService) AddOrder(order Order) error {
	p.orders[order.OrderID] = order
	return nil
}

func (p PaymentService) GetPayment() ([]Payment, error) {
	payments := make([]Payment, 0)
	for _, v := range p.data {
		payments = append(payments, v)
	}
	return payments, nil
}
