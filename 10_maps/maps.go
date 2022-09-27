package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//             key    value
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	funcionario := map[string]map[string]string{
		"dados_pessoais": {
			"nome":      "João",
			"sobrenome": "Costa",
		},
		"dados_funcionais": {
			"matricula": "20221027",
			"setor":     "Manutenção Predial",
		},
	}
	fmt.Println(funcionario)

	delete(funcionario, "dados_funcionais")
	fmt.Println(funcionario)

	delete(funcionario["dados_pessoais"], "sobrenome")
	fmt.Println(funcionario)

	funcionario["dados_pessoais"]["sobrenome"] = "Melo"
	funcionario["dados_funcionais"] = map[string]string{
		"filial": "Sul",
	}
	fmt.Println(funcionario)
}
