package main

import (
	"context"
	"fmt"
	"log"

	commonlib "stream/commonlib"
)

type OrderService struct {
	data  map[string]Order
	event *commonlib.Emitter
}

func NewOrderService(em *commonlib.Emitter) *OrderService {
	return &OrderService{
		data:  make(map[string]Order),
		event: em,
	}
}

func (s *OrderService) NewOrder(command NewOrderCommand) (string, error) {
	order := NewOrder(command.CustomerID, command.LineItems)
	s.data[order.ID] = order

	event := OrderCreatedEvent{
		ID:            order.ID,
		CustomerID:    order.CustomerID,
		LineItems:     order.LineItems,
		Price:         order.Price,
		PaymentStatus: order.PaymentStatus,
	}

	err := s.event.Publish(context.Background(), "order-created", order.ID, event, nil)
	if err != nil {
		return "", fmt.Errorf("failed to publish new order event: %v", err)
	}

	log.Printf("New Order is published: %v", order)

	return order.ID, nil
}

func (s *OrderService) GetOrders() ([]Order, error) {
	orders := make([]Order, 0)
	for _, v := range s.data {
		orders = append(orders, v)
	}
	return orders, nil
}
