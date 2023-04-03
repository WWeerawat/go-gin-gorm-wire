package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-gin-gorm-wire/app/constant"
	"go-gin-gorm-wire/app/domain/dao"
	"go-gin-gorm-wire/app/pkg"
	"go-gin-gorm-wire/app/repository"
	"net/http"
	"strconv"
	"time"
)

type RoleService interface {
	AddRoleData(c *gin.Context)
	GetAllRole(c *gin.Context)
	GetRoleById(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
}

func RoleServiceInit(RoleRepository repository.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{
		RoleRepository: RoleRepository,
	}
}

func (r RoleServiceImpl) AddRoleData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data Role")
	var request dao.Role
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	data, err := r.RoleRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (r RoleServiceImpl) GetAllRole(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data Role")

	data, err := r.RoleRepository.FindAllRole()
	if err != nil {
		log.Error("Happened Error when find all Role data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (r RoleServiceImpl) GetRoleById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get Role by id")
	RoleID, _ := strconv.Atoi(c.Param("RoleID"))

	data, err := r.RoleRepository.FindRoleById(RoleID)
	if err != nil {
		log.Error("Happened Error when find all Role data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (r RoleServiceImpl) UpdateRole(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update Role data by id")
	RoleID, _ := strconv.Atoi(c.Param("RoleID"))

	var request dao.Role
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := r.RoleRepository.FindRoleById(RoleID)
	if err != nil {
		log.Error("Happened Error when find all Role data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	data.Role = request.Role
	data.UpdatedAt = time.Now()

	_, err = r.RoleRepository.Save(&data)
	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (r RoleServiceImpl) DeleteRole(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete Role data by id")
	RoleID, _ := strconv.Atoi(c.Param("RoleID"))

	err := r.RoleRepository.DeleteRoleById(RoleID)
	if err != nil {
		log.Error("Happened Error when try delete data User from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}
