package main

func conditional() {
	age := 20
	if age < 18 {
		println("Minor")
	}
	if age >= 18 && age < 65 {
		println("Adult")
	}
	if age >= 65 {
		println("Senior Citizen")
	}
	//if-else
	if age < 18 {
		println("Minor")
	} else if age >= 18 && age < 65 {
		println("Adult")
	}	 else {
		println("Senior Citizen")
	}	
	//Switch case
	day := 3
	switch day {
	case 1:
		println("Monday")		
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}

	//Switch without expression
	switch {
	case age < 18:
		println("Minor")	
	case age >= 18 && age < 65:
		println("Adult")
	default:
		println("Senior Citizen")
	}
}