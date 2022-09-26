package main

import (
	"fmt"
)

func main() {
	for hora := 0; hora < 12; hora++ {
		for minuto := 0; minuto < 60; minuto++ {
			fmt.Printf("%02d:%02d\n", hora, minuto)
		}
		fmt.Printf("\n\n\n")
	}
}
