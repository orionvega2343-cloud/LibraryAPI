package repository

import (
	"LibrariAPI/internal/model"

	"github.com/jmoiron/sqlx"
)

// Struct
type ReviewRepository struct {
	DB *sqlx.DB
}

// Constructor
func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

// Create
func (repo *ReviewRepository) Create(r *model.Reviews) error {
	err := repo.DB.QueryRow(`INSERT INTO reviews (user_id,book_id,text,grade) VALUES ($1,$2,$3,$4) RETURNING id`, r.UserID, r.BookID, r.Text, r.Grade).Scan(&r.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetByID

func (repo *ReviewRepository) GetByID(id int) (*model.Reviews, error) {
	var r model.Reviews
	err := repo.DB.QueryRow(`SELECT id,user_id,book_id,text,grade FROM reviews WHERE id = $1 `, id).Scan(&r.ID, &r.UserID, &r.BookID, &r.Text, &r.Grade)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
