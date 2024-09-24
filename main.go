package main

import (
	"ordersAPI/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/customers", handlers.GetCustomers)
	r.POST("/customers", handlers.CreateCustomer)
	r.PUT("/customers/:id", handlers.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)

	r.GET("/orders", handlers.GetOrders)
	r.POST("/orders", handlers.CreateOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)

	r.Run(":8080")
}
