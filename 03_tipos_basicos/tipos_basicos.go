package main

import (
	"errors"
	"fmt"
)

func inteiros() {
	// (u)int8, (u)int16, (u)int32, (u)int64 :: 8, 16, 32, 64 bits respectivamente
	// int :: tamanho padrão da palavra da arquitetura (64 bits em intel ou amd 64)

	var ubits8 uint8 = 255 // 0 - 255. Maior que 255 causa problemas de overflow.
	fmt.Printf("%v, %T\n", ubits8, ubits8)

	var bits8 int8 = -128 // -128 - 127.
	fmt.Printf("%v, %T\n", bits8, bits8)

	// alias
	// int32 = rune
	var letra rune = 97 // 'a' em unicode
	fmt.Printf("%c, %T\n", letra, letra)

	// alias
	// uint8 = byte
	var bt byte = 97 // 'a' em ascii
	fmt.Printf("%c, %T\n", bt, bt)
}

func numerosReais() {
	var real32 float32 = 3.14
	fmt.Printf("%v, %T\n", real32, real32)

	var real64 float64 = 3.14
	fmt.Printf("%v, %T\n", real64, real64)
}

func booleanos() {
	var b1 bool = true
	fmt.Printf("%v, %T\n", b1, b1)
}

func strings() {
	var str string = "Texto"
	fmt.Printf("%v, %T\n", str, str)

	char := 'á' // int32
	fmt.Printf("%v, %c, %T\n", char, char, char)
}

func erros() {
	var erro error = errors.New("Meu Erro")
	fmt.Printf("%v, %T\n", erro, erro)
}

func valoresZeros() {
	// Todo tipo tem seu valor zero de inicialização
	// para evitar comportamento inesperado.

	var a int32
	fmt.Printf("%v, %T\n", a, a)

	var b float32
	fmt.Printf("%v, %T\n", b, b)

	var c bool
	fmt.Printf("%v, %T\n", c, c)

	var d string
	fmt.Printf("%v, %T\n", d, d)

	var e error
	fmt.Printf("%v, %T\n", e, e)
}

func main() {

	fmt.Println("# Inteiros")
	inteiros()

	fmt.Println("\n# Números reais")
	numerosReais()

	fmt.Println("\n# Booleanos")
	booleanos()

	fmt.Println("\n# Strings")
	strings()

	fmt.Println("\n# Erros")
	erros()

	fmt.Println("\n# Valores Zeros")
	valoresZeros()
}
