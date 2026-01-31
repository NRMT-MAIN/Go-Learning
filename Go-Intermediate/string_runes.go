package main

import "fmt"

func string_runes() {
	message := "Hello, 世界" // A string containing both ASCII and non-ASCII characters
	println("Original String:", message)

	// for i, ch := range message {
	// 	println(i, ":", string(ch), "->", ch)
	// }

	// // Convert string to a slice of runes to properly handle Unicode characters
	// runes := []rune(message)
	// println("Runes in the string:")
	// for i, r := range runes {
	// 	println(i, ":", string(r), "->", r)
	// }

	//Output
	// Original String: Hello, 世界
	// 0 : H -> 72
	// 1 : e -> 101
	// 2 : l -> 108
	// 3 : l -> 108
	// 4 : o -> 111
	// 5 : , -> 44
	// 6 :   -> 32
	// 7 : 世 -> 19990
	// 10 : 界 -> 30028
	// Runes in the string:
	// 0 : H -> 72
	// 1 : e -> 101
	// 2 : l -> 108
	// 3 : l -> 108
	// 4 : o -> 111
	// 5 : , -> 44
	// 6 :   -> 32
	// 7 : 世 -> 19990
	// 8 : 界 -> 30028

	message = "Hello, \nGo!"
	fmt.Println("Updated String with escape sequence:")
	fmt.Println(message)

	for i, ch := range message {
		fmt.Printf("Rune %c at byte position %d\n", ch, i)
	}

	fmt.Println("Length of string in bytes:", len(message)) // 11 -- includes escape character
	fmt.Println("Length of string in runes:", len([]rune(message))) // 11

	fmt.Println(message[0]) ; // 72 (byte value of 'H')
	fmt.Println(string(message[0])) ; // 'H'

	greeting := "Hii there!"
	name := "Alice"
	fullMessage := greeting + " " + name
	fmt.Println("Concatenated String:", fullMessage) // Hii there! Alice

	str1 := "Go"
	str2 := "Go"
	if str1 == str2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}

	str2 = "Golang"
	fmt.Println(str1 > str2) // false
	fmt.Println(str1 < str2) // true

	var ch rune = 'A'
	fmt.Println("Rune value:", ch)                     // Rune value: 65
	fmt.Printf("Character: %c, Unicode code point: %U\n", ch, ch) // Character: A, Unicode code point: U+0041

	fmt.Println(string(ch))
	fmt.Printf("Type of ch: %T\n", ch) // Type of ch: int32

	var ch3 string = "✅"
	fmt.Println("String value:", ch3)                     // String value: ✅
	fmt.Printf("Character: %s\n", ch3)

	var ch2 rune = '✅'
	fmt.Println("Rune value:", ch2)                     // Rune value: 10003
	fmt.Printf("Character: %c, Unicode code point: %U\n", ch2, ch2) // Character: ✅, Unicode code point: U+2705

	// Both represent the same character but in different types
	// ch3 is a string, while ch2 is a rune (int32)
	// This demonstrates the difference between string and rune types in Go
	// Output:
	// String value: ✅
	// Character: ✅
	// Rune value: 10003
	// Character: ✅, Unicode code point: U+2705

	var ch4 string = "✅" // Correct way to assign a Unicode character to a string
	fmt.Println(ch4) // Error: cannot use '✅' (untyped rune constant) as string value
	// '✅' is a rune constant, not a string. To assign it to a string variable, use double quotes.
	// "✅" is a string literal. Internally, it is represented as a sequence of bytes corresponding to the UTF-8 encoding of the character.

	//Diiference between string and rune
	//======================================
	// A string is a sequence of bytes, while a rune is an alias for int32 and represents a single Unicode code point.
	// Strings can contain multiple runes (characters), while a rune represents just one character.
	// Strings are immutable, whereas runes can be manipulated as integer values.
	// Strings are used for text data, while runes are often used for character-level operations.
	// Strings are enclosed in double quotes, while runes are enclosed in single quotes.
	// Strings can be concatenated, while runes can be converted to strings or used in character comparisons.
	// Strings have a length in bytes, while runes have a length of 4 bytes (int32).
	// Strings can contain escape sequences, while runes represent actual characters.

	//var ch5 rune = "A" // Error: cannot use "A" (type string) as type rune in assignment
	// 'A' is a rune constant, while "A" is a string literal. To assign a character to a rune variable, use single quotes.

	//var ch6 rune  = 'AB' // Error: invalid rune literal (more than one character)
	// A rune can only represent a single Unicode code point. 'AB' contains two characters, which is invalid for a rune literal.

	//var ch7 string = 'A' // Error: cannot use 'A' (type rune) as type string in assignment
	// 'A' is a rune constant, while a string variable requires a string literal enclosed in double quotes.

	
	
}