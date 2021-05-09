package modelos

import (
	"errors"
	"strings"
	"time"
)

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `json:"nome, omitempty"`
	Nick     string    `json:"nick, omitempty"`
	Email    string    `json:"email, omitempty"`
	Senha    string    `json:"senha, omitempty"`
	CriadoEm time.Time `json:"criadoEm, omitempty"`
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório, não pode ser em vazio!")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório, não pode ser em vazio!")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório, não pode ser em vazio!")
	}

	if usuario.Senha == "" {
		return errors.New("O senha é obrigatório, não pode ser em vazio!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
