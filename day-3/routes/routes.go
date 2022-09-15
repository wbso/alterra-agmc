package routes

import (
	"alterrathree/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Controller *controllers.Controller
	Router     *echo.Echo
}

func New(c *controllers.Controller) *Router {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} ${latency_human}]\n",
	}))

	router := &Router{
		Controller: c,
		Router:     e,
	}
	router.InitRouter()
	return router
}

func (r *Router) ListenAndServe(address string) error {
	return r.Router.Start(address)
}
