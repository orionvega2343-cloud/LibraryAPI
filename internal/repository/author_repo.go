package repository

import (
	"LibrariAPI/internal/model"

	"github.com/jmoiron/sqlx"
)

// Struct
type AuthorRepository struct {
	DB *sqlx.DB
}

// Constructor
func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

// Create()
func (repo *AuthorRepository) Create(a *model.Author) error {
	err := repo.DB.QueryRow(`INSERT INTO authors (name, surname, bio) VALUES ($1, $2, $3) RETURNING id`, a.Name, a.Surname, a.Bio).Scan(&a.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll()
func (repo *AuthorRepository) GetById(id int) (*model.Author, error) {
	var a model.Author
	err := repo.DB.QueryRow(`SELECT name, surname, bio  FROM authors WHERE id = $1`, id).Scan(&a.Name, &a.Surname, &a.Bio)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
