package usecase

import (
	"github.com/fabianogoes/dev-books-api/application/domain/entity"
	"github.com/fabianogoes/dev-books-api/application/domain/port/output/repository"
)

type CreateBookUseCase struct {
	repository repository.BaseRepository
}

func NewCreateBookUseCase(r repository.BaseRepository) *CreateBookUseCase {
	return &CreateBookUseCase{repository: r}
}

func (buc *CreateBookUseCase) Create(title string, description string, author string) error {
	return buc.repository.Save(entity.NewBook(title, description, author))
}
