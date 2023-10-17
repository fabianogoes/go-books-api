package entity

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID
	Title       string
	Description string
	Author      string
}

func NewBook(title string, description string, author string) *Book {
	return &Book{ID: uuid.New(), Title: title, Description: description, Author: author}
}
