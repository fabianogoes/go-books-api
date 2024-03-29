package entity

import uuid "github.com/satori/go.uuid"

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
}

func NewBook(title string, description string, author string) *Book {
	return &Book{ID: uuid.NewV4(), Title: title, Description: description, Author: author}
}
