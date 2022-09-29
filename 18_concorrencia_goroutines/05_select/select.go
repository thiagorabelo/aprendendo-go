package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// for {
	// 	/* Acaba que o tempo em espera é o do canal que demorar mais. */

	// 	msg1 := <-canal1 // Para e aguarda canal1 repassar a mensagem
	// 	fmt.Println(msg1)

	// 	msg2 := <-canal2 // Para e aguarda canal2 repassar a mensagem
	// 	fmt.Println(msg2)
	// }

	for {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
