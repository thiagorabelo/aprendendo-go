package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	response, err := http.Post(config.API("/usuarios"), "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodPost,
		fmt.Sprintf("/usuarios/%d/parar-de-seguir", usuarioId),
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

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodPost,
		fmt.Sprintf("/usuarios/%d/seguir", usuarioId),
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

func CarregaPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, err := modelos.BuscarUsuarioCompleto(usuarioId, r)
	if err != nil {
		respostas.InformaErro(w, http.StatusInternalServerError, err)
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}

func CarregarPaginaDeEdicaoDoUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, err := modelos.RecuperarDadosDoUsuarioViaAPI(usuarioId, r)
	if err != nil {
		respostas.InformaErro(w, http.StatusInternalServerError, err)
		return
	}

	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodPut,
		fmt.Sprintf("/usuarios/%d", usuarioId),
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

	respostas.JSON(w, response.StatusCode, nil)
}
