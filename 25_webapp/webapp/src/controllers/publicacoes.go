package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(
		r,
		http.MethodPost,
		"/publicacoes",
		bytes.NewBuffer(publicacao),
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

	respostas.JSON(w, response.StatusCode, nil)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(
		r,
		http.MethodPost,
		fmt.Sprintf("/publicacoes/%d/curtir", publicacaoId),
		nil,
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

	respostas.JSON(w, response.StatusCode, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(
		r,
		http.MethodPost,
		fmt.Sprintf("/publicacoes/%d/descurtir", publicacaoId),
		nil,
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

	respostas.JSON(w, response.StatusCode, nil)
}

func CarregarTelaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/publicacoes/%d", publicacaoId),
		nil,
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

	var publicacao modelos.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacao); err != nil {
		respostas.InformaErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	r.ParseForm()
	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoCOmAutenticacao(
		r,
		http.MethodPut,
		fmt.Sprintf("/publicacoes/%d", publicacaoId),
		bytes.NewBuffer(publicacao),
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

	respostas.JSON(w, response.StatusCode, nil)
}
