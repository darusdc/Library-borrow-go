package dto

import "time"

type BookStocksData struct {
	BookId     string     `json:"book_id"`
	Code       string     `json:"code"`
	Status     string     `json:"status"`
	BorrowerId *string    `json:"borrower_id"`
	BorrowerAt *time.Time `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"`
}

type UpdateBookStocksDataRequest struct {
	BookId     string `json:"book_id" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Status     string `json:"status" validate:"required"`
	BorrowerId string `json:"borrower_id" validate:"required"`
}

type DeleteBookStocksByIdDataRequest struct {
	BookId string `json:"book_id" validate:"required"`
}

type DeleteBookStockByCodeDataRequest struct {
	BookId string `json:"book_id" validate:"required"`
	Code   string `json:"code" validate:"required"`
}
