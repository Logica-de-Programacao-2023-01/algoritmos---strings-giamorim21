package main

import (
	"fmt"
	"sort"
	"strings"
)

//Escreva um programa que receba duas strings
//e verifique se elas são anagramas.
//Informe ao usuário se são ou não.

func main() {
	var s1, s2 string
	fmt.Println("Digite a primeira string:")
	fmt.Scan(&s1)

	fmt.Println("Digite a segunda string:")
	fmt.Scan(&s2)

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		fmt.Println("As strings não são anagramas")
		return
	}
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	if string(runes1) == string(runes2) {
		fmt.Println("As strings são anagramas")
	} else {
		fmt.Println("As strings não são anagramas")
	}
}

