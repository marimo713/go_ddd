package mysql

import (
	entity_book "my-app/domain/entity/book"
	"my-app/domain/repository"
	myerror "my-app/error"
	"my-app/infrastructure/mysql/gorm_model"

	"github.com/jinzhu/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(database Database) repository.BookRepository {
	return bookRepository{
		db: database.conn(),
	}
}

func (repo bookRepository) GetByID(id uint64) (*entity_book.Book, error) {
	var book gorm_model.Book
	err := repo.db.First(&book, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, myerror.NewNotFoundError(err, "book not found")
		}
		return nil, err
	}
	domain, err := book.ToDomain()
	if err != nil {
		return nil, err
	}
	return domain, nil
}
