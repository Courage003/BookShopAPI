package models

import "gorm.io/gorm"

type Books struct{
	gorm.Model
	Author *string `json: "author"`
	Title *string `json: "title"`
	Publisher *string `json: "publisher"`
}