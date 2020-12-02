package entity_book

type Book struct {
	id     uint64
	isbn   string
	title  string
	author string
}

func NewBook(id uint64, isbn string, title string, author string) Book {
	return Book{
		id:     id,
		isbn:   isbn,
		title:  title,
		author: author,
	}
}

func (b Book) ID() uint64 {
	return b.id
}

func (b Book) Isbn() string {
	return b.isbn
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Author() string {
	return b.author
}
