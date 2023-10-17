package entity

import uuid "github.com/satori/go.uuid"

type Book struct {
	ID          uuid.UUID
	Title       string
	Description string
	Author      string
}

func NewBook(title string, description string, author string) *Book {
	return &Book{ID: uuid.NewV4(), Title: title, Description: description, Author: author}
}
