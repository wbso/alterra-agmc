package controllers

import (
	"net/http"
	"strconv"

	"alterrathree/models"

	"github.com/labstack/echo/v4"
)

func (con *Controller) GetAllBookController(c echo.Context) error {
	books := con.BookDB.GetAll()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (con *Controller) GetBookByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	books := con.BookDB.FindByID(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (con *Controller) UpdateBookByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var input models.Book
	err = c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	books := con.BookDB.Update(id, input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (con *Controller) DeleteBookByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	con.BookDB.Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   nil,
	})
}

func (con *Controller) CreateBookController(c echo.Context) error {
	var input models.Book
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	book := con.BookDB.Create(input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   book,
	})
}
