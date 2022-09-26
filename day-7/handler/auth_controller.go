package handler

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginControllerInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//
//func (con *Controller) LoginController(secretKey []byte) func(echo.Context) error {
//	return func(c echo.Context) error {
//		var input LoginControllerInput
//		err := c.Bind(&input)
//		if err != nil {
//			return c.String(http.StatusBadRequest, "bad request")
//		}
//
//		// get user
//		user, err := con.Model.GetUserByEmail(input.Email)
//		if err != nil {
//			return c.String(http.StatusNotFound, "user not found")
//		}
//
//		// compare hash
//		err = bcrypt.CompareHashAndPassword(user.Password, []byte(input.Password))
//		if err != nil {
//			return c.String(http.StatusUnauthorized, "invalid credentials")
//		}
//
//		// generate claims
//		issuedAt := time.Now()
//		expiresAt := time.Now().Add(5 * time.Minute)
//		accessClaims := Claims{
//			Email: user.Email,
//			StandardClaims: jwt.StandardClaims{
//				ExpiresAt: expiresAt.Unix(),
//				IssuedAt:  issuedAt.Unix(),
//				NotBefore: issuedAt.Unix(),
//				Issuer:    "echo app",
//				Subject:   strconv.Itoa(int(user.ID)),
//			},
//		}
//
//		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
//		if err != nil {
//			return c.String(http.StatusUnauthorized, "failed to sign access token")
//		}
//
//		return c.JSON(http.StatusOK, map[string]interface{}{
//			"status":       "success",
//			"access_token": token,
//			"expires_at":   expiresAt.Unix(),
//		})
//	}
//}
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
