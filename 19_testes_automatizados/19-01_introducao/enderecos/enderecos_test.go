// TESTE DE UNIDADE
package enderecos

/* ou */
// package enderecos_test

import (
	"testing"
)

/*
import (
	...

	"introducao-testes/enderecos"    // enderecos.TipoEndereco(...)
	. "introducao-testes/enderecos"  // TipoEndereco(...)
)
*/

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoEndereco(t *testing.T) {
	/* Executa os testes em paralelo */
	/* OBS: Precisa ser colocado em todos os testes que devem ser paralelizados */
	// t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodoviária dos Imigrantes", "Rodoviária"},
		{"Estrada Qualquer", "Estrada"},
		{"Praça das Rosas", "Inválido"}, // comente aqui e faça o teste com a opção --cover
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Inválido"}, // comente aqui e faça o teste com a opção --cover
	}

	for _, cenario := range cenariosDeTeste {
		recebido := TipoEndereco(cenario.enderecoInserido)
		if recebido != cenario.enderecoEsperado {
			t.Errorf(
				"O tipo de endereço recebido (%s) é diferente do esperado (%s)",
				recebido,
				cenario.enderecoEsperado,
			)
		}
	}
}
