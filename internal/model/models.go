package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Author struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Bio     string `json:"bio"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Year     int    `json:"year"`
}

type Reviews struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Text   string `json:"text"`
	Grade  int    `json:"grade"`
}
