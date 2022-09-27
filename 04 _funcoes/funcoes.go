package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1 int32, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Printf("%v, %T\n", soma, soma)

	var f = func(i int32) int32 {
		// fmt.Printf("Eu sou uma função f(%d)\n", i)
		return i
	}
	fmt.Printf("'%T', %v\n", f, f(10)) // 'func(int32) int32', 10

	calcSoma, calSub := calculosMatematicos(896, 128)
	fmt.Printf("A soma é %d e a subtração é %d\n", calcSoma, calSub)

	resultadoSoma, _ := calculosMatematicos(768, 256)
	fmt.Printf("A soma é %d\n", resultadoSoma)
}
