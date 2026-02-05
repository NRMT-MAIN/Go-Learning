package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func stringExample() {

	str1 := "Hello  Go!"
	str2 := "World"

	result := str1 + " " + str2

	fmt.Println(len(str2)) //5
	fmt.Println(str1[0]) //72
	fmt.Println(str1[1 : 7]) //ello

	fmt.Println(result) //Hello  Go! World
	
	// string conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3)) // 2

	// string splitting
	fruits := "apple,orange,banana"
	parts := strings.Split(fruits , ",")
	fmt.Println(parts) // [apple orange banana]

	// Concvatenation using slice
	countries := []string{"Germany" , "France" , "India"}
	joined := strings.Join(countries , " , ")
	fmt.Println(joined) // Germany , France , India

	// Contains
	fmt.Println(strings.Contains(joined , "Italy")) //false
	fmt.Println(strings.Contains(joined , "India")) //true

	// Replacement
	replaced := strings.Replace(joined , "India" , "Italy" , 1)
	fmt.Println(replaced) // Germany , France , Italy

	//Leading and trailing space
	strwspace := " Hello Everyone!"
	fmt.Println(strwspace) // Hello Everyone!
	fmt.Println(strings.TrimSpace(strwspace)) //Hello Everyone!

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))
	//  hello everyone!
 	//  HELLO EVERYONE!

	fmt.Println(strings.Repeat("foo ", 3)) //foo foo foo

	fmt.Println(strings.Count("Hello", "l")) // 2
	fmt.Println(strings.HasPrefix("Hello", "he")) // false
	fmt.Println(strings.HasSuffix("Hello", "lo")) // true
	fmt.Println(strings.HasSuffix("Hello", "la")) // false

	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches) // [1 123 11]

	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6)) // 9 

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result2 := builder.String()
	fmt.Println(result2) // Hello, world!

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result) // Hello, world! How are you

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result) // Starting fresh!
}