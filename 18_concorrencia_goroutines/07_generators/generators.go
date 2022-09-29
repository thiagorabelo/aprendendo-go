package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

// Encapsula a complexidade da goroutine apenas devolvendo um canal, de onde
// os dados de ineresse podem ser extraídos.
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; ; i++ {
			canal <- fmt.Sprintf("%s [%d]", texto, i)
			time.Sleep(time.Second / 2)
		}
	}()

	return canal
}
