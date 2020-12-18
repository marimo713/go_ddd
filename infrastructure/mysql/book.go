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

func (repo bookRepository) GetAll() ([]entity_book.Book, error) {
	var books []gorm_model.Book
	err := repo.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	var domainBooks []entity_book.Book
	for i := range books {
		domainBook, err := books[i].ToDomain()
		if err != nil {
			return nil, err
		}
		domainBooks = append(domainBooks, *domainBook)
	}
	return domainBooks, nil
}

func (repo bookRepository) Create(book entity_book.Book) (*entity_book.Book, error) {
	gormBook := gorm_model.NewBookFromDomain(book)
	result := repo.db.Create(&gormBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	return gormBook.ToDomain()
}

func (repo bookRepository) Update(book entity_book.Book) (*entity_book.Book, error) {
	gormBook := gorm_model.NewBookFromDomain(book)
	result := repo.db.Save(&gormBook)
	if err := result.Error; err != nil {
		return nil, err
	}

	return gormBook.ToDomain()
}

func (repo bookRepository) Delete(id uint64) error {
	result := repo.db.Delete(&gorm_model.Book{}, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
