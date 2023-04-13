package main

//Solicite ao usuário uma string
//e informe se ela é uma sequência numérica
//crescente (exemplo: "123" ou "4567").

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("Digite uma sequência numérica: ")
	fmt.Scan(&s)

	previonum := -1
	isIncreasing := true
	for _, char := range s {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Entrada inválida.")
			return
		}
		if num <= previonum {
			isIncreasing = false
			break
		}
		previonum = num
	}

	if isIncreasing {
		fmt.Println("A sequência é numérica crescente.")
	} else {
		fmt.Println("A sequência não é numérica crescente.")
	}
}

