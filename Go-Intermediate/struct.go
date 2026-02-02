package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}


func (p *Person) haveBirthday() {
	p.age++
}

// Differnce between value receiver and pointer receiver
// Value receiver: method operates on a copy of the struct
// Pointer receiver: method operates on the original struct

// Struct and Method should be defined outside main function or any other function
// and then called inside main or any other function

//Embedded Struct 
type Employee struct {
	Person      // Embedded Person struct
	// Inherits fields and methods of Person
	employeeID  string
	department  string
}

func (e Employee) getEmployeeInfo() string {
	return e.firstName + " " + e.lastName + " - ID: " + e.employeeID + ", Department: " + e.department
}

func structsExample() {
	p1 := Person{
		firstName: "Alice",
		lastName:  "Johnson",
		age:       30,
	}

	p2 := Person{"Alice", "Johnson", 30} // Alternative way

	// Accessing struct fields
	println("First Name:", p1.firstName)
	println("Last Name:", p1.lastName)
	println("Age:", p1.age)
	println("First Name:", p2.firstName)
	println("Last Name:", p2.lastName)
	println("Age:", p2.age)

	// Modulfying struct fields
	p1.age = 31
	println("Updated Age of p1:", p1.age)

	// Anonymous Struct
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "user123@example.com",
	}
	fmt.Println(user)

	fmt.Println("Full Name of p1:", p1.fullName())
	fmt.Println("Full Name of p2:", p2.fullName())

	println("Age of p1 before birthday:", p1.age)
	p1.haveBirthday()
	println("Age of p1 after birthday:", p1.age)

	// Without pointer receiver, age won't change
	// To see the effect, we need to change haveBirthday to use pointer receiver
	
	fmt.Println("---- Embedded Struct Example ----")
	emp := Employee{
		Person: Person{
			firstName: "Bob",
			lastName:  "Smith",
			age:       28,
		},
		employeeID: "E12345",
		department: "Engineering",
	}

	fmt.Println("Employee Info:", emp.getEmployeeInfo())
	fmt.Println("Employee Full Name:", emp.fullName())
	fmt.Println("Employee Age before birthday:", emp.age)
	emp.haveBirthday()
	fmt.Println("Employee Age after birthday:", emp.age)

}