package handler

import (
	"alterraseven/app/book"
	"alterraseven/app/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type Handler struct {
	router *echo.Echo
}

func New(user user.Service, book book.Service) *Handler {

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	bookHandler := NewBookHandler(book)
	userHandler := NewUserHandler(user)
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} m${uri} ${latency_human}]\n",
	}))

	e.Use(middleware.Recover())
	v1 := e.Group("/v1")

	v1.GET("/books", bookHandler.All)
	v1.GET("/books/:id", bookHandler.GetByID)
	v1.PUT("/books/:id", bookHandler.UpdateByID)
	v1.DELETE("/books/:id", bookHandler.DeleteByID)
	v1.POST("/books", bookHandler.Create)

	v1.POST("/auth/login", userHandler.Login(secretKey))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Hello, World!")
	})

	return &Handler{
		router: e,
	}
}

func (h *Handler) ListenAndServe(address string) error {
	return h.router.Start(address)
}
