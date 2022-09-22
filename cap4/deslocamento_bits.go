package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb
	gb
)

func main() {
	x := 1
	y := x << 1
	z := x >> 1
	fmt.Println(x, y, z)

	fmt.Println()

	bt := 1
	// kb := 1 << 10
	// mb := 1 << 20
	// gb := 1 << 30
	fmt.Printf("%b\t\t\t\t%d\n", bt, bt)
	fmt.Printf("%b\t\t\t%d\n", kb, kb)
	fmt.Printf("%b\t\t%d\n", mb, mb)
	fmt.Printf("%b\t%d\n", gb, gb)
}
