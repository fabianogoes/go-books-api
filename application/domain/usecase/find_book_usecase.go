package usecase

import (
	"github.com/fabianogoes/dev-books-api/application/domain/entity"
	"github.com/fabianogoes/dev-books-api/application/domain/port/output/repository"
	uuid "github.com/satori/go.uuid"
)

type FindBookUseCase struct {
	repository repository.BaseRepository
}

func NewFindBookUseCase(r repository.BaseRepository) *FindBookUseCase {
	return &FindBookUseCase{repository: r}
}

func (buc *FindBookUseCase) FindById(id uuid.UUID) *entity.Book {
	return buc.repository.FindById(id)
}
