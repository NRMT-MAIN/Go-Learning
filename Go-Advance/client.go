package main

import (
	"fmt"
	"io"
	"net/http"
)

func clientExample() {
	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
	defer resp.Body.Close()

	body , err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(body))
}
