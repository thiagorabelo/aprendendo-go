package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
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
