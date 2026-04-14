package main

import (
	"fmt"
	"net/http"
)

func sevrerExample() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Fprintln(w , "Hello Server")
	})

	const serverAddr = ":8080"

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		panic(err)
	}

}