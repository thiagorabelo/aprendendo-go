package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		greeting("mundo")
		waitGroup.Done()
	}()

	go func() {
		greeting("GoLang")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func greeting(nome string) {
	for {
		fmt.Printf("Olá, %s\n", nome)
		time.Sleep(time.Second / 2)
	}
}
