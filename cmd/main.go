package main

import (
	"fmt"
	"movies/internal/config"
)

func main() {
	db := config.SetDatabaseConnection()
	fmt.Println(db) // временно
}
