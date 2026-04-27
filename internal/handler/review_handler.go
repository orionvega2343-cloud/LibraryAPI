package handler

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ReviewHandler struct {
	rev *service.ReviewService
}

func NewReviewHandler(r *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{rev: r}
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var review model.Reviews
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.rev.Create(&review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *ReviewHandler) GetReviewsById(w http.ResponseWriter, r *http.Request) {
	getId := chi.URLParam(r, "id")
	parseId, err := strconv.Atoi(getId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.rev.FindById(parseId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
