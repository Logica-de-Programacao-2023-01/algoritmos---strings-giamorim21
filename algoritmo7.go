package main

import (
	"fmt"
	"unicode"
)

//Solicite ao usuário uma string
//e informe se ela contém pelo menos um número.

func main() {
	var s string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	numero := false
	for _, char := range s {
		if unicode.IsNumber(char) {
			numero = true
			break
		}
	}
	if numero {
		fmt.Println("A string contém pelo menos um número")
	} else {
		fmt.Println("A string não contém nenhum número")
	}
}
