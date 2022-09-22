package handler

import (
	"alterrasix/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllBookController(c echo.Context) error {
	books := h.bookService.Index(c.Request().Context())
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func (h *Handler) GetBookByIDController(c echo.Context) error {
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

func (h *Handler) UpdateBookByIDController(c echo.Context) error {
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

func (h *Handler) DeleteBookByIDController(c echo.Context) error {
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

func (h *Handler) CreateBookController(c echo.Context) error {
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
