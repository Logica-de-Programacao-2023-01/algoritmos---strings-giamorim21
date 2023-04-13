package main

import "fmt"

//Solicite ao usuário duas strings
//e informe se a segunda é uma substring da primeira.

func main() {
	var s1, s2 string
	fmt.Println("Digite a primeira string:")
	fmt.Scan(&s1)

	fmt.Println("Digite a segunda string:")
	fmt.Scan(&s2)

	if len(s1) < len(s2) {
		fmt.Println("Erro: a segunda string é maior que a primeira")
	} else {
		found := false
		for i := 0; i <= len(s1)-len(s2); i++ {
			if s1[i:i+len(s2)] == s2 {
				found = true
				break
			}
		}
		if found {
			fmt.Println("A segunda string é uma substring da primeira")
		} else {
			fmt.Println("A segunda string não é uma substring da primeira")
		}
	}
}

