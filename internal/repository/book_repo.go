package repository

import (
	"LibrariAPI/internal/model"

	"github.com/jmoiron/sqlx"
)

// Struct
type BookRepository struct {
	DB *sqlx.DB
}

// Constructor
func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// Create()
func (repo *BookRepository) Create(book *model.Book) error {
	err := repo.DB.QueryRow(`INSERT INTO books (title,author_id,year) VALUES ($1, $2, $3) RETURNING id`, book.Title, book.AuthorID, book.Year).Scan(&book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll()
func (repo *BookRepository) GetByID(id int) (*model.Book, error) {
	var book model.Book
	err := repo.DB.QueryRow(`SELECT id,title,author_id,year FROM books WHERE id = $1`, id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
