package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-gin-gorm-wire/app/constant"
	"go-gin-gorm-wire/app/domain/dao"
	"go-gin-gorm-wire/app/pkg"
	"go-gin-gorm-wire/app/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update User data by id")
	UserID, _ := strconv.Atoi(c.Param("UserID"))

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.UserRepository.FindUserById(UserID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.RoleID = request.RoleID
	data.Email = request.Email
	data.Name = request.Password
	data.Status = request.Status
	data.UpdatedAt = time.Now()

	_, err = u.UserRepository.Save(&data)
	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get User by id")
	UserID, _ := strconv.Atoi(c.Param("UserID"))

	data, err := u.UserRepository.FindUserById(UserID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data User")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	data, err := u.UserRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data User")

	data, err := u.UserRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all User data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data User by id")
	UserID, _ := strconv.Atoi(c.Param("UserID"))

	err := u.UserRepository.DeleteUserById(UserID)
	if err != nil {
		log.Error("Happened Error when try delete data User from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func UserServiceInit(UserRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: UserRepository,
	}
}
