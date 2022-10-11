package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := append(
		rotasUsuarios,
		rotaLogin,
	)

	for _, rota := range rotas {
		var handler http.HandlerFunc = nil

		if rota.RequerAutenticacao {
			handler = middlewares.Autenticar(rota.Funcao)
		}

		// Adicionar outros middlewares

		// Default - Sem outros middleware
		if handler == nil {
			handler = rota.Funcao
		}

		router.HandleFunc(
			rota.URI,
			middlewares.Logger(handler), // Middleware de Logger é sempre aplicado
		).Methods(rota.Metodo)
	}

	return router
}
