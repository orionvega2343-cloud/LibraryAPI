package main

import (
	"LibrariAPI/internal/config"
	"fmt"
)

func main() {
	fmt.Println(config.MustLoad())
}
