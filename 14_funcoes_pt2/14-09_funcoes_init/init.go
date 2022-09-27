package main

import "fmt"

func main() {
	fmt.Println("Dentro da função main")
}

// Pode haver uma função init por arquivo
func init() {
	fmt.Println("Dentro da função init (executo antes da main)")
}
