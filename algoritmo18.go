package main

import (
	"fmt"
	"strconv"
)

//Solicite ao usuário uma string
//e informe se ela contém somente números (0 a 9).

func main() {
	var s string
	fmt.Println("Digite uma string:")
	fmt.Scan(&s)

	_, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println("A string contém somente números")
	} else {
		fmt.Println("A string não contém somente números")
	}
}

