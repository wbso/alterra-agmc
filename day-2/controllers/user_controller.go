package controllers

import (
	"net/http"
	"strconv"

	"alterratwo/models"

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

	res := con.Model.DB.Create(&data)

	if res.Error != nil {
		return c.String(http.StatusBadRequest, res.Error.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": Output{
			ID:    data.ID,
			Name:  data.Name,
			Email: data.Email,
		},
	})
}

func (con *Controller) GetUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var user models.User
	res := con.Model.DB.First(&user, id)
	if res.Error != nil {
		return c.String(http.StatusBadRequest, res.Error.Error())
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
	var user models.User
	res := con.Model.DB.First(&user, id)
	if res.Error != nil {
		return c.String(http.StatusBadRequest, res.Error.Error())
	}

	var input UpdateInput
	err = c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user.Name = input.Name
	user.Email = input.Email

	res = con.Model.DB.Save(&user)

	if res.Error != nil {
		return c.String(http.StatusBadRequest, res.Error.Error())
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

func (con *Controller) DeleteUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var user models.User
	res := con.Model.DB.First(&user, id)
	if res.Error != nil {
		return c.String(http.StatusBadRequest, res.Error.Error())
	}

	con.Model.DB.Delete(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   nil,
	})
}
