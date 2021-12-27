package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Dio","LastName":"Surandito","Age":24,"Job":"Programmer","Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Job)
	fmt.Println(customer.Married)

	expected := &Customer{
		FirstName: "Dio",
		LastName:  "Surandito",
		Age:       24,
		Job:       "Programmer",
		Married:   false,
	}

	assert.Equal(t, expected, customer, "Both should be equal!")

}
