package handler

import (
	"alterraseven/app/user"
	"alterraseven/dto"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(user user.Service) *UserHandler {
	return &UserHandler{userService: user}
}

func (uh *UserHandler) Login(secretKey []byte) func(echo.Context) error {
	return func(c echo.Context) error {
		var input dto.AuthLoginRequest
		if err := c.Bind(&input); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		fmt.Println(input)
		result, err := uh.userService.Login(c.Request().Context(), input.Email, input.Password)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		// generate claims
		issuedAt := time.Now()
		expiresAt := time.Now().Add(5 * time.Minute)
		accessClaims := dto.Claims{
			Email: result.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiresAt.Unix(),
				IssuedAt:  issuedAt.Unix(),
				NotBefore: issuedAt.Unix(),
				Issuer:    "echo app",
				Subject:   strconv.Itoa(int(result.ID)),
			},
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
		if err != nil {
			return c.String(http.StatusUnauthorized, "failed to sign access token")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":       "success",
			"access_token": token,
			"expires_at":   expiresAt.Unix(),
		})
	}
}

func (uh *UserHandler) Index(c echo.Context) error {
	response, err := uh.userService.Index(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, response)
}

func (uh *UserHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	response, err := uh.userService.Get(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, response)
}

func (uh *UserHandler) Update(c echo.Context) error {
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

	var input dto.UserUpdateRequest
	err = c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	update, err := uh.userService.Update(c.Request().Context(), id, input)
	if err != nil {
		return err
	}

	return c.JSON(200, update)
}

func (uh *UserHandler) Delete(c echo.Context) error {
	return c.JSON(200, "not implemented")
}

func (uh *UserHandler) Create(c echo.Context) error {
	return c.JSON(200, "not implemented")
}

func extractUserId(c echo.Context) (userId int, err error) {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(*dto.Claims)
		userId, err = strconv.Atoi(claims.Subject)
		return userId, err
	}
	return userId, errors.New("invalid claims")
}
