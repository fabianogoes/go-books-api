package usecase

import (
	"github.com/fabianogoes/dev-books-api/application/domain/entity"
	"github.com/fabianogoes/dev-books-api/application/domain/port/output/repository"
)

type ListBookUseCase struct {
	repository repository.BaseRepository
}

func NewListBookUseCase(r repository.BaseRepository) *ListBookUseCase {
	return &ListBookUseCase{repository: r}
}

func (buc *ListBookUseCase) List() []*entity.Book {
	return buc.repository.List()
}
