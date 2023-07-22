package main

var (
	a string
	b bool
	c int
	d float64
)

const e = "Variavel Constante"

func main() {
	println("Hello, Wolrd!")
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	shortHand()
}

func shortHand() {
	var x string = "Método short-hand abaixo substitui o formato de declaração 'var nomeVariavel tipo = valor'"
	y := "Método short-hand 'nomeVariavel := valor' usado para atribuir valores pela primeira vez a uma variavel" // String

	println(x + "\n" + y)
}
