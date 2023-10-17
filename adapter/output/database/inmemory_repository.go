package database

import "github.com/fabianogoes/dev-books-api/application/domain/entity"

type inMemoryRepository struct {
	db []*entity.Book
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{db: make([]*entity.Book, 0)}
}

func (imr *inMemoryRepository) Save(b *entity.Book) error {
	imr.db = append(imr.db, b)
	return nil
}

func (imr *inMemoryRepository) List() []*entity.Book {
	return imr.db
}

func (imr *inMemoryRepository) FindById(id string) *entity.Book {
	for _, b := range imr.db {
		if b.ID.String() == id {
			return b
		}
	}
	return nil
}

func (imr *inMemoryRepository) Update(b *entity.Book) error {
	for i, book := range imr.db {
		if book.ID == b.ID {
			imr.db[i] = b
		}
	}
	return nil
}

func (imr *inMemoryRepository) Delete(id string) error {
	for i, b := range imr.db {
		if b.ID.String() == id {
			imr.db = append(imr.db[:i], imr.db[i+1:]...)
		}
	}
	return nil
}
