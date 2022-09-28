package main

import "fmt"

func fibonacci(indice int) int {
	if indice <= 1 {
		return 1
	}
	return fibonacci(indice-1) + fibonacci(indice-1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	const total int = 45

	tarefas := make(chan int, total)
	resultados := make(chan int, total)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < total; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < total; i++ {
		resultado := <-resultados
		fmt.Printf("%d) %v\n", i, resultado)
	}
}
