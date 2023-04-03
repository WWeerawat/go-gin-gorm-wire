package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-wire/app/service"
)

type RoleController interface {
	AddRoleData(c *gin.Context)
	GetAllRole(c *gin.Context)
	GetRoleById(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type RoleControllerImpl struct {
	svc service.RoleService
}

func RoleControllerInit(RoleService service.RoleService) *RoleControllerImpl {
	return &RoleControllerImpl{
		svc: RoleService,
	}
}

func (r RoleControllerImpl) AddRoleData(c *gin.Context) {
	r.svc.AddRoleData(c)
}

func (r RoleControllerImpl) GetAllRole(c *gin.Context) {
	r.svc.GetAllRole(c)
}

func (r RoleControllerImpl) GetRoleById(c *gin.Context) {
	r.svc.GetRoleById(c)
}

func (r RoleControllerImpl) UpdateRole(c *gin.Context) {
	r.svc.UpdateRole(c)
}

func (r RoleControllerImpl) DeleteRole(c *gin.Context) {
	r.svc.DeleteRole(c)
}
