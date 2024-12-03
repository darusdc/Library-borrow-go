package domain

import (
	"context"
	"database/sql"

	"github.com/darusdc/belajar-go/dto"
)

const (
	STATUS_AVAILABLE = "AVAILABLE"
	STATUS_BORROWED  = "BORROWED"
	STATUS_NOT_AVAIL = "NOT_AVAIL"
	STATUS_DELETED   = "DELETED"
)

type BookStocks struct {
	BookId     string         `db:"book_id"`
	Code       string         `db:"code"`
	Status     string         `db:"status"`
	BorrowerId sql.NullString `db:"borrower_id"`
	BorrowerAt sql.NullTime   `db:"borrowed_at"`
	ReturnedAt sql.NullTime   `db:"returned_at"`
}

type BookStocksRepository interface {
	FindByBookId(ctx context.Context, bookId string) ([]BookStocks, error)
	FindByCodeAndId(ctx context.Context, code string, bookId string) (BookStocks, error)
	Save(context context.Context, data []BookStocks) error
	Update(ctx context.Context, bookStock *BookStocks) error
	DeleteByBookId(ctx context.Context, bookId string) error
	DeleteByCodeAndId(ctx context.Context, code string, bookId string) error
}

type BookStockService interface {
	Borrow(ctx context.Context, bookId string, code string, borrowerId string) error
	Returned(ctx context.Context, bookId string, code string) error
	CheckStock(ctx context.Context, bookId string) ([]dto.BookStocksData, error)
	DeleteByBookId(ctx context.Context, bookId string) error
	DeleteByCodeAndId(ctx context.Context, code string, bookId string) error
}
