package repository

import (
	"LibrariAPI/internal/model"

	"github.com/jmoiron/sqlx"
)

// Struct
type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create
func (repo *UserRepository) Create(user *model.User) error {
	err := repo.DB.QueryRow(`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetById
func (repo *UserRepository) GetByID(id int) (*model.User, error) {
	user := &model.User{}
	err := repo.DB.QueryRow(`SELECT id, email, password FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (repo *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.DB.QueryRow(`SELECT id,email,password FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
