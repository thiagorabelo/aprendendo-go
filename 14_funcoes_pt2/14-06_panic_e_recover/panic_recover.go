package main

import "fmt"

/*
 * Panic é uma função interna que interrompe o fluxo normal de controle e começa a entrar em pânico.
 * Quando a função F chama panic, a execução de F é interrompida, quaisquer funções adiadas em F são
 * executadas normalmente e então F retorna ao seu chamador. Para o chamador, F então se comporta
 * como um chamada à panic. O processo continua na pilha até que todas as funções na goroutine atual
 * tenham retornado, momento em que o programa trava. Pânicos podem ser iniciados invocando o panic
 * diretamente. Eles também podem ser causados ​​por erros de tempo de execução, como acessos de matriz
 * fora dos limites.
 *
 * Recover é uma função interna que permite ao programa gerenciar o comportamento de uma goroutine em
 * pânico. Executar uma chamada para recover dentro de uma função adiada (mas não qualquer função
 * chamada por ela) interrompe a sequência de pânico restaurando a execução normal e recupera o valor
 * de erro passado para a chamada de panic. Se a recover for chamada fora da função 'defer', ela
 * não interromperá uma sequência em pânico. Neste caso, ou quando a goroutine não está em panico, ou
 * se o argumento fornecido ao panic foi nil, recover retorna nil. Assim, o valor de retorno de
 * recover informa se a goroutine está em pânico.
 *
 * - <https://go.dev/blog/defer-panic-and-recover>
 */

func main() {
	f()
	fmt.Println("Retornado normalmente de f().")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado em f().", r)
		}
	}()

	fmt.Println("Chamando g(0).")
	g(0)
	fmt.Println("Retornado normalmente de g(0).")
}

func g(i int) {
	if i > 3 {
		fmt.Printf("Estou entrando em Pânico em g(%v)!\n", i)
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Printf("'Defer' em g(%v)\n", i)
	fmt.Printf("Imprimindo em g(%v)\n", i)
	g(i + 1)
}
