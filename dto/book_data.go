package dto

type BookData struct {
	Id          string `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateBookDataRequest struct {
	Title       string `json:"title" validate:"required"`
	Isbn        string `json:"isbn" validate:"required"`
	Description string `json:"description" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}

type UpdateBookDataRequest struct {
	Id          string `json:"-"`
	Isbn        string `json:"isbn" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
