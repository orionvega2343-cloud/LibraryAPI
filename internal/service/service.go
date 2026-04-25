package service

import (
	"LibrariAPI/internal/model"
	"LibrariAPI/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   *repository.UserRepository
	secret string
}

func NewUserService(repo *repository.UserRepository, secret string) *UserService {
	return &UserService{repo: repo, secret: secret}
}

type Claims struct {
	UserID int
	Email  string
	jwt.RegisteredClaims
}

func (u *UserService) HashPassword(p string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (u *UserService) Register(user *model.User) error {
	hash, err := u.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = u.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(user *model.User, email string) (string, error) {
	res, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	var c Claims
	c.UserID = res.ID
	c.Email = email
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
