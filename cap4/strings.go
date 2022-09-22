package main

import (
	"fmt"
)

func stringLength() {
	sample := "ab£c" // 4 chars

	fmt.Println(len(sample))    // 5
	fmt.Println([]byte(sample)) // [97 98 194 163 99]

	s := []rune(sample)
	fmt.Println(len(s))

	len := 0
	for range sample {
		len++
	}
	fmt.Println(len)
}

/**
 * %b, %v, %T
 * Raw strings literals
 * Conversão para slices of bytes: []byte(x)
 * %#U, %#x
 */

func main() {
	s := `Hello
 Go
 Lang (éøâ 香)`

	for _, v := range s { // loop sobre os caracteres
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ { // loop sobre os bytes
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}

	fmt.Println("")

	stringLength()
}
