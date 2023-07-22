package main

import "fmt"

var (
	a = "Class about Maps - HashTables[key:value]"
)

func main() {
	fmt.Println(a)

	salarios := map[string]int{"Bruno": 1000, "João": 2000, "Maria": 3000}
	//wages := make(map[string]int)
	//wage := map[string]int{} // empty
	fmt.Println(salarios)

	delete(salarios, "Bruno")
	salarios["Ana"] = 5000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
