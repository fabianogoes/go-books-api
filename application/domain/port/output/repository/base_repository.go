package repository

import (
	"github.com/fabianogoes/dev-books-api/application/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type BaseRepository interface {
	Save(b *entity.Book) error
	List() []*entity.Book
	FindById(id uuid.UUID) *entity.Book
	Update(b *entity.Book) error
	Delete(id uuid.UUID) error
}
