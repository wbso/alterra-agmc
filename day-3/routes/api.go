package routes

import (
	"net/http"

	"alterrathree/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (r *Router) InitRouter() error {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &controllers.Claims{},
		SigningKey: r.SecretKey,
	})

	r.Router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	v1 := r.Router.Group("/v1")
	// auth routes
	v1.POST("/auth/login", r.Controller.LoginController(r.SecretKey))

	// books routes
	v1.GET("/books", r.Controller.GetAllBookController)
	v1.GET("/books/:id", r.Controller.GetBookByIDController)
	v1.PUT("/books/:id", r.Controller.UpdateBookByIDController, jwtMiddleware)
	v1.DELETE("/books/:id", r.Controller.DeleteBookByIDController, jwtMiddleware)
	v1.POST("/books", r.Controller.CreateBookController, jwtMiddleware)

	// api routes
	v1.GET("/users", r.Controller.GetAllUserController, jwtMiddleware)
	v1.GET("/users/:id", r.Controller.GetUserByIDController, jwtMiddleware)
	v1.PUT("/users/:id", r.Controller.UpdateUserByIDController, jwtMiddleware)
	v1.DELETE("/users/:id", r.Controller.DeleteUserByIDController, jwtMiddleware)
	v1.POST("/users", r.Controller.CreateUserController)

	return nil
}
