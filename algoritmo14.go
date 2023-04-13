package main

import (
	"fmt"
	"strconv"
)

//Solicite ao usuário uma string
//e informe se ela é uma sequência numérica decrescente
//(exemplo: "987" ou "54321").

func main() {
	var s string
	fmt.Println("Digite uma sequencia numerica:")
	fmt.Scan(&s)

	previonum := 10
	decrescente := true
	for _, char := range s {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Entrada inválida")
			return
		}
		if num >= previonum {
			decrescente = false
			break
		}
		previonum = num
	}
	if decrescente {
		fmt.Println("A sequencia é numérica decrescente")
	} else {
		fmt.Println("A sequencia não é numérica decrescente")
	}
}

