package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

func main() {
	var s string
	fmt.Println("Digite um string:")
	fmt.Scan(&s)

	fmt.Println(strings.ToUpper(s))
}
