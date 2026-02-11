package main

import (
	"encoding/base64"
	"fmt"
)

func encodingExample() {
	data := []byte("Hello , Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded , err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error : " , err)
		return
	}

	fmt.Println("Decoded : " , string(decoded))

	// URL Safe encoding
	urlEncoded := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println("URL-safe:", urlEncoded)
    
    decoded, _ = base64.URLEncoding.DecodeString(urlEncoded)
    fmt.Println("Decoded:", string(decoded))
}