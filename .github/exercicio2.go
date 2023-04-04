package main

import "fmt"

//Peça ao usuário para digitar uma string
//e retorne o número de caracteres nessa string.

func main() {
	var s string
	fmt.Println("Digite algo:")
	fmt.Scan(&s)

	fmt.Println("len(s) =", len(s))
}
