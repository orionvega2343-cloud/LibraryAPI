package main

import (
	"LibrariAPI/internal/config"
	"LibrariAPI/internal/db"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected", db)

}
