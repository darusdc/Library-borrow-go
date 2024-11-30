package repository

import (
	"context"
	"database/sql"

	"github.com/darusdc/belajar-go/domain"
	"github.com/doug-martin/goqu/v9"
)

type bookStocksRepository struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStocksRepository {
	return &bookStocksRepository{
		db: goqu.New("default", con),
	}
}

// DeleteByBookId implements domain.BookStocksRepository.
func (b *bookStocksRepository) DeleteByBookId(ctx context.Context, bookId string) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("book_id").
			Eq(bookId)).Set(goqu.Record{
		"status": domain.STATUS_DELETED,
	}).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

// DeleteByCode implements domain.BookStocksRepository.
func (b *bookStocksRepository) DeleteByCode(ctx context.Context, code string) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("code").
			Eq(code)).Set(goqu.Record{
		"status": domain.STATUS_DELETED,
	}).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

// FindByBookAndCode implements domain.BookStocksRepository.
func (b *bookStocksRepository) FindByCode(ctx context.Context, code string) (bookStock domain.BookStocks, err error) {
	dataset := b.db.From("book_stocks").
		Where(goqu.C("code").Eq(code), goqu.C("status").Neq(domain.STATUS_DELETED))

	err = dataset.ScanStructsContext(ctx, &bookStock)
	return
}

// FindByBookId implements domain.BookStocksRepository.
func (b *bookStocksRepository) FindByBookId(ctx context.Context, bookId string) (bookStocks []domain.BookStocks, err error) {
	dataset := b.db.From("book_stocks").
		Where(goqu.C("book_id").Eq(bookId),
			goqu.C("status").Neq(domain.STATUS_DELETED),
		)

	err = dataset.ScanStructsContext(ctx, &bookStocks)
	return
}

// Save implements domain.BookStocksRepository.
func (b *bookStocksRepository) Save(context context.Context, data []domain.BookStocks) error {
	executor := b.db.Insert("book_stocks").Rows(
		data,
	).Executor()

	_, err := executor.ExecContext(context)
	return err
}

// Update implements domain.BookStocksRepository.
func (b *bookStocksRepository) Update(ctx context.Context, bookStock *domain.BookStocks) error {
	executor := b.db.Update("book_stocks").Where(
		goqu.C("book_id").Eq(bookStock.BookId),
	).Set(bookStock).Executor()

	_, err := executor.ExecContext(ctx)

	return err
}
