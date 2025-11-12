package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qobilovvv/1uchet/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.service.Create(req.PhoneNumber, req.Password)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, user)
}
