package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-wire/app/service"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func AuthControllerInit(AuthService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: AuthService,
	}
}

func (a AuthControllerImpl) Login(c *gin.Context) {
	a.svc.Login(c)
}
