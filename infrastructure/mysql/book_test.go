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
		if err := mock.ExpectationsWereMet(); err != nil {
			log.Fatalf("there were unfulfilled expectations: %s", err)
		}
		if err := gdb.Close(); err != nil {
			log.Print(err)
		}
	}

	return &database{
		db: gdb,
	}, mock, cleanup
}

func TestBook_GetByID_ReturnBook(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookByIDQuery(*seed)
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "title", "author"}).
			AddRow(seed.ID(), seed.Isbn(), seed.Title(), seed.Author()))

	br := NewBookRepository(db)
	actual, err := br.GetByID(seed.ID())

	assert.NoError(t, err)
	assert.Equal(t, seed, actual)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestBook_GetByID_ReturnErrorWhenDBError(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookByIDQuery(*seed)
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("DB error"))

	br := NewBookRepository(db)
	actual, err := br.GetByID(seed.ID())

	assert.Error(t, err)
	assert.Nil(t, actual)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestBook_GetByID_ReturnErrorWhenDataNotFound(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookByIDQuery(*seed)
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows(nil))

	br := NewBookRepository(db)
	actual, err := br.GetByID(seed.ID())

	assert.IsType(t, myerror.NotFoundError{}, err)
	assert.Nil(t, actual)

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestBook_GetByID_ReturnErrorWhenToDomainFailed(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookByIDQuery(*seed)
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "title", "author"}).
			AddRow(seed.ID(), "invalid_isbn", seed.Title(), seed.Author()))

	br := NewBookRepository(db)
	actual, err := br.GetByID(seed.ID())

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_GetAll_ReturnsAllBooks(t *testing.T) {
	seed1, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}
	seed2, err := entity_book.NewBookForRebuild(234, "1234567890123", "実践ドメイン駆動設計", "	ヴァーン・ヴァーノン")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookAllQuery()
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "title", "author"}).
			AddRow(seed1.ID(), seed1.Isbn(), seed1.Title(), seed1.Author()).
			AddRow(seed2.ID(), seed2.Isbn(), seed2.Title(), seed2.Author()))

	br := NewBookRepository(db)
	actual, err := br.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, *seed1, actual[0])
	assert.Equal(t, *seed2, actual[1])
}

func TestBook_GetAll_ReturnsErrorWhenDBError(t *testing.T) {
	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookAllQuery()
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("DB error"))

	br := NewBookRepository(db)
	actual, err := br.GetAll()

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_GetAll_ReturnErrorWhenToDomainFailed(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	query := newGetBookAllQuery()
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "title", "author"}).
			AddRow(seed.ID(), "invalid_isbn", seed.Title(), seed.Author()))

	br := NewBookRepository(db)
	actual, err := br.GetAll()

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_Create_InsertsBook(t *testing.T) {
	seed, err := entity_book.NewBook("9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newCreateBookSQL()
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WithArgs(seed.Isbn(), seed.Title(), seed.Author()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	br := NewBookRepository(db)

	actual, err := br.Create(*seed)

	assert.NoError(t, err)
	assert.NotNil(t, actual.ID())
	assert.Equal(t, seed.Isbn(), actual.Isbn())
	assert.Equal(t, seed.Title(), actual.Title())
	assert.Equal(t, seed.Author(), actual.Author())
}

func TestBook_Create_ReturnsErrorWhenDBReturnsError(t *testing.T) {
	seed, err := entity_book.NewBook("9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newCreateBookSQL()
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("db error"))

	br := NewBookRepository(db)

	actual, err := br.Create(*seed)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_Update_UpdatesBook(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newUpdateBookSQL()
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WithArgs(seed.Isbn(), seed.Title(), seed.Author(), seed.ID()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	br := NewBookRepository(db)

	actual, err := br.Update(*seed)

	assert.NoError(t, err)
	assert.Equal(t, seed, actual)
}

func TestBook_Update_ReturnsErrorWhenDBReturnsError(t *testing.T) {
	seed, err := entity_book.NewBookForRebuild(123, "9784798121963", "エリック・エヴァンスのドメイン駆動設計", "エリック・エヴァンス")
	if err != nil {
		log.Fatal(err)
	}

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newUpdateBookSQL()
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("db error"))

	br := NewBookRepository(db)

	actual, err := br.Update(*seed)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_Delete_DeletesBook(t *testing.T) {
	targetID := uint64(123)

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newDeleteBookSQL(targetID)
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	br := NewBookRepository(db)

	err := br.Delete(targetID)

	assert.NoError(t, err)
}

func TestBook_Delete_ReturnsErrorWhenDBReturnsError(t *testing.T) {
	targetID := uint64(123)

	db, mock, cleanup := OpenTestDB()
	defer cleanup()

	mock.ExpectBegin()
	sql := newDeleteBookSQL(targetID)
	mock.ExpectExec(regexp.QuoteMeta(sql)).
		WillReturnError(errors.New("db error"))

	br := NewBookRepository(db)

	err := br.Delete(targetID)

	assert.Error(t, err)
}

func newGetBookByIDQuery(book entity_book.Book) string {
	return "SELECT * FROM `books`  WHERE (`books`.`id` = " +
		strconv.FormatUint(book.ID(), 10) +
		") ORDER BY `books`.`id` ASC LIMIT 1"
}

func newGetBookAllQuery() string {
	return "SELECT * FROM `books`"
}

func newCreateBookSQL() string {
	return "INSERT INTO `books` (`isbn`,`title`,`author`) VALUES (?,?,?)"
}

func newUpdateBookSQL() string {
	return "UPDATE `books` SET `isbn` = ?, `title` = ?, `author` = ?  WHERE `books`.`id` = ?"
}

func newDeleteBookSQL(id uint64) string {
	return "DELETE FROM `books`  WHERE (`books`.`id` = " + strconv.FormatUint(id, 10) + ")"
}
