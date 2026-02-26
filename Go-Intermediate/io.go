package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ioExample() {
	reader := strings.NewReader("Hello, World! This is an example of io.Reader in Go.")
	fmt.Println("Reader:", reader)

	buf := make([]byte, 5)

	for {
		n , err := reader.Read(buf)
		
		if err == io.EOF {
			fmt.Println("End of reader reached")
			break
		}

		if err != nil {
			fmt.Println("Error reading from reader:", err)
			break
		}
		fmt.Println("Read", n, "bytes:", string(buf[:n]))
	}

	message := []byte("Hello, io.Writer!")

	n , err := os.Stdout.Write(message)
	if err != nil {
		fmt.Println("Error writing to stdout:", err)
		return
	}

	fmt.Println("Wrote", n, "bytes to stdout")

}