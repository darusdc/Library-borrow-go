package main

import (
	"net/http"

	"github.com/darusdc/belajar-go/config"
	"github.com/darusdc/belajar-go/dto"
	"github.com/darusdc/belajar-go/internal/api"
	"github.com/darusdc/belajar-go/internal/connection"
	"github.com/darusdc/belajar-go/internal/repository"
	"github.com/darusdc/belajar-go/internal/service"
	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	cnf := config.Get()

	dbConnect := connection.GetDatabases(cnf.Database)

	app := fiber.New()

	jwtMid := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("Unauthorize Access"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnect)
	booksRepository := repository.NewBook(dbConnect)
	bookStocksRepository := repository.NewBookStock(dbConnect)

	UserRepository := repository.NewUser(dbConnect)

	customerService := service.NewCustomer(customerRepository)
	booksService := service.NewBookService(booksRepository, bookStocksRepository)
	bookStocksService := service.NewBookStockService(bookStocksRepository, customerRepository)
	authService := service.NewAuth(cnf, UserRepository)

	api.NewCustomer(app, customerService, jwtMid)
	api.NewAuth(app, authService)
	api.NewBook(app, booksService, jwtMid)
	api.NewBookStock(app, bookStocksService, jwtMid)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
