package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db gorm.DB

type url struct {
	gorm.Model
	longURL  string
	shortURL string
}

func dbInit() {
	db, err := gorm.Open(sqlite.Open("url-shortener.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("DB connection failed")
	}

	db.AutoMigrate(&url{})
}
