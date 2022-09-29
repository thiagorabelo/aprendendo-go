package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodoviária"}
	enderecoEmMinuscilo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmMinuscilo, " ")[0]

	valido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			valido = true
			break
		}
	}

	if valido {
		return strings.Title(primeiraPalavra)
	}

	return "Inválido"
}
