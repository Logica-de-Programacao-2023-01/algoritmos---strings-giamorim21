package main

import (
	"fmt"
	"strings"
)

//Solicite ao usuário uma string
//e substitua todas as ocorrências de uma letra por outra
//informada pelo usuário.

func main() {
	var s, c, n string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	fmt.Println("Digite a ser substituída:")
	fmt.Scan(&c)

	fmt.Println("Digite a letra substituita:")
	fmt.Scan(&n)

	output := strings.ReplaceAll(s, c, n)

	fmt.Println("A string modificada é: %s\n", output)
}

