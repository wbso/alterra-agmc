package routes

import (
	"alterrafour/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Controller *controllers.Controller
	Router     *echo.Echo
	SecretKey  []byte
}

func New(c *controllers.Controller) *Router {
	e := echo.New()

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "[${time_rfc3339}] ${status} ${method} m${uri} ${latency_human}]\n",
	// }))
	e.Use(middleware.Recover())

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	router := &Router{
		Controller: c,
		Router:     e,
		SecretKey:  secretKey,
	}
	router.InitRouter()
	return router
}

func (r *Router) ListenAndServe(address string) error {
	return r.Router.Start(address)
}
