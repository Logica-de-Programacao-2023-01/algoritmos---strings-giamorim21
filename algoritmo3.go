package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e substitua todas as ocorrências desse caractere
//na string
//por outro caractere especificado pelo usuário.
//Informe ao usuário o resultado.

func main() {
	var s string
	var c string
	var n string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	fmt.Println("Digite o caracter a ser substituído:")
	fmt.Scan(&c)
	fmt.Println("Digite o novo caracter:")
	fmt.Scan(&n)

	novastring := strings.ReplaceAll(s, c, n)

	fmt.Println("Nova string:", novastring)

}

