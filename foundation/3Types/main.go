package main

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
	println("Hello, Wolrd!")
	println(a)
	println(e)
	println(f)

}
