package main

import (
	"fmt"
)

var k = "olá mundo" // Correto
// w := "vai teia"  // Errado

var a int
var b float64
var c string
var d bool

func parte1() {
	var x float32 = 16.0
	y := "texto" // Declaração de variável
	z := true

	nbytes, _ := fmt.Println("Hello World!", "Fala, Piãozada!")
	fmt.Println(nbytes)

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", y, y)
	fmt.Printf("x: %v, %T\n", z, z)

	x = 30.5 // Atribuição de um novo valor
	fmt.Printf("x: %v, %T\n", x, x)

	fmt.Println("6%8 = ", 6%8)
}

type hotdog int

func parte3() {
	var bom hotdog = 3
	var x int = 10

	fmt.Printf("\n%v, %T\n", bom, bom)
	fmt.Printf("%v, %T\n", x, x)

	// bom = x //  Não funciona (hotdog != int)
	x = int(bom)
	fmt.Println(x)

	x = 15
	bom = hotdog(x)
	fmt.Println(bom)
}

func parte2() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	s := "\nOi bom dia\ncomo vai?\nespero \"que\" tudo bem\n"
	fmt.Println(s)

	s1 := `Oi bom dia
	como vai?
	espero \"que\" tudo bem`
	fmt.Println(s1)

	var x float32 = 3.14
	cat := fmt.Sprint("\nO valor de x é ", x)
	fmt.Println(cat)
}

func main() {
	parte1()
	parte2()
	parte3()
}
