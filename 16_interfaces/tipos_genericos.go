package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// função genérica
	fmt.Println(1, 2, "string", false, true, float64(1234))

	// mapa genérico
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "Vish",
		true:         float64(12),
	}
	fmt.Println(mapa)
}
