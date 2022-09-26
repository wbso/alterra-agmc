package handler

import (
	"alterraseven/app/user"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(user user.Service) *UserHandler {
	return &UserHandler{userService: user}
}

//func (r *Router) userRouter(g *echo.Group) {
//	g.GET("/users", r.Controller.GetAllUserController, jwtMiddleware)
//	g.GET("/users/:id", r.Controller.GetUserByIDController, jwtMiddleware)
//	g.PUT("/users/:id", r.Controller.UpdateUserByIDController, jwtMiddleware)
//	g.DELETE("/users/:id", r.Controller.DeleteUserByIDController, jwtMiddleware)
//	g.POST("/users", r.Controller.CreateUserController)
//}
