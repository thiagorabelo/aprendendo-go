package main

import (
	"fmt"
	"runtime"
)

func main() {
	// x = 10.5  // Erro de compilação
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
