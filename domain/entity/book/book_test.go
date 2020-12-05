package entity_book

import (
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
	actual := NewBook(
		expect.isbn,
		expect.title,
		expect.author,
	)

	assert.Equal(t, expect, actual)
}

func TestNewBookForRebuild_ReturnsBookEntity(t *testing.T) {
	actual := NewBookForRebuild(
		seedBook.id,
		seedBook.isbn,
		seedBook.title,
		seedBook.author,
	)

	assert.Equal(t, seedBook, actual)
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
