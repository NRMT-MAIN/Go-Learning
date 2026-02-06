package main

import "fmt"

func sfExample() {

	num := 424
	fmt.Printf("%05d\n", num) //00424

	message := "Hello"
	fmt.Printf("|%10s|\n", message) // |     Hello|
	fmt.Printf("|%-10s|\n", message) // |Hello     |

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)
	// Hello
	// World!
	fmt.Println(message2) // Hello \nWorld!

	// sqlQuery := `SELECT * FROM users WHERE age > 30`
}