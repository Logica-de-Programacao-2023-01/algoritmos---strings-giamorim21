package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e remova todas as vogais.
//Informe ao usuário o resultado.

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&s)

	output := removeVowels(s)

	fmt.Printf("A string sem vogais é: %s\n", output)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

