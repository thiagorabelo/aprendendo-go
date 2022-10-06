package modelos

import (
	"errors"
	"strings"
	"time"
)

type Etapa uint64

const (
	Cadastro Etapa = iota
	Edicao
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"` // omitempty - Remove o campo do json se for valor zero
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa Etapa) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa Etapa) error {
	if usuario.Nome == "" {
		return errors.New("o campo nome é obrigatório")
	}
	if usuario.Nick == "" {
		return errors.New("o campo nick é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("o campo email é obrigatório")
	}
	if etapa == Cadastro && usuario.Senha == "" {
		return errors.New("o campo senha é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
