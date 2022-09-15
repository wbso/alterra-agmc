package inmemdb

// simple in memory database for golang
import (
	"alterrathree/models"
)

// struct for holding the data
type DB struct {
	cursor int
	data   map[int]models.Book
}

// get all data
func (d *DB) GetAll() []models.Book {
	var books []models.Book
	for _, book := range d.data {
		books = append(books, book)
	}
	return books
}

// get data by id
func (d *DB) FindByID(id int) models.Book {
	return d.data[id]
}

// update data by id
func (d *DB) Update(id int, book models.Book) models.Book {
	book.ID = id
	d.data[id] = book
	return book
}

// create new data
func (d *DB) Create(book models.Book) models.Book {
	// auto increment id
	d.cursor++
	book.ID = d.cursor
	d.data[d.cursor] = book

	return book
}

func (d *DB) Delete(id int) {
	delete(d.data, id)
}

func New() *DB {
	initialData := map[int]models.Book{
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

	return &DB{
		cursor: 4,
		data:   initialData,
	}
}
