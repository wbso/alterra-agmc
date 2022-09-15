package routes

import (
	"alterrathree/controllers"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Controller *controllers.Controller
	Router     *echo.Echo
}

func New(c *controllers.Controller) *Router {
	e := echo.New()

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
