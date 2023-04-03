package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-wire/app/service"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

// GetAllUser godoc
// @summary Get All User
// @description  Get all user
// @tags user
// @accept json
// @produce json
// @response 200 {object} dto.ApiResponse[dao.User] "OK"
// @Router /api/user [get]
func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

// CreateUser godoc
// @summary Create User
// @description  Create new user
// @tags user
// @accept json
// @produce json
// @param User body dto.CreateUserDto true "User data to be created"
// @response 200 {object} dto.ApiResponse[dao.User] "OK"
// @Router /api/user [post]
func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func UserControllerInit(UserService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: UserService,
	}
}
