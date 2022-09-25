package main

import (
	"fmt"
)

const valor int = 100

func main() {
	if x := 100; x > valor {
		fmt.Printf("Xís (%d) é maior que %d\n", x, valor)
	} else if x < valor {
		fmt.Printf("Xís (%d) é menor que %d\n", x, valor)
	} else {
		fmt.Printf("Xís (%d) é igual a %d\n", x, valor)
	}
}
