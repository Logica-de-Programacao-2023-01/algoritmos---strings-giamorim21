package main

import (
	"fmt"
	"unicode"
)

//Solicite ao usuário uma string e informe se ela é está em
//camelCase e quantas palavras possuí. CamelCase é quando a
//primeira letra de cada palavra é maiúscula, e as demais
//são minúsculas. Exemplo: "estaStringEstaEmCamelCase".

func main() {
	var s string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	camelCase := true
	count := 1
	for i, char := range s {
		if i == 0 {
			if !unicode.IsUpper(char) {
				camelCase = false
				break
			}
		} else if char == ' ' {
			camelCase = false
			break
		} else if unicode.IsUpper(char) {
			count++
		}
	}
	if camelCase {
		fmt.Printf("A string '%s' está em camelCase e possui %d palavra(s)\n", s, count)
	} else {
		fmt.Printf("A string '%s' não está em camelCase\n", s)
	}
}

