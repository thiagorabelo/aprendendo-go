package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func EscreveArea(f Forma) {
	fmt.Printf("A área do %T é %0.2f\n", f, f.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
