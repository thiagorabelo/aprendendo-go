package modelos

import "time"

type Usuario struct {
	Id       uint64    `json:"id,omitempty"` // omitempty - Remove o campo do json se for valor zero
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}
