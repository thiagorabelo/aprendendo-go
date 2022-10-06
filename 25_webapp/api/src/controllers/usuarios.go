package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("Erro ao ler o corpo da requisição"))
		log.Fatal(err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	id, err := repositorio.Criar(usuario)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", id)))
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar Usuários"))
}

func MostrarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostrando Usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}
