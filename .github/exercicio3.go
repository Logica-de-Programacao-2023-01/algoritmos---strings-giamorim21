package main

import (
	"fmt"
	"strings"
)

//Peça ao usuário para digitar uma string e um caractere
//e retorne o número de
//ocorrências desse caractere na string.

func main() {
	var s string
	var c string

	fmt.Println("Digite algo:")
	fmt.Scan(&s)
	fmt.Println("Digite um caracter:")
	fmt.Scan(&c)

	fmt.Print(strings.Count(s, c))
}
