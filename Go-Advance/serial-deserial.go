package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func serialDeserialExample() {
	// Serialization
	user := User{Name: "Alice", Email: "alice@example.com"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized JSON:", string(jsonData))
	// Output: Serialized JSON: {"name":"Alice","email":"alice@example.com"}

	// Deserialization
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized User: %+v\n", user)
	// Output: Deserialized User: {Name:Alice Email:alice@example.com}

	// Encoder and Decoder - used for streaming data from files or network

	jsonData1 := `{"name":"Bob","email":"bob@example.com"}`
	reader := json.NewDecoder(strings.NewReader(jsonData1))
	var user1 User
	err = reader.Decode(&user1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized User1: %+v\n", user1)
	// Output: Deserialized User1: {Name:Bob Email:bob@example.com}
	fmt.Printf("User1 Name: %s, Email: %s\n", user1.Name, user1.Email)
	// Output: User1 Name: Bob, Email: bob@example.com

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user1)
	
	if err != nil {
		panic(err)
	}
	fmt.Println("Encoded JSON:", buf.String()) 
	// Encoded JSON: {"name":"Bob","email":"bob@example.com"}
}