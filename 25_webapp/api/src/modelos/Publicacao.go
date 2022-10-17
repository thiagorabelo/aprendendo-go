package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	Id        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o campo título é obrigatorio")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o campo conteúdo é obrigatorio")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
