package main

import "fmt"

//Peça ao usuário para digitar duas strings
//e retorne a concatenação das duas.

func main() {
	var s string
	var t string

	fmt.Println("Digite algo:")
	fmt.Scan(&s)

	fmt.Println("Digite algo novamente:")
	fmt.Scan(&t)

	s3 := s + "," + t
	fmt.Println(s3)
}
