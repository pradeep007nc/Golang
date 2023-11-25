package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
)

type Product struct {
	ID           int     `json:"id"`
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	UnitPrice    float64 `json:"unitPrice"`
	ImageURL     string  `json:"imageUrl"`
	Active       bool    `json:"active"`
	UnitsInStock int     `json:"unitsInStock"`
	DateCreated  string  `json:"dateCreated"`
	LastUpdated  string  `json:"lastUpdated"`
}

type ProductList struct {
	Embedded struct {
		Products []Product `json:"products"`
	} `json:"_embedded"`
}

func main() {

	url := "https://localhost:8443/api/products?size=100"

	// Disable TLS verification for this example (not recommended for production)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create an HTTP client with the custom transport
	client := &http.Client{Transport: tr}

	response, err := client.Get(url);
	if err != nil{
		fmt.Println("bad request error ", err)
		return
	}

	defer response.Body.Close();

	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println("Error reading the response body:", err)
		return
	} 

	fmt.Println(string(body))
}
