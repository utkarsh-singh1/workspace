package bookapi

import (
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Name   string
	Author string
	Date   time.Time
}

type BookDb struct {
	Bookmap map[int]Book
	Id      int
}

func New() *BookDb {

	bd := &BookDb{}

	bd.Bookmap = make(map[int]Book)

	bd.Id = 0

	return bd
}

func (bd *BookDb) CreateEntry(Name string, Author string, Date time.Time) int {

	b := Book{
		ID:     bd.Id,
		Name:   Name,
		Author: Author,
		Date:   Date,
	}

	bd.Bookmap[bd.Id] = b

	bd.Id++

	return b.ID

}

func (bd *BookDb) DeleteEntry(id int) error {

	if _, ok := bd.Bookmap[id] ; ok {

		return fmt.Errorf("Sorry couldn't any book by this id - %v ", id)
		
	}
	
	delete(bd.Bookmap, id)

	return nil

}

func (bd *BookDb) DeleteAll() {

	bd.Bookmap = make(map[int]Book)

}

func (bd *BookDb) GetById(id int) (Book, error) {

	if book, err := bd.Bookmap[id] ; !err {

		return book, nil
	}else {
		return Book{}, fmt.Errorf("This book with this id -%v does not Exist",id)

	}
}

func (bd *BookDb) GetAllEntry() []Book{

	allBook := make([]Book, 0, len(bd.Bookmap))
	
	for _, book := range bd.Bookmap {

		allBook = append(allBook, book)

	}

	return allBook
}

func (bd *BookDb) GetByAuthor(author string) []Book {

	allBook := make([]Book, 0, len(bd.Bookmap))

	for _, book := range bd.Bookmap {

		if book.Author == author {

			allBook = append(allBook, book)
		}
	}

	return allBook
}
