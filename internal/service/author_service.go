package service

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(a *model.Author) (*model.Author, error) {
	err := s.repo.Create(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AuthorService) GetById(id int) (*model.Author, error) {
	res, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
