package rotas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasUsuarios...)

	for _, rota := range rotas {
		var handler http.HandlerFunc

		if rota.RequerAutenticacao {
			handler = middlewares.Autenticar(rota.Funcao)
		}

		// Adicionar outros middlewares

		// Default - Sem outros middlewares
		if handler == nil {
			handler = rota.Funcao
		}

		router.HandleFunc(
			rota.URI,
			middlewares.Logger(handler), // Middleware de Logger é sempre aplicado
		).Methods(rota.Metodo)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Não foi possível obter o diretório de trabalho atual", err)
	}

	assetsPath := fmt.Sprintf("%s/assets", cwd)
	fileServer := http.FileServer(http.Dir(assetsPath))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	return router
}
