package handler

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	s *service.UserService
}

func NewHandler(s *service.UserService) *Handler {
	return &Handler{s: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := h.s.Register(&user)
	if res != nil {
		http.Error(w, res.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
