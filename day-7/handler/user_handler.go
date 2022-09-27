package handler

import (
	"alterraseven/app/user"
	"alterraseven/dto"
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

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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
		accessClaims := Claims{
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

//
//func extractUserId(c echo.Context) (userId int, err error) {
//	user := c.Get("user").(*jwt.Token)
//	if user.Valid {
//		claims := user.Claims.(*Claims)
//		userId, err = strconv.Atoi(claims.Subject)
//		return userId, err
//	}
//	return userId, errors.New("invalid claims")
//}

//func (r *Router) userRouter(g *echo.Group) {
//	g.GET("/users", r.Controller.GetAllUserController, jwtMiddleware)
//	g.GET("/users/:id", r.Controller.GetUserByIDController, jwtMiddleware)
//	g.PUT("/users/:id", r.Controller.UpdateUserByIDController, jwtMiddleware)
//	g.DELETE("/users/:id", r.Controller.DeleteUserByIDController, jwtMiddleware)
//	g.POST("/users", r.Controller.CreateUserController)
//}
