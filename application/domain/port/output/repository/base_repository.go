package repository

import "github.com/fabianogoes/dev-books-api/application/domain/entity"

type BaseRepository interface {
	Save(b *entity.Book) error
	List() []*entity.Book
	FindById(id string) *entity.Book
	Update(b *entity.Book) error
	Delete(id string) error
}
