package api

import (
	"encoding/json"
	"github.com/BigGold1310/spend-smart-save/internal/domain"
	"github.com/BigGold1310/spend-smart-save/internal/service"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

func sendJSON(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, err error, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, req bunrouter.Request) error {
	users, err := h.userService.ListUsers()
	if err != nil {
		return sendError(w, err, http.StatusInternalServerError)
	}
	return sendJSON(w, users, http.StatusOK)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, req bunrouter.Request) error {
	id, err := strconv.Atoi(req.Param("id"))
	if err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		return sendError(w, err, http.StatusNotFound)
	}
	return sendJSON(w, user, http.StatusOK)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, req bunrouter.Request) error {
	var createReq domain.UserCreate
	if err := json.NewDecoder(req.Body).Decode(&createReq); err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}

	user, err := h.userService.CreateUser(&createReq)
	if err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}
	return sendJSON(w, user, http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, req bunrouter.Request) error {
	id, err := strconv.Atoi(req.Param("id"))
	if err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}

	var updateReq domain.UserUpdate
	if err := json.NewDecoder(req.Body).Decode(&updateReq); err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}

	user, err := h.userService.UpdateUser(id, &updateReq)
	if err != nil {
		return sendError(w, err, http.StatusBadRequest)
	}
	return sendJSON(w, user, http.StatusOK)
}
