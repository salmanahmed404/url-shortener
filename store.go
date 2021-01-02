package main

import "github.com/jinzhu/gorm"

//URL is the model for storage
type URL struct {
	gorm.Model
	LongURL string
}
