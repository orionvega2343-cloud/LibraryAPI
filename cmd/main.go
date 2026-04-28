package main

import (
	"LibrariAPI/internal/config"
	"LibrariAPI/internal/db"
	"LibrariAPI/internal/handler"
	"LibrariAPI/internal/repository"
	"LibrariAPI/internal/service"
	"fmt"
	"net/http"

	mw "LibrariAPI/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.MustLoad()
	datb, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepository(datb)
	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	userHandler := handler.NewHandler(userService)

	authorRepo := repository.NewAuthorRepository(datb)
	authorService := service.NewAuthorService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	bookRepo := repository.NewBookRepository(datb)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	reviewsRepo := repository.NewReviewRepository(datb)
	reviewsService := service.NewReviewService(reviewsRepo)
	reviewsHandler := handler.NewReviewHandler(reviewsService)

	r := chi.NewRouter()
	r.Use(mw.LoggerMiddleware)
	r.Post("/auth/register", userHandler.CreateUser)
	r.Post("/auth/login", userHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return mw.AuthMiddleware(next, cfg.JWT.Secret)
		})
		r.Get("/authors/{id}", authorHandler.GetAuthorById)
		r.Post("/authors", authorHandler.AddAuthor)
		r.Get("/books/{id}", bookHandler.GetBookById)
		r.Post("/books", bookHandler.CreateBook)
		r.Get("/reviews/{id}", reviewsHandler.GetReviewsById)
		r.Post("/reviews", reviewsHandler.CreateReview)

	})
	fmt.Println("Сервер запщуен на порту 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
