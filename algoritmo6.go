package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e conte quantas palavras ela contém.
//Informe ao usuário o resultado.

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&input)

	words := strings.Fields(input)
	numWords := len(words)

	fmt.Printf("A string contém %d palavra(s).\n", numWords)
}
