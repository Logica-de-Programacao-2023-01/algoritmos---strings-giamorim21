package main

import "fmt"

//Solicite ao usuário uma string
//e imprima a mesma string invertida.

func main() {
	var s string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	inverso := ""
	for i := len(s) - 1; i >= 0; i-- {
		inverso += string(s[i])
	}
	fmt.Println("String invertida:", inverso)
}

