package main


func adder() func() int {
	sum := 0
	return func() int {
		sum += 1
		return sum
	}
}

