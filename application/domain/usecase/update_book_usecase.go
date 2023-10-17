package usecase

import (
	"fmt"

	"github.com/fabianogoes/dev-books-api/application/domain/port/output/repository"
	uuid "github.com/satori/go.uuid"
)

type UpdateBookUseCase struct {
	repository repository.BaseRepository
}

func NewUpdateBookUseCase(r repository.BaseRepository) *UpdateBookUseCase {
	return &UpdateBookUseCase{repository: r}
}

func (buc *UpdateBookUseCase) Update(id uuid.UUID, title string, description string, author string) error {
	book := buc.repository.FindById(id.String())
	if book == nil {
		return fmt.Errorf(fmt.Sprintf("book %s not found", title))
	}

	book.Title = title
	book.Description = description
	book.Author = author
	return buc.repository.Update(book)
}
