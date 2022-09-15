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

	// api routes
	v1.GET("/users", r.Controller.GetAllUserController)
	v1.GET("/users/:id", r.Controller.GetUserByIDController)
	v1.PUT("/users/:id", r.Controller.UpdateUserByIDController)
	v1.DELETE("/users/:id", r.Controller.DeleteUserByIDController)
	v1.POST("/users", r.Controller.CreateUserController)

	return nil
}
