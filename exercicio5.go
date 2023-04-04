package main

//Peça ao usuário para digitar uma string
//e um número n e retorne a mesma string com as
//n primeiras letras em maiúsculo.

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int

	fmt.Println("Digite uma string:")
	fmt.Scan(&s)
	fmt.Println("Digite um número:")
	fmt.Scan(&n)

	var result string
	for i := 0; i < n; i++ {
		result = result + string(s[i])
	}
	fmt.Println(strings.ToUpper(result))
}