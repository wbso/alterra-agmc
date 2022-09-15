package controllers

import (
	"net/http"
	"strconv"

	"alterrathree/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (con *Controller) GetAllUserController(c echo.Context) error {
	users, err := con.Model.GetAllUser()
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

type Input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Output struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (con *Controller) CreateUserController(c echo.Context) error {
	var input Input
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	}

	user, err := con.Model.CreateUser(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": Output{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

func (con *Controller) GetUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user, err := con.Model.GetUserByID((id))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": Output{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

type UpdateInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (con *Controller) UpdateUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	authId, err := extractUserId(c)
	if err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}

	// validate authorization
	if id != authId {
		return c.String(http.StatusForbidden, "you dont have permission to access this resource")
	}

	var input UpdateInput
	err = c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user, err := con.Model.GetUserByID(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.Name = input.Name
	user.Email = input.Email

	u, err := con.Model.UpdateUser(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": Output{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		},
	})
}

func (con *Controller) DeleteUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	authId, err := extractUserId(c)
	if err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}

	// validate authorization
	if id != authId {
		return c.String(http.StatusForbidden, "you dont have permission to access this resource")
	}

	user, err := con.Model.GetUserByID(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = con.Model.DeleteUser(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   nil,
	})
}
