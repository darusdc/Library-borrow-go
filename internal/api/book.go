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

type BookAPI struct {
	bookServices domain.BookServices
}

func NewBook(app *fiber.App,
	bookServices domain.BookServices,
	auzMid fiber.Handler) {
	ba := BookAPI{bookServices: bookServices}

	bookApp := app.Group("/books", auzMid)

	bookApp.Get("", ba.Index)
	bookApp.Post("", ba.Create)
	bookApp.Get("/:id", ba.Show)
	bookApp.Put("/:id", ba.Update)
	bookApp.Delete("/:id", ba.Delete)
}

func (ba BookAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)

	defer cancel()
	res, err := ba.bookServices.Index(c)

	if err != nil {
		return ctx.Status(
			http.StatusInternalServerError).
			JSON(dto.CreateResponseError(
				err.Error(),
			),
			)
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ba BookAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookDataRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(
			http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				"validation failed", fails))
	}

	err := ba.bookServices.Create(c, req)

	if err != nil {
		return ctx.Status(
			http.StatusInternalServerError).JSON(
			dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(
		http.StatusCreated).JSON(
		dto.CreateResponseSuccess(""))
}

func (ba BookAPI) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := ba.bookServices.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusAccepted).
		JSON(dto.
			CreateResponseSuccess(res),
		)
}

func (ba BookAPI) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookDataRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}

	id := ctx.Params("id")
	req.Id = id
	if fails := util.Validate(req); len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed:", fails))
	}

	err := ba.bookServices.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (ba BookAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	err := ba.bookServices.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusOK)
}
