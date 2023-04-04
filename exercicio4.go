package main

//Peça ao usuário para digitar uma string
//e retorne a mesma string com as letras em
//maiúsculo.

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	fmt.Println(strings.ToUpper(s))
}
