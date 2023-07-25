package main

import "fmt"

func main() {
	canal := make(chan string, 200)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	canal <- "Programando em Go De Novo!"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
