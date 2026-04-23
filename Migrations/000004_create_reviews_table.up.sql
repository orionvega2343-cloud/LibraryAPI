CREATE TABLE IF NOT EXISTS reviews(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    book_id INT NOT NULL REFERENCES books(id),
    text TEXT NOT NULL,
    grade INT NOT NULL
);