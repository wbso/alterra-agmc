package handler

import (
	"alterrasix/app/book"
	"alterrasix/app/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	router      *echo.Echo
	bookService book.Service
	userService user.Service
}

func New(user user.Service, book book.Service) *Handler {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} m${uri} ${latency_human}]\n",
	}))

	e.Use(middleware.Recover())

	//secretKey := []byte(os.Getenv("SECRET_KEY"))
	//router := &Router{
	//	Controller: c,
	//	Router:     e,
	//	SecretKey:  secretKey,
	//}
	return &Handler{
		router:      e,
		userService: user,
		bookService: book,
	}
}

func (h *Handler) ListenAndServe(address string) error {
	return h.router.Start(address)
}
