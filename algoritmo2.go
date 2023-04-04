package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva um programa que receba uma string
//e remova todos os espaços em branco.
//Informe ao usuário o resultado.

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite uma string:")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(strings.ReplaceAll(frase, " ", ""))

}
