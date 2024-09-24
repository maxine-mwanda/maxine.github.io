package models

import "time"

type Order struct {
	OrderID    int       `json:"order_id"`
	CustomerID int       `json:"customer_id"`
	Item       string    `json:"item"`
	Amount     float64   `json:"amount"`
	OrderTime  time.Time `json:"order_time"`
}
