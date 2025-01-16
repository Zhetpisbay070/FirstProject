package server

import "time"

type CreateOrderRequest struct {
	UserID       string   `json:"user_id"`
	Products     []string `json:"products"`
	Price        float64  `json:"price"`
	DeliveryType string   `json:"delivery_type"`
	AddressID    string   `json:"address_id"`
}

type Order struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	ProductIDs       []string  `json:"product_i_ds"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeliveryDeadLine time.Time `json:"delivery_dead_line"`
	Price            float64   `json:"price"`
	DeliveryType     string    `json:"delivery_type"`
	Address          string    `json:"address"`
	OrderStatus      string    `json:"order_status"`
}
