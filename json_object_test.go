package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Job       string
	Married   bool
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Dio",
		LastName:  "Surandito",
		Age:       24,
		Job:       "Programmer",
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
