package entity_book

import (
	string_util "my-app/utils/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

var seedBook = Book{
	id:     1,
	isbn:   "9784798121963",
	title:  "seed_title",
	author: "seed_author",
}

func TestNewBook_ReturnsBookEntity(t *testing.T) {
	expect := Book{
		isbn:   "9784798121963",
		title:  "seed_title",
		author: "seed_author",
	}
	actual, err := NewBook(
		expect.isbn,
		expect.title,
		expect.author,
	)

	assert.Equal(t, expect, *actual)
	assert.NoError(t, err)
}

func TestNewBook_ReturnsErrorWithInvalidData(t *testing.T) {
	actual, err := NewBook(
		"",
		seedBook.title,
		seedBook.author,
	)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestNewBookForRebuild_ReturnsBookEntity(t *testing.T) {
	actual, err := NewBookForRebuild(
		seedBook.id,
		seedBook.isbn,
		seedBook.title,
		seedBook.author,
	)

	assert.Equal(t, seedBook, *actual)
	assert.NoError(t, err)
}

func TestNewBookForRebuild_ReturnsErrorWithInvalidData(t *testing.T) {
	actual, err := NewBookForRebuild(
		seedBook.id,
		"",
		seedBook.title,
		seedBook.author,
	)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestBook_validate_ReturnNilOnValidDate(t *testing.T) {
	err := seedBook.validate()
	assert.NoError(t, err)
}

func TestBook_validate_ReturnsErrorOnInvalidateData(t *testing.T) {
	cases := []struct {
		name string
		book Book
	}{
		{
			name: "EmptyISBN",
			book: Book{
				id:     seedBook.id,
				isbn:   "",
				title:  seedBook.title,
				author: seedBook.author,
			},
		},
		{
			name: "TooShortISBN",
			book: Book{
				id:     seedBook.id,
				isbn:   "123456789012",
				title:  seedBook.title,
				author: seedBook.author,
			},
		},
		{
			name: "TooLongISBN",
			book: Book{
				id:     seedBook.id,
				isbn:   "12345678901234",
				title:  seedBook.title,
				author: seedBook.author,
			},
		},
		{
			name: "TooLongTitle",
			book: Book{
				id:     seedBook.id,
				isbn:   seedBook.isbn,
				title:  string_util.RandamString(256),
				author: seedBook.author,
			},
		},
		{
			name: "TooLongAuthor",
			book: Book{
				id:     seedBook.id,
				isbn:   seedBook.isbn,
				title:  seedBook.title,
				author: string_util.RandamString(256),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			err := c.book.validate()
			assert.Error(t, err)
		})
	}
}

func TestBook_ID_ReturnsID(t *testing.T) {
	assert.Equal(t, seedBook.id, seedBook.ID())
}

func TestBook_Isbn_ReturnsIsbn(t *testing.T) {
	assert.Equal(t, seedBook.isbn, seedBook.Isbn())
}

func TestBook_Title_ReturnsTitle(t *testing.T) {
	assert.Equal(t, seedBook.title, seedBook.Title())
}

func TestBook_Author_ReturnsAuthor(t *testing.T) {
	assert.Equal(t, seedBook.author, seedBook.Author())
}
