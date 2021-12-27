package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Dio",
		LastName:  "Surandito",
		Age:       24,
		Job:       "Programmer",
		Married:   false,
	}

	encoder.Encode(customer)
}
