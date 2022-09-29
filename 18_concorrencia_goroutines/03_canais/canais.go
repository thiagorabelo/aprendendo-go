package main

import (
	"fmt"
	"time"
)

func canais1() {
	canal := make(chan string)
	go escrever("Olá, mundo!", canal)

	// Com loop infinito, seta e break
	for {
		mensagem, aberto := <-canal
		if !aberto {
			fmt.Println("O canal foi fechado")
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do trecho")
	fmt.Println()
}

func canais2() {
	canal := make(chan string)
	go escrever("Hello World!", canal)

	// Com range
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do trecho")
	fmt.Println()
}

func main() {
	fmt.Println("Canais com sitaxe de loop, infinito, seta e break")
	canais1()

	fmt.Println("Canais com sitaxe de range")
	canais2()
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%s [%d]", texto, i)
		time.Sleep(time.Second)
	}
	close(canal)
}
