package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo!"), escrever("Olá GoLang"))

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

// Encapsula a complexidade da goroutine apenas devolvendo um canal, de onde
// os dados de ineresse podem ser extraídos.
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; ; i++ {
			sleepTime := time.Millisecond * time.Duration(rand.Intn(2000))
			canal <- fmt.Sprintf("%s [%d] (sleep %d millis)", texto, i, sleepTime/time.Millisecond)
			time.Sleep(sleepTime)
		}
	}()

	return canal
}
