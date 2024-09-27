package api

import (
	"encoding/json"
	"github.com/BigGold1310/spend-smart-save/internal/repository"
	"github.com/uptrace/bunrouter"
	"net/http"
)

type UserHandler struct {
	U repository.User
}

func (u UserHandler) Get(w http.ResponseWriter, req bunrouter.Request) error {
	users, err := u.U.GetAll()
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return err
	}
	return nil
}

func (u UserHandler) Create(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (u UserHandler) Update(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (u UserHandler) Delete(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}
