package main

import (
	"fmt"
)

// iota representa inteiros não tipados sucessivos
const (
	a = iota
	_ // = iota
	c // = iota
	x // = iota
	_ = iota + 10
	z // = iota
)

func main() {
	fmt.Println(a, c, x, z)
}
