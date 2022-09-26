package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v2++
	fmt.Println(v1, v2)

	var ptr *int
	fmt.Println(ptr)

	ptr = &v2
	fmt.Println(ptr, *ptr)

	*ptr++
	fmt.Println(*ptr, v2)
}
