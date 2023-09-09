package main

func main() {
	mapaInt := map[string]int{
		"key1": 10,
		"key2": 20,
		"key3": 30,
	}
	println(Sum(mapaInt))

	mapaFloat := map[string]float64{
		"key1": 10.5,
		"key2": 20.5,
		"key3": 30.5,
	}
	println(Sum(mapaFloat))

	m3 := map[string]MyInt{
		"key1": 10,
		"key2": 20,
		"key3": 30,
	}
	println(Sum(m3))

	// println(IsSameNumber(10, 20.5)) // this will not compile
	println(IsSameNumber(10, 10))

	// IsTheSame is a generic function that can receive any comparable type and return true if they are the same
	println(IsTheSame(10, 10.00))
}

type MyInt int

// Number is a generic interface that can be int or float64 and the symbol ~ is used to define the type constraint.
// Any type that is ~int can be used as a type for Number interface
type Number interface {
	~int | float64
}

// Sum is a generic function that can receive a map of int or float64 type and return the sum of all values
func Sum[T Number](m map[string]T) T {
	var soma T
	// ignore the key and get the value
	for _, v := range m {
		soma += v
	}
	return soma
}

// IsSameNumber is a generic function that can receive a Number type and return true if they are the same
func IsSameNumber[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false

}

// IsTheSame is a generic function that can receive any comparable type and return true if they are the same
func IsTheSame[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}
