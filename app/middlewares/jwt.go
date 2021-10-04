package middlewares

import (
	_businesses "alterra/business"
	"alterra/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, _businesses.ErrInvalidTokenCredential)
		},
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint) (string, error) {
	claims := JWTCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, nil
}

func GetClaimUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*JWTCustomClaims)
}
