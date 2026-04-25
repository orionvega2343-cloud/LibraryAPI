package service

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(review *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: review}
}

func (s *ReviewService) Create(review *model.Reviews) error {
	err := s.repo.Create(review)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewService) FindById(id int) (*model.Reviews, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
