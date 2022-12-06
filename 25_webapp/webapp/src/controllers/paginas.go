package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	if cookie, _ := cookies.Ler(r); cookie != nil && cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, "/publicacoes", nil)
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

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioId   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioId:   usuarioId,
	})
}

func CarregarTelaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
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

func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/usuarios?usuario=%s", url.PathEscape(nomeOuNick)),
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

	var usuarios []modelos.Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuarios); err != nil {
		respostas.InformaErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.InformaErro(w, http.StatusBadRequest, err)
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioId == usuarioLogadoId {
		http.Redirect(w, r, "/perfil", http.StatusFound)
		return
	}

	usuario, err := modelos.BuscarUsuarioCompleto(usuarioId, r)
	if err != nil {
		respostas.InformaErro(w, http.StatusInternalServerError, err)
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         modelos.Usuario
		UsuarioLogadoId uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoId: usuarioLogadoId,
	})
}

func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)
}
