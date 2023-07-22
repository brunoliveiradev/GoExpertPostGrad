package main

import "fmt"

type ID int

var (
	a string  = "Com o type voce cria sua propria variavel tipada"
	b bool    = true
	c int     = 10
	d float64 = 1.2
	e ID      = 1
)

const f string = "sou uma constante"

func main() {
	fmt.Printf("O tipo da variavel 'd' é: %T \n", d)
	fmt.Printf("O valor da variavel 'd' é: %v \n", d)
	fmt.Printf("O tipo da variavel 'e' é: %T \n", e)
	fmt.Printf("O valor da variavel 'e' é: %v \n", e)
}
