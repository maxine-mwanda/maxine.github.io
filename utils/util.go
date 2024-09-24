package utils

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/url"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "A$$aP123"
	dbname   = "ordersAPI"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SendPostRequest() error {
	apiURL := "https://api.sandbox.africastalking.com/version1/messaging"
	data := url.Values{}
	data.Set("username", "sandbox")
	data.Set("to", "+254710367960")
	data.Set("message", "Hello World!")
	data.Set("from", "myShortCode")

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", "MyAppApiKey")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	fmt.Println("SMS sent successfully:", string(body))
	return nil
}
