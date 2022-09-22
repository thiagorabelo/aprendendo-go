package main

import (
	"fmt"
)

/*********************/
const x = 10 // o tipo fica indefinido até o momento do uso (int, float, etc)
var y = 10   // definido no momento da declaração

var c int
var d float64

/*********************/

/*********************/
const ( // declaração múltipla
	i = 10
	j = 20
	k = 30
)

/*********************/

func main() {
	/*********************/
	c = x // funciona de boas (x passa a ser int)
	// c = y // funciona de boas
	// d = y  // Não funciona
	d = x // funciona de boas (x passa a ser float64)
	fmt.Println(c, d)
	/*********************/

	/*********************/
	fmt.Println(i, j, k)
	/*********************/
}
