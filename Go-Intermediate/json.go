package main

import (
	"encoding/json"
	"fmt"
)

type PersonJ struct {
	Name         string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty`
	Address      Address `json:"address"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}

func JsonExample() {
	person := PersonJ{
		Name: "Nirmit Sahu",
		Age:  22,
	}

	// Marshalling
	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error marshalling to JSON : ", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := PersonJ {
		Name: "Jane",
		Age: 30,
		EmailAddress: "jane@fakemail.com",
		Address: Address{
			City: "New York",
			State: "U.S",
		},
	}

	jsonData1 , err := json.Marshal(person1)

	if err != nil {
		fmt.Println("Error marshalling to JSON : " , err)
	}

	fmt.Println(string(jsonData1))

	jsonData2 := "{\"name\": \"John\", \"age\": 25, \"email\": \"john@fakemail.com\", \"address\": {\"city\": \"Los Angeles\", \"state\": \"U.S\"}}"
	var person2 PersonJ

	err = json.Unmarshal([]byte(jsonData2) , &person2)
	if err != nil {
		fmt.Println("Error unmarshalling JSON : ", err)
		return
	}
	fmt.Println(person2)
	fmt.Println(person2.Name)
	fmt.Println(person2.Age)
	fmt.Println(person2.EmailAddress) 
	fmt.Println(person2.Address.City)

	// Handling JSON with unknown fields
	jsonData3 := "{\"name\": \"Alice\", \"age\": 28, \"email\": \"alice@fakemail.com\", \"unknown_field\": \"some_value\"}"
	var person3 PersonJ

	err = json.Unmarshal([]byte(jsonData3) , &person3)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData3) , &data)

	if err != nil {
		fmt.Println("Error unmarshalling JSON : ", err)
		return
	}
	fmt.Println(person3)
	fmt.Println(data)
	
}
