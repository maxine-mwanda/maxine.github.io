# Orders API project

## Overview
MyProject is a Go-based RESTful API that interacts with a PostgreSQL database to manage `Customers` and `Orders`. It includes CRUD operations for both tables and integrates with Africa's Talking API to send SMS notifications upon successfully creating an order.

## Features
- **CRUD Operations**: Create, Read, Update, and Delete operations for `Customers` and `Orders`.
- **SMS Notifications**: Sends an SMS notification using Africa's Talking API when a new order is created.

## Prerequisites
- Go 1.16 or higher
- PostgreSQL
- Africa's Talking API credentials

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/myproject.git
    cd myproject
    ```

2. **Install dependencies**:
    ```sh
    go get github.com/gin-gonic/gin
    go get github.com/lib/pq
    ```

3. **Set up PostgreSQL**:
    - Create a database named `ordersAPI`.
    - Create the `Customers` and `Orders` tables using the following SQL:
      ```sql
      CREATE TABLE Customers (
          CustomerID SERIAL PRIMARY KEY,
          Name VARCHAR(100) NOT NULL,
          Code VARCHAR(50) NOT NULL
      );

      CREATE TABLE Orders (
          OrderID SERIAL PRIMARY KEY,
          CustomerID INT REFERENCES Customers(CustomerID),
          Item VARCHAR(100) NOT NULL,
          Amount DECIMAL(10, 2) NOT NULL,
          OrderTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
      );
      ```

## Running the Project

1. **Run the application**:
    ```sh
    go run main.go
    ```

2. **API Endpoints**:

    ### Customers
    - **Create Customer**: `POST /customers`
      ```json
      {
          "name": "John Doe",
          "code": "C001"
      }
      ```
    - **Get Customers**: `GET /customers`
    - **Update Customer**: `PUT /customers/:id`
      ```json
      {
          "customer_id": 1,
          "name": "Jane Doe",
          "code": "C002"
      }
      ```
    - **Delete Customer**: `DELETE /customers/:id`

    ### Orders
    - **Create Order**: `POST /orders`
      ```json
      {
          "customer_id": 1,
          "item": "Item1",
          "amount": 100.50
      }
      ```
    - **Get Orders**: `GET /orders`
    - **Update Order**: `PUT /orders/:id`
      ```json
      {
          "order_id": 1,
          "item": "Item2",
          "amount": 150.75
      }
      ```
    - **Delete Order**: `DELETE /orders/:id`

## SMS Notification
Upon successfully creating an order, an SMS notification is sent using Africa's Talking API. The SMS contains a custom message and is sent to the specified phone numbers.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Gin Gonic
- pq
- Africa's Talking

