package main

import (
	"fmt"

	"github.com/t-okuji/go-chi-gorm-jwt/db"
	"github.com/t-okuji/go-chi-gorm-jwt/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
