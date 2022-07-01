package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"github.com/watarun54/spbill-api/server/domain"
	"github.com/watarun54/spbill-api/server/interfaces/database"
	"github.com/watarun54/spbill-api/server/usecase"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

func NewJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: signingKey,
	}
}

type AuthController struct {
	Interactor usecase.UserInteractor
}

func NewAuthController(sqlHandler database.SqlHandler) *AuthController {
	return &AuthController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AuthController) Login(c Context) (err error) {
	uForm := domain.UserForm{}
	c.Bind(&uForm)
	user, _ := controller.Interactor.UserByEmail(uForm.Email)
	if user.ID == 0 || user.HashedPassword != generateHash(uForm.Email, uForm.Password) {
		c.JSON(500, NewError(errors.New("invalid name or password")))
		return
	}

	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return
	}

	c.JSON(200, map[string]string{
		"token": t,
	})
	return
}

func userIDFromToken(c Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}

func generateHash(email string, password string) string {
	salt := "test" //TODO :change
	result := sha256.Sum256([]byte(email + ":" + password + ":" + salt))
	return hex.EncodeToString(result[:])
}
