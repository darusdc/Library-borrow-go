package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/darusdc/belajar-go/domain"
	"github.com/darusdc/belajar-go/dto"
)

type bookStocksService struct {
	bookStockRepository domain.BookStocksRepository
	customerRepository  domain.CustomerRepository
}

// Borrow implements domain.BookStockService.
func (b *bookStocksService) Borrow(ctx context.Context, bookId string, code string, borrowerId string) error {
	persisted, err := b.bookStockRepository.FindByCodeAndId(ctx, code, bookId)
	if err != nil {
		return err
	}

	user, err := b.customerRepository.FindById(ctx, borrowerId)

	if err != nil {
		return err
	}

	if user.Id == "" {
		return errors.New("user not found")
	}

	persisted.Status = domain.STATUS_BORROWED
	persisted.BorrowerAt = sql.NullTime{Valid: true, Time: time.Now()}
	persisted.BorrowerId = sql.NullString{Valid: true, String: borrowerId}
	return b.bookStockRepository.Update(ctx, &persisted)
}

// Check implements domain.BookStockService.
func (b *bookStocksService) CheckStock(ctx context.Context, bookId string) ([]dto.BookStocksData, error) {
	persisted, err := b.bookStockRepository.FindByBookId(ctx, bookId)
	if err != nil {
		return nil, err
	}

	var books []dto.BookStocksData
	for _, book := range persisted {
		books = append(books, dto.BookStocksData{
			BookId:     book.BookId,
			Code:       book.Code,
			Status:     book.Status,
			BorrowerId: &book.BorrowerId.String,
			BorrowerAt: &book.BorrowerAt.Time,
			ReturnedAt: &book.ReturnedAt.Time,
		})
	}

	return books, nil
}

// DeleteByBookId implements domain.BookStockService.
func (b *bookStocksService) DeleteByBookId(ctx context.Context, bookId string) error {
	existed, err := b.bookStockRepository.FindByBookId(ctx, bookId)

	if err != nil {
		return err
	}

	if len(existed) <= 0 {
		return errors.New("book is not found")
	}

	return b.bookStockRepository.DeleteByBookId(ctx, bookId)

}

// DeleteByCode implements domain.BookStockService.
func (b *bookStocksService) DeleteByCodeAndId(ctx context.Context, code string, bookId string) error {
	existed, err := b.bookStockRepository.FindByCodeAndId(ctx, code, bookId)
	if err != nil {
		return err
	}

	if existed.BookId == "" {
		return errors.New("book not found")
	}

	return b.bookStockRepository.DeleteByCodeAndId(ctx, code, bookId)
}

// Returned implements domain.BookStockService.
func (b *bookStocksService) Returned(ctx context.Context, bookId string, code string) error {
	persisted, err := b.bookStockRepository.FindByCodeAndId(ctx, code, bookId)
	if err != nil {
		return err
	}

	persisted.Status = domain.STATUS_AVAILABLE
	persisted.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}
	return b.bookStockRepository.Update(ctx, &persisted)
}

func NewBookStockService(bookStockRepository domain.BookStocksRepository, customerRepository domain.CustomerRepository) domain.BookStockService {
	return &bookStocksService{
		bookStockRepository: bookStockRepository,
		customerRepository:  customerRepository,
	}
}
