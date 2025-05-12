package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	auxiliar.Escrever2()
}

//se uma função começa com letra maiúscula, ela é pública e pode ser acessada por outros pacotes
//se uma função começa com letra minúscula, ela é privada e só pode ser acessada dentro do mesmo pacote
//para executar o código, é necessário rodar o comando go run main.go
//para compilar o código, é necessário rodar o comando go build main.go
//para executar o código compilado, é necessário rodar o comando ./main
//para instalar o pacote, é necessário rodar o comando go install
