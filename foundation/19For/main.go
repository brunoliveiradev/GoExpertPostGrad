package main

func main() {

	// for is the only loop in Go
	for i := 0; i < 4; i++ {
		println(i)
	}

	// using for to transversing a slice
	numeros := []string{"zero", "um", "dois", "tres", "quatro", "cinco"}
	for _, v := range numeros {
		println(v)
	}

	// for can be used as a while loop
	i := 8
	for i < 10 {
		println(i)
		i++
	}

}
