package usecase

import (
	"errors"
	"fmt"

	"github.com/fabianogoes/dev-books-api/application/domain/port/output/repository"
	uuid "github.com/satori/go.uuid"
)

type DeleteBookUseCase struct {
	repository repository.BaseRepository
}

func NewDeleteBookUseCase(r repository.BaseRepository) *DeleteBookUseCase {
	return &DeleteBookUseCase{repository: r}
}

func (buc *DeleteBookUseCase) Delete(id uuid.UUID) error {
	if book := buc.repository.FindById(id); book == nil {
		return errors.New(fmt.Sprintf("book %s not found", id))
	}

	return buc.repository.Delete(id)
}
