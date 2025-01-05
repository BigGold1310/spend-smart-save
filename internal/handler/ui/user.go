package ui

import (
	"net/http"
	"strconv"

	"github.com/BigGold1310/spend-smart-save/internal/repository"
	"github.com/BigGold1310/spend-smart-save/internal/template"
	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, req bunrouter.Request) error {
	users, err := h.userRepo.GetUsers()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"Users": users,
		"Title": "User List",
	}

	return template.Render(w, "user/list.gohtml", data)
}

func (h *UserHandler) ShowUser(w http.ResponseWriter, req bunrouter.Request) error {
	idStr := req.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	user, err := h.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"User":  user,
		"Title": "User Details",
	}

	return template.Render(w, "user/detail.gohtml", data)
}

func (h *UserHandler) ShowCreateForm(w http.ResponseWriter, req bunrouter.Request) error {
	data := map[string]interface{}{
		"Title": "Create User",
	}
	return template.Render(w, "user/create.gohtml", data)
}
