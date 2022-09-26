package repository

import "alterraseven/entity"

type BookRepository interface {
	All() []entity.Book
	Find(int) entity.Book
	Update(int, entity.Book) entity.Book
	Create(book entity.Book) entity.Book
	Delete(int)
}

type Book struct {
	cursor int
	data   map[int]entity.Book
}

// get all data
func (d *Book) All() []entity.Book {
	var books []entity.Book
	for _, book := range d.data {
		books = append(books, book)
	}
	return books
}

// get data by id
func (d *Book) Find(id int) entity.Book {
	return d.data[id]
}

// update data by id
func (d *Book) Update(id int, book entity.Book) entity.Book {
	book.ID = id
	d.data[id] = book
	return book
}

// create new data
func (d *Book) Create(book entity.Book) entity.Book {
	// auto increment id
	d.cursor++
	book.ID = d.cursor
	d.data[d.cursor] = book

	return book
}

func (d *Book) Delete(id int) {
	delete(d.data, id)
}

func NewBookRepository() *Book {
	initialData := map[int]entity.Book{
		1: {
			ID:     1,
			Author: "Chinua Achebe",
			Title:  "Things Fall Apart",
			Year:   1958,
		},
		2: {
			ID:     2,
			Author: "Hans Christian Andersen",
			Title:  "Fairy tales",
			Year:   1836,
		},
		3: {
			ID:     3,
			Author: "Dante Alighieri",
			Title:  "The Divine Comedy",
			Year:   1315,
		},
	}

	return &Book{
		cursor: 4,
		data:   initialData,
	}
}
