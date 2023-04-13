package main

import (
	"fmt"
	"strconv"
)

//Escreva um programa que receba uma string
//e verifique se ela é um número válido em ponto flutuante.
//Informe ao usuário se é ou não.

func main() {
	var s string
	fmt.Println("Digite um número em ponto flutuante:")
	fmt.Scan(&s)

	_, err := strconv.ParseFloat(s, 64)
	if err == nil {
		fmt.Println("A string digitada é um número em ponto flutuante válido")
	} else {
		fmt.Println("A string digitada nao é um número em ponto flutuante válido")
	}
}

