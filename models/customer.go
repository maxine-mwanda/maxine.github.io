package models

type Customer struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
}
