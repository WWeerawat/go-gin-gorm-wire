package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"go-gin-gorm-wire/app/constant"
	"go-gin-gorm-wire/app/domain/dao"
	"go-gin-gorm-wire/app/domain/dto"
	"go-gin-gorm-wire/app/pkg"
	"go-gin-gorm-wire/app/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func AuthServiceInit(UserRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: UserRepository,
	}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (a AuthServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program login user")
	var request dto.LoginUserDto
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := a.UserRepository.FindUserById(request.ID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	valid, err := PasswordMatches(data.Password, request.Password)
	if err != nil || !valid {
		log.Error("Happened error when get invalid credentials. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	token, err := GenerateToken(data)
	if err != nil {
		log.Error("Happened error when generate token from user data. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, token))
}

func GenerateToken(u dao.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "test",
		ID:        strconv.Itoa(u.ID),
		Subject:   u.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func AuthMiddleware(c *gin.Context) {
	defer pkg.PanicHandler(c)
	headerToken := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(headerToken, "Bearer ")

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Error("Happened error when request had invalid credential. Error", err)
		pkg.PanicException(constant.Unauthorized)
	}
}

func PasswordMatches(userPassword string, requestPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(requestPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
