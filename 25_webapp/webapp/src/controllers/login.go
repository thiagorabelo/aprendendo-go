package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := http.Post(
		config.API("/login"),
		"application/json",
		bytes.NewBuffer(usuario),
	)
	if err != nil {
		respostas.InformaErro(w, http.StatusInternalServerError, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if err := json.NewDecoder(response.Body).Decode(&dadosAutenticacao); err != nil {
		respostas.InformaErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := cookies.Salvar(w, dadosAutenticacao.Id, dadosAutenticacao.Token); err != nil {
		respostas.InformaErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
