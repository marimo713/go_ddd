package mysql

import (
	"errors"
	"log"
	entity_book "my-app/domain/entity/book"
	myerror "my-app/error"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func OpenTestDB() (Database, sqlmock.Sqlmock, func()) {
	mdb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	gdb, err := gorm.Open("mysql", mdb)
	if err != nil {
		log.Fatal(err)
	}

	cleanup := func() {
		if err := gdb.Close(); err != nil {
			log.Print(err)
		}
	}

	return &database{
		db: gdb,
	}, mock, cleanup
}

func TestBook_GetByID_ReturnBook(t *testing.T) {
	seed := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := "SELECT * FROM `books`  WHERE (`books`.`id` = " +
		strconv.FormatUint(seed.ID(), 10) +
		") ORDER BY `books`.`id` ASC LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "title", "author"}).
			AddRow(seed.ID(), seed.Isbn(), seed.Title(), seed.Author()))

	br := NewBookRepository(db)
	actual, err := br.GetByID(123)

	assert.NoError(t, err)
	assert.Equal(t, seed, *actual)
}

func TestBook_GetByID_ReturnNotFoundErrorWhenDBError(t *testing.T) {
	seed := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := "SELECT * FROM `books`  WHERE (`books`.`id` = " +
		strconv.FormatUint(seed.ID(), 10) +
		") ORDER BY `books`.`id` ASC LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("DB error"))

	br := NewBookRepository(db)
	actual, err := br.GetByID(123)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_GetByID_ReturnErrorWhenDataNotFound(t *testing.T) {
	seed := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := "SELECT * FROM `books`  WHERE (`books`.`id` = " +
		strconv.FormatUint(seed.ID(), 10) +
		") ORDER BY `books`.`id` ASC LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows(nil))

	br := NewBookRepository(db)
	actual, err := br.GetByID(123)

	assert.IsType(t, myerror.NotFoundError{}, err)
	assert.Nil(t, actual)
}
