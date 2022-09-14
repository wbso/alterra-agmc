package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) InitRouter() error {
	r.Router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	v1 := r.Router.Group("/v1")
	// books routes
	v1.GET("/books", r.Controller.GetAllBookController)
	v1.GET("/books/:id", r.Controller.GetBookByIDController)
	v1.PUT("/books/:id", r.Controller.UpdateBookByIDController)
	v1.DELETE("/books/:id", r.Controller.DeleteBookByIDController)
	v1.POST("/books", r.Controller.CreateBookController)

	return nil
}
