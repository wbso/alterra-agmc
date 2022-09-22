package book

import (
	"alterrasix/dto"
	"alterrasix/entity"
	"alterrasix/repository"
	"context"
)

type Service interface {
	Index(context.Context) []dto.BookResponse
	Get(context.Context, int) dto.BookResponse
	Create(context.Context, dto.BookRequest) dto.BookResponse
	Update(context.Context, int, dto.BookRequest) dto.BookResponse
	Delete(context.Context, int)
}

type BookService struct {
	repo repository.BookRepository
}

func New(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (bs *BookService) Index(ctx context.Context) []dto.BookResponse {
	books := bs.repo.All()
	return bookEntitiesToDTO(books)
}

func (bs *BookService) Get(ctx context.Context, i int) dto.BookResponse {
	book := bs.repo.Find(i)
	return bookEntityToDTO(book)
}

func (bs *BookService) Create(ctx context.Context, request dto.BookRequest) dto.BookResponse {
	book := bs.repo.Create(entity.Book{
		Title:  request.Title,
		Author: request.Author,
		Year:   request.Year,
	})

	return bookEntityToDTO(book)
}

func (bs *BookService) Update(ctx context.Context, id int, request dto.BookRequest) dto.BookResponse {
	book := bs.repo.Update(id, entity.Book{
		Title:  request.Title,
		Author: request.Author,
		Year:   request.Year,
	})

	return bookEntityToDTO(book)
}

func (bs *BookService) Delete(ctx context.Context, id int) {
	bs.repo.Delete(id)
}

func bookEntityToDTO(book entity.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
}

func bookEntitiesToDTO(books []entity.Book) (res []dto.BookResponse) {
	for _, book := range books {
		res = append(res, bookEntityToDTO(book))
	}
	return res
}
