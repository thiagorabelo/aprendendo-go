package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Evita erro quando o status code é "No Content (204)",
	// pois mesmo nil é tratado como "conteúdo".
	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: err.Error(),
	})
}
