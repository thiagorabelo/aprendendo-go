package main

import "fmt"

func closure(extra string) func() {
	texto := "Dentro da função closure"

	return func() {
		fmt.Printf("%s [%s]\n", texto, extra)
	}
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	fn := closure("Texto extra")
	fn()
}
