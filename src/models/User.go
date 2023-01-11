package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

func (user *User) Prepare(etapa string) error {
	if erro := user.validate(etapa); erro != nil {
		return erro
	}

	if erro := user.format(etapa); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(etapa string) error {
	if user.Name == "" {
		return errors.New("O nome e obrigatorio")
	}

	if user.Nick == "" {
		return errors.New("O nick e obrigatorio")
	}

	if user.Email == "" {
		return errors.New("O email e obrigatorio")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido e invalido")
	}

	if user.Password == "" && etapa == "Cadastro" {
		return errors.New("A senha e obrigatoria")
	}

	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if etapa == "cadastro" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHash)
	}

	return nil
}
