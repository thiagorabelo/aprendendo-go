package main

import "fmt"

func aritimeticos() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10.0 / 4.0
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var num1 int16 = 10
	var num2 int32 = 25
	// soma2 := num1 + num2 // Não funciona a operação de tipos diferentes nem há conversão implícita.
	soma2 := int32(num1) + num2
	fmt.Println(soma2)

	fmt.Printf("%v, %T\n", num1, num1)
	num1++     // num = num + 1
	num1 += 10 // num = num + 10
	fmt.Printf("%v, %T\n", num1, num1)

	// ++num1 // Não existe na linguagem (préfixado)
}

func atribuicao() {
	var var1 string = "Atribuição com declaração explícita"
	var2 := "Atribuição com inferência de tipo"
	fmt.Println(var1, "\n", var2)

	var1 = "Atrinuição para mudar valor de variável"
	fmt.Println(var1)
}

func relacionais() {
	fmt.Printf("1  < 2 = %v\n", 1 < 2)
	fmt.Printf("1 <= 2 = %v\n", 1 <= 2)
	fmt.Printf("2 == 2 = %v\n", 2 == 2)
	fmt.Printf("3 >= 2 = %v\n", 3 >= 2)
	fmt.Printf("3  > 2 = %v\n", 3 > 2)
	fmt.Printf("3 != 2 = %v\n", 3 != 2)
}

func logicos() {
	v, f := true, false

	fmt.Printf("v && f = %v\n", v && f)
	fmt.Printf("v || f = %v\n", v || f)
	fmt.Printf("!v = %v\n", !v)
	fmt.Printf("!f = %v\n", !f)
}

func bitwise() {
	var num int32 = 1
	fmt.Printf("%v, %T\n", num, num)
	num <<= 1
	fmt.Printf("%v, %T\n", num, num)
	num >>= 1
	fmt.Printf("%v, %T\n", num, num)
}

func main() {
	fmt.Println("# Aritimeticos")
	aritimeticos()

	fmt.Println("\n# Atribuição")
	atribuicao()

	fmt.Println("\n# Relacionais")
	relacionais()

	fmt.Println("\n# Lógicos")
	logicos()

	fmt.Println("\n# Bitwise")
	bitwise()
}
