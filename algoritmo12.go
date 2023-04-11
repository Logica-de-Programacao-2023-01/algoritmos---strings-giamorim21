package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Digite uma string:")
	fmt.Scan(&x)

	letras := strings.Split(x, "")

	if len(letras)%2 != 0 {
		palindromo := true
		for i := 0; i < len(letras)/2; i++ {
			if letras[i] != letras[len(letras)-1-i] {
				palindromo = false
				break
			}
		}
		if palindromo {
			fmt.Printf("É palindromo")
		}

	} else {
		fmt.Printf("A string %x nao é palindromo", x)
	}
}

