package main

import (
	"fmt"
	"os"
)

func envVarExample() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var : " , user)
	fmt.Println("Home env var : " , home)

	os.Setenv("FRUIT" , "APPLE")


	fruit := os.Getenv("FRUIT")
	fmt.Println("Fruit env var : " , fruit) // Output: Fruit env var :  APPLE	

	// Unset the environment variable
	os.Unsetenv("FRUIT")
	fruit = os.Getenv("FRUIT")
	fmt.Println("Fruit env var after unsetting : " , fruit) // Output: Fruit env var after unsetting :

	// List all environment variables
	fmt.Println("All Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Check if an environment variable exists
	if value, exists := os.LookupEnv("USER"); exists {
		fmt.Println("USER env var exists with value:", value)
	} else {
		fmt.Println("USER env var does not exist")
	}

	// Get the value of an environment variable with a default fallback
	path := os.Getenv("PATH")
	if path == "" {
		path = "/usr/local/bin:/usr/bin:/bin"
	}
	fmt.Println("PATH env var:", path)
	
	// Get the value of an environment variable with a default fallback using os.LookupEnv
	if value, exists := os.LookupEnv("SHELL"); exists {
		fmt.Println("SHELL env var:", value)
	} else {
		fmt.Println("SHELL env var not set, using default: /bin/bash")
	}

	// Get the value of an environment variable with a default fallback using os.LookupEnv and a default value
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}
	fmt.Println("SHELL env var with default fallback:", shell)
}