package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khanendy",
		Age:        30,
		Married:    true,
		Addresses: []Address{
			{
				Street:     "Jalan Setapak",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "Jalan Raya",
				Country:    "Indonesia",
				PostalCode: "67890",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}