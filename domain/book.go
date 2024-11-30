package domain

import (
	"context"
	"database/sql"

	"github.com/darusdc/belajar-go/dto"
)

type Book struct {
	Id          string       `db:"id"`
	Isbn        string       `db:"isbn"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(context context.Context) ([]Book, error)
	FindById(context context.Context, id string) (Book, error)
	Save(context context.Context, book *Book) error
	Update(context context.Context, book *Book) error
	Delete(context context.Context, id string) error
}

type BookServices interface {
	Index(context context.Context) ([]dto.BookData, error)
	Create(context context.Context, req dto.CreateBookDataRequest) error
	Update(context context.Context, req dto.UpdateBookDataRequest) error
	Delete(context context.Context, id string) error
	Show(context context.Context, id string) (dto.BookData, error)
}
