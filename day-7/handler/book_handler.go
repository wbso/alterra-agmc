package handler

import (
	"alterraseven/app/book"
	"alterraseven/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookService book.Service
}

func NewBookHandler(book book.Service) *BookHandler {
	return &BookHandler{bookService: book}
}

func (h *BookHandler) All(c echo.Context) error {
	books := h.bookService.Index(c.Request().Context())
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (h *BookHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	books := h.bookService.Get(c.Request().Context(), id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (h *BookHandler) UpdateByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var input dto.BookRequest
	err = c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	books := h.bookService.Update(c.Request().Context(), id, input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (h *BookHandler) DeleteByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	h.bookService.Delete(c.Request().Context(), id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   nil,
	})
}

func (h *BookHandler) Create(c echo.Context) error {
	var input dto.BookRequest
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	savedBook := h.bookService.Create(c.Request().Context(), input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   savedBook,
	})
}
