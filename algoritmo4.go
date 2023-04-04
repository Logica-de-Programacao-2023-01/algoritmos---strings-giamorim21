package main

import "fmt"

//Solicite ao usuário duas strings
//e informe se elas são iguais ou diferentes.

func main() {
	var s string
	var t string

	fmt.Println("Digite uma string:")
	fmt.Scan(&s)
	fmt.Println("Digite outra string:")
	fmt.Scan(&t)

	if s == t {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}
}