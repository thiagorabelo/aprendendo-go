package main

import (
	"fmt"
)

var x = 10   // default é int (total de bits depende da arquitetura do processador. 32, 64, etc)
var y = 10.5 // default é float64

func main() {
	// x = 10.5  // Erro de compilação
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

}
