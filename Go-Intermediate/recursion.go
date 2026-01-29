package main

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n / 10)
}

func fibonacci(n int) int {	
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}