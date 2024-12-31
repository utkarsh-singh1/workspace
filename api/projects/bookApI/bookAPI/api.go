package bookapi

import "time"

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

func new() *BookDb {

	bd := &BookDb{}

	bd.Bookmap = make(map[int]Book)

	bd.Id = 0

	return bd
}

func (bs *BookDb) CreateEntry(Name string, Author string, Date time.Time) int {

	b := Book{
		ID:     bs.Id,
		Name:   Name,
		Author: Author,
		Date:   Date,
	}

	bs.Bookmap[bs.Id] = b

	bs.Id++

	return b.ID

}
