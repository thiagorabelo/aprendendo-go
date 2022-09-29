package formas

import (
	"math"
	"testing"
)

/*
 * - <https://dev.to/juliaferraioli/testing-in-go-testing-floating-point-numbers-4i0a>
 * - <https://pkg.go.dev/github.com/google/go-cmp/cmp>
 */

func TestArea(t *testing.T) {
	const tolerance = .00001
	comp := func(a, b float64) bool {
		diff := math.Abs(a - b)
		mean := math.Abs(a+b) / 2.0
		return (diff / mean) < tolerance
	}

	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		esperado := float64(120)
		recebido := ret.Area()

		if !comp(esperado, recebido) {
			t.Fatalf("Recebido (%f) é diferente do esperado (%f)", recebido, esperado) // Força parada
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		const raio = 10
		circ := Circulo{raio}
		esperado := float64(math.Pi * math.Pow(raio, 2))
		recebido := circ.Area()

		if !comp(esperado, recebido) {
			t.Fatalf("Recebido (%f) é diferente do esperado (%f)", recebido, esperado) // Força parada
		}
	})
}
