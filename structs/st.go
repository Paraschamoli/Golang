package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Name string
	ID   int64
}

type Order struct {
	OrderID        string
	OrderAmount    float32
	OrderStatus    string
	OrderCreatedAt time.Time
	Customer
}

// UpdateAmount updates the order's amount
func (o *Order) UpdateAmount(newAmount float32) {
	o.OrderAmount = newAmount
}

// NewOrder is a constructor for creating a new Order
func NewOrder(id string, amount float32, status string, createdAt time.Time, customerName string, customerID int64) *Order {
	return &Order{
		OrderID:        id,
		OrderAmount:    amount,
		OrderStatus:    status,
		OrderCreatedAt: createdAt,
		Customer: Customer{
			Name: customerName,
			ID:   customerID,
		},
	}
}

func main() {
	order1 := Order{
		OrderID:     "55",
		OrderAmount: 90.0,
		OrderStatus: "pending",
		OrderCreatedAt: time.Now(),
		Customer: Customer{
			Name: "Paras",
			ID:   101,
		},
	}
	order1.UpdateAmount(88.8)

	order2 := NewOrder("66", 8878, "delivered", time.Now(), "Ravi", 202)

	fmt.Println(order2)
}
