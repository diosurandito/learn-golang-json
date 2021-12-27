package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Dio",
		LastName:  "Surandito",
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Dio","LastName":"Surandito","Age":0,"Job":"","Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

	expected := &Customer{
		FirstName: "Dio",
		LastName:  "Surandito",
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}
	assert.Equal(t, expected, customer, "Both should be equal!")
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Dio",
		Addresses: []Address{
			{
				Street:     "Jl. H. Maja",
				Country:    "Jakarta",
				PostalCode: "52218",
			},
			{
				Street:     "Jl. Desa Wangandalem",
				Country:    "Brebes",
				PostalCode: "52214",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `{"FirstName":"Dio","LastName":"","Age":0,"Job":"","Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl. H. Maja","Country":"Jakarta","PostalCode":"52218"},{"Street":"Jl. Desa Wangandalem","Country":"Brebes","PostalCode":"52214"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

	expected := &Customer{
		FirstName: "Dio",
		Addresses: []Address{
			{
				Street:     "Jl. H. Maja",
				Country:    "Jakarta",
				PostalCode: "52218",
			},
			{
				Street:     "Jl. Desa Wangandalem",
				Country:    "Brebes",
				PostalCode: "52214",
			},
		},
	}
	assert.Equal(t, expected, customer, "Both should be equal!")
}

func TestOnlyJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `[{"Street":"Jl. H. Maja","Country":"Jakarta","PostalCode":"52218"},{"Street":"Jl. Desa Wangandalem","Country":"Brebes","PostalCode":"52214"}]`
	jsonBytes := []byte(jsonString)

	addresess := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresess)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresess)

	expected := &[]Address{
		{
			Street:     "Jl. H. Maja",
			Country:    "Jakarta",
			PostalCode: "52218",
		},
		{
			Street:     "Jl. Desa Wangandalem",
			Country:    "Brebes",
			PostalCode: "52214",
		},
	}

	assert.Equal(t, expected, addresess, "Both should be equal!")
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresess := []Address{
		{
			Street:     "Jl. H. Maja",
			Country:    "Jakarta",
			PostalCode: "52218",
		},
		{
			Street:     "Jl. Desa Wangandalem",
			Country:    "Brebes",
			PostalCode: "52214",
		},
	}

	bytes, _ := json.Marshal(addresess)
	fmt.Println(string(bytes))
}
