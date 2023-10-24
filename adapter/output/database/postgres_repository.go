package database

import (
	"github.com/fabianogoes/dev-books-api/application/domain/entity"
	"github.com/fabianogoes/dev-books-api/configuration/logger"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (pr *PostgresRepository) Save(b *entity.Book) error {
	if err := pr.DB.Create(b).Error; err != nil {
		logger.Error("Erro no DB.Create", err)
		panic(err)
	}
	pr.DB.Create(b)
	return nil
}

func (pr *PostgresRepository) List() []*entity.Book {
	var books []*entity.Book
	if err := pr.DB.Find(&books).Error; err != nil {
		logger.Error("Erro no DB.Find", err)
		panic(err.Error)
	}

	pr.DB.Find(&books)
	logger.Info("Find ok")
	return books
}

func (pr *PostgresRepository) FindById(id uuid.UUID) *entity.Book {
	var book *entity.Book
	if err := pr.DB.First(&book, id).Error; err != nil {
		logger.Error("Erro no DB.First id = "+id.String(), err)
		panic(err.Error)
	}

	logger.Info("FindById ok")
	return book
}

func (pr *PostgresRepository) Update(b *entity.Book) error {
	if err := pr.DB.Model(b).Update("title", b.Title).Error; err != nil {
		logger.Error("Erro no DB.Model(b).Update", err)
		panic(err.Error)
	}

	return nil
}

func (pr *PostgresRepository) Delete(id uuid.UUID) error {
	if err := pr.DB.Delete(&entity.Book{}, id).Error; err != nil {
		logger.Error("Erro no DB.Delete = "+id.String(), err)
		panic(err.Error)
	}
	return nil
}
