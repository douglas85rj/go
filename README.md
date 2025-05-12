# Desenvolvimento Go

Prática em desenvolvimento na linguagem de programação Go

Anotações

* Se uma função começa com letra maiúscula, ela é pública e pode ser acessada por outros pacotes.
* Se uma função começa com letra minúscula, ela é privada e só pode ser acessada dentro do mesmo pacote.
* Para executar o código, é necessário rodar o comando `go run main.go`
* Para compilar o código, é necessário rodar o comando `go build main.go`
* Para executar o código compilado, é necessário rodar o comando `./main`
* Para instalar o pacote, é necessário rodar o comando

  `go install`
* Para utilizar um pacote externo utilizar o comando
  `go get + *url do pacote*`
* Remover dependências que não estão sendo utilizadas no *go.mod   `go mod tidy`
