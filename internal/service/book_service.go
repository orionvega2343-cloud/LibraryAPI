package service

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(book *repository.BookRepository) *BookService {
	return &BookService{repo: book}
}

func (s *BookService) Create(book *model.Book) (*model.Book, error) {
	err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) GetById(id int) (*model.Book, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
