package main

import "strconv"

func typeConversionExample() {
	var a int = 32
	b := int32(a) // Type conversion from int to int32
	println("Value of a:", a)
	println("Value of b (int32):", b)

	var c float64 = 3.14
	d := int(c) // Type conversion from float64 to int (truncates the decimal part)
	println("Value of c:", c)
	println("Value of d (int):", d)

	var e string = "123"
	f, err := strconv.Atoi(e) // Type conversion from string to int using strconv.Atoi
	if err != nil {
		println("Error converting string to int:", err.Error())
	} else {
		println("Value of e:", e)
		println("Value of f (int):", f)
	}

	g := "hELLO WOEKD あいうえお"
	var h []byte
	h = []byte(g) // Type conversion from string to byte slice
	println("Value of g:", g)
	println("Value of h (byte slice):", h)

	i :=  []byte{104, 69, 76, 76, 79} // Type conversion from byte slice to string
	j := string(i)
	println("Value of i (byte slice):", i)
	println("Value of j (string):", j)
}