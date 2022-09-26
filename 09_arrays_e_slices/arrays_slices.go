package main

import "fmt"

func arrays() {
	// Array é sempre estático
	var ar1 [5]int
	fmt.Println(len(ar1), ar1)
	ar1[0] = 5
	fmt.Println(ar1)

	ar2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(ar2), ar2)

	ar3 := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println(len(ar3), ar3)
}

func slices() {
	slc1 := []int{10, 11, 12, 13}
	fmt.Println(len(slc1), slc1)

	slc1[0] = 9
	fmt.Println(len(slc1), slc1)

	slc1 = append(slc1, 14)
	fmt.Println(len(slc1), slc1)

	slc2 := slc1[1:3]
	fmt.Println("slc2 := slc1[1:3] =>", slc2, "/ len(slc2) =", len(slc2))

	slc1[1] = 55
	fmt.Println(slc2)
}

func arraysInternos() {
	slc1 := make([]float32, 10, 11) // make(tipo, tamanho, capacidade_maxima)
	fmt.Println(slc1, "/", len(slc1), "/", cap(slc1))

	slc1 = append(slc1, 5.0)
	fmt.Println(slc1, "/", len(slc1), "/", cap(slc1))

	slc1 = append(slc1, 6.0)
	fmt.Println(slc1, "/", len(slc1), "/", cap(slc1))

	slc2 := []int{1, 2}
	fmt.Println(slc2, "/", len(slc2), "/", cap(slc2))
	slc3 := slc2[0:1]
	fmt.Println(slc3, "/", len(slc3), "/", cap(slc3))
	slc2 = append(slc2, 3, 4, 5)
	slc2[0] = 100
	fmt.Println(slc2, "/", len(slc2), "/", cap(slc2))
	fmt.Println(slc3, "/", len(slc3), "/", cap(slc3))
}

func main() {
	fmt.Println("# Arrays")
	arrays()

	fmt.Println("\n# Slices")
	slices()

	fmt.Println("\n# Arrays Internos")
	arraysInternos()
}
