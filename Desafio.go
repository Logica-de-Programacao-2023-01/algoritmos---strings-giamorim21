package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e um padrão (outro string) e retorne todos os índices em que
//o padrão ocorre na string. Informe ao usuário o resultado.

func main() {
	var s, p string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)
	fmt.Println("Digite um padrão:")
	fmt.Scan(&p)

	indice := []int{}
	start := 0

	for {
		index := strings.Index(s[start:], p)
		if index == -1 {
			break
		}
		indice = append(indice, index+start)
		start += index + len(p)
	}
	fmt.Printf("O padrão '%s' ocorre nas posições: %v", p, indice)
}
