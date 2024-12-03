package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/darusdc/belajar-go/domain"
	"github.com/darusdc/belajar-go/dto"
	"github.com/google/uuid"
)

type bookServices struct {
	bookRepository       domain.BookRepository
	bookStocksRepository domain.BookStocksRepository
}

func NewBookService(bookRepository domain.BookRepository, bookStockRepository domain.BookStocksRepository) domain.BookServices {
	return &bookServices{
		bookRepository:       bookRepository,
		bookStocksRepository: bookStockRepository,
	}
}

// Create implements domain.BookServices.
func (b *bookServices) Create(context context.Context, req dto.CreateBookDataRequest) error {
	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	var books []domain.BookStocks
	for i := 0; i < req.Stock; i++ {
		books = append(books, domain.BookStocks{
			BookId: book.Id,
			Code:   uuid.NewString(),
			Status: domain.STATUS_AVAILABLE,
		})
	}

	if err := b.bookStocksRepository.Save(context, books); err != nil {
		return err
	}

	return b.bookRepository.Save(context, &book)
}

// Delete implements domain.BookServices.
func (b *bookServices) Delete(context context.Context, id string) error {
	existed, err := b.bookRepository.FindById(context, id)

	if err != nil {
		return err
	}

	if existed.Id == "" {
		return errors.New("id isn't exist")
	}

	_, err = b.bookStocksRepository.FindByBookId(context, existed.Id)

	if err != nil {
		return err
	}

	if err := b.bookStocksRepository.DeleteByBookId(context, existed.Id); err != nil {
		return err
	}

	return b.bookRepository.Delete(context, existed.Id)
}

// Index implements domain.BookServices.
func (b *bookServices) Index(context context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(context)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var booksData []dto.BookData

	for _, v := range books {
		booksData = append(booksData, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
		})
	}

	return booksData, nil
}

// Show implements domain.BookServices.
func (b *bookServices) Show(context context.Context, id string) (dto.BookData, error) {
	book, err := b.bookRepository.FindById(context, id)

	if err != nil {
		return dto.BookData{}, err
	}

	if book.Id == "" {
		return dto.BookData{}, errors.New("data not found")
	}

	return dto.BookData{
		Id:          book.Id,
		Isbn:        book.Isbn,
		Title:       book.Title,
		Description: book.Description,
	}, nil

}

// Update implements domain.BookServices.
func (b *bookServices) Update(context context.Context, req dto.UpdateBookDataRequest) error {
	persisted, err := b.bookRepository.FindById(context, req.Id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("data not found")
	}

	persisted.Description = req.Description
	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return b.bookRepository.Update(context, &persisted)

}
