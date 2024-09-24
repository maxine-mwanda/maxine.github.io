package handlers

import (
	"net/http"
	"ordersAPI/models"
	"ordersAPI/utils"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT CustomerID, Name, Code FROM Customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.CustomerID, &customer.Name, &customer.Code); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}
	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO Customers (Name, Code) VALUES ($1, $2) RETURNING CustomerID`
	err := db.QueryRow(sqlStatement, customer.Name, customer.Code).Scan(&customer.CustomerID)
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

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customerID := c.Param("id")
	sqlStatement := `UPDATE Customers SET Name=$2, Code=$3 WHERE CustomerID=$1`
	_, err := db.Exec(sqlStatement, customerID, customer.Name, customer.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

func DeleteCustomer(c *gin.Context) {
	db, _ := utils.ConnectDB()
	defer db.Close()

	customerID := c.Param("id")
	sqlStatement := `DELETE FROM Customers WHERE CustomerID=$1`
	_, err := db.Exec(sqlStatement, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
