package models

import "gorm.io/gorm"

type Books struct{
	gorm.Model
	Author *string `json: "author"`
	Title *string `json: "title"`
	Publisher *string `json: "publisher"`
}


//creating a function for migration
func MigrateBooks(db *gorm.DB) error{
	err:= db.AutoMigrate(&Books{})
	return err
}