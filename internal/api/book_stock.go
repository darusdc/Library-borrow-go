package api

import (
	"context"
	"net/http"
	"time"

	"github.com/darusdc/belajar-go/domain"
	"github.com/darusdc/belajar-go/dto"
	"github.com/darusdc/belajar-go/internal/util"
	"github.com/gofiber/fiber/v2"
)

type BookStockAPI struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App,
	bookStockService domain.BookStockService,
	auzMidd fiber.Handler) {
	bookStockAPI := BookStockAPI{
		bookStockService: bookStockService,
	}

	bookstockGroup := app.Group("/books/stock", auzMidd)

	bookstockGroup.Get("/:id", bookStockAPI.StockBook)
	bookstockGroup.Post("/borrow/", bookStockAPI.BorrowBook)
	bookstockGroup.Post("/return/", bookStockAPI.ReturnBook)
	bookstockGroup.Delete("/:id", bookStockAPI.DeleteBookStockById)
	bookstockGroup.Delete("/:id/:code", bookStockAPI.DeleteBookStockByCode)

}

func (bookStockAPI BookStockAPI) StockBook(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	res, err := bookStockAPI.bookStockService.CheckStock(c, id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).JSON(dto.CreateResponseSuccess(res))

}

func (bookStockAPI BookStockAPI) BorrowBook(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.UpdateBookStocksDataRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed:", fails))
	}

	err := bookStockAPI.bookStockService.Borrow(c, req.BookId, req.Code, req.BorrowerId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).JSON(dto.CreateResponseSuccess(""))

}

func (bookStockAPI BookStockAPI) ReturnBook(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.UpdateBookStocksDataRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed:", fails))
	}

	err := bookStockAPI.bookStockService.Returned(c, req.BookId, req.Code)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).JSON(dto.CreateResponseSuccess(""))

}

func (bookStockAPI BookStockAPI) DeleteBookStockById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := bookStockAPI.bookStockService.DeleteByBookId(c, id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).JSON(dto.CreateResponseSuccess(""))

}

func (bookStockAPI BookStockAPI) DeleteBookStockByCode(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	code := ctx.Params("code")
	id := ctx.Params("id")
	err := bookStockAPI.bookStockService.DeleteByCodeAndId(c, code, id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).JSON(dto.CreateResponseSuccess(""))

}
