package main

import (
	"fmt"
)

func main() {
	var x uint16
	x = 65535
	// x = 65536  # vai dar pau
	fmt.Printf("%v, %T\n", x, x)

	x++ // volta a zero
	fmt.Printf("%v, %T\n", x, x)

	x++ // 1
	fmt.Printf("%v, %T\n", x, x)
}
