package handlers

import (
	"net/http"
	"ordersAPI/models"
	"ordersAPI/utils"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT OrderID, CustomerID, Item, Amount, OrderTime FROM Orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.CustomerID, &order.Item, &order.Amount, &order.OrderTime); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO Orders (CustomerID, Item, Amount, OrderTime) VALUES ($1, $2, $3, DEFAULT) RETURNING OrderID`
	err := db.QueryRow(sqlStatement, order.CustomerID, order.Item, order.Amount).Scan(&order.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Send POST request to Africa's Talking API
	err = utils.SendPostRequest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()
	orderID := c.Param("id")
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE Orders SET Item=$2, Amount=$3 WHERE OrderID=$1`
	_, err := db.Exec(sqlStatement, orderID, order.Item, order.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	orderID := c.Param("id")
	sqlStatement := `DELETE FROM Orders WHERE OrderID=$1`
	_, err := db.Exec(sqlStatement, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
