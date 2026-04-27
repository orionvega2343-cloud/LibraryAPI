package handler

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorHandler struct {
	a *service.AuthorService
}

func NewAuthorHandler(a *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{a: a}
}

func (h *AuthorHandler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.a.Create(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *AuthorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	getId := chi.URLParam(r, "id")
	convId, err := strconv.Atoi(getId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.a.GetById(convId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
