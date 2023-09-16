package main

func main() {

	a := 2
	b := 3

	// if is the same as other languages
	if a > b {
		println("a is greater than b")
	} else if a < b {
		println("a is less than b")
	} else {
		println("a is equal to b")
	}

	// can be used to initialize variables
	if c := 10; c > 0 {
		println("c is greater than zero")
	}

	// switch is the same as other languages
	switch a {
	case 1:
		println("a is equal to 1")
	case 2:
		println("a is equal to 2")
	default:
		println("a is not equal to 1 or 2")
	}

	// switch can be used to initialize variables
	switch c := 10; {
	case c > 0:
		println("c is greater than zero")
	case c < 0:
		println("c is less than zero")
	default:
		println("c is equal to zero")
	}
}
