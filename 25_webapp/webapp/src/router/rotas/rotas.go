package rotas

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	rotas = append(rotas, rotasUsuarios...)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
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
