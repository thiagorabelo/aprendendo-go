package main

import (
	"fmt"
)

func minDaUmNumero() int {
	return 2
}

func main() {

	if x := minDaUmNumero(); x == 2 || x%2 == 0 {
		fmt.Println("É dois ou múltiplo de 2")
	}
}
