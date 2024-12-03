package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/darusdc/belajar-go/domain"
	"github.com/doug-martin/goqu/v9"
)

type bookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &bookRepository{
		db: goqu.New("default", con),
	}
}

// Delete implements domain.BookRepository.
func (b *bookRepository) Delete(context context.Context, id string) error {
	executor := b.db.
		Update("books").
		Where(
			goqu.C("id").Eq(id),
		).Set(
		goqu.Record{
			"deleted_at": sql.NullTime{
				Valid: true,
				Time:  time.Now(),
			}}).Executor()
	_, err := executor.ExecContext(context)
	return err
}

// FindAll implements domain.BookRepository.
func (b *bookRepository) FindAll(context context.Context) (books []domain.Book, err error) {
	dataset := b.db.From("books").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(context, &books)
	return
}

// FindById implements domain.BookRepository.
func (b *bookRepository) FindById(context context.Context, id string) (book domain.Book, err error) {
	dataset := b.db.From("books").
		Where(goqu.C("deleted_at").IsNull(),
			goqu.C("id").Eq(id),
		)

	_, err = dataset.ScanStructContext(context, &book)

	return
}

// Save implements domain.BookRepository.
func (b *bookRepository) Save(context context.Context, book *domain.Book) error {
	executor := b.db.Insert("books").Rows(book).Executor()

	_, err := executor.ExecContext(context)
	return err
}

// Update implements domain.BookRepository.
func (b *bookRepository) Update(context context.Context, book *domain.Book) error {
	executor := b.db.Update("books").
		Where(
			goqu.C("id").Eq(book.Id),
		).Set(book).Executor()

	_, err := executor.ExecContext(context)

	return err
}
