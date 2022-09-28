package main

import (
	"fmt"
	"time"
)

func main() {
	go greeting("mundo")
	greeting("GoLang")
}

func greeting(nome string) {
	for {
		fmt.Printf("Olá, %s\n", nome)
		time.Sleep(time.Second / 2)
	}
}
