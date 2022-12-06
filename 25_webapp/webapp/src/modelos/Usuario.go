package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/requisicoes"
)

type Usuario struct {
	Id          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguidos := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioId, r)
	go BuscarSeguidoresDoUsuario(canalSeguidores, usuarioId, r)
	go BuscarSeguidosDoUsuario(canalSeguidos, usuarioId, r)
	go BuscarPublicacoesDoUsuario(canalPublicacoes, usuarioId, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguidos    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.Id == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar os seguidores do usuário")
			}

			seguidores = seguidoresCarregados

		case seguidosCarregados := <-canalSeguidos:
			if seguidosCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar os seguidos do usuário")
			}

			seguidos = seguidosCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar as publicações do usuário")
			}

			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguidos
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

func RecuperarDadosDoUsuarioViaAPI(usuarioId uint64, r *http.Request) (Usuario, error) {
	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/usuarios/%d", usuarioId),
		nil,
	)
	if err != nil {
		return Usuario{}, err
	}
	defer response.Body.Close()

	var usuario Usuario
	if err := json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}

func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	usuario, err := RecuperarDadosDoUsuarioViaAPI(usuarioId, r)
	if err != nil {
		canal <- Usuario{}
		return
	}
	canal <- usuario
}

func BuscarSeguidoresDoUsuario(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/usuarios/%d/seguidores", usuarioId),
		nil,
	)
	if err != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if err := json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

func BuscarSeguidosDoUsuario(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/usuarios/%d/seguindo", usuarioId),
		nil,
	)
	if err != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidos []Usuario
	if err := json.NewDecoder(response.Body).Decode(&seguidos); err != nil {
		canal <- nil
		return
	}

	if seguidos == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidos
}

func BuscarPublicacoesDoUsuario(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodGet,
		fmt.Sprintf("/usuarios/%d/publicacoes", usuarioId),
		nil,
	)
	if err != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if err := json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
