package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10

	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Println("len de x é", len(x))
	fmt.Println("len de y é", len(y))

	fmt.Println()

	fmt.Printf("x é um %T\n", x)
	fmt.Printf("y é um %T\n", y)

	y[0] = x[1]

	fmt.Println("y é", y)
}
