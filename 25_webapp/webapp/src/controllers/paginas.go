package controllers

import (
	"encoding/json"
	"net/http"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(r, http.MethodGet, "/publicacoes", nil)
	if err != nil {
		respostas.InformaErro(w, http.StatusInternalServerError, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if err := json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		respostas.InformaErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.ExecutarTemplate(w, "home.html", publicacoes)
}
