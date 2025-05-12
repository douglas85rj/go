package main

import "fmt"

func main() {
	var variavel1 string = "first"
	variavel2 := "second"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// invertendo valores das variaveis
	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)

	const constante1 string = "first const"
	fmt.Println(constante1)
}
