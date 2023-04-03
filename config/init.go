package config

import (
	"go-gin-gorm-wire/app/controller"
	"go-gin-gorm-wire/app/repository"
	"go-gin-gorm-wire/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	roleRepo repository.RoleRepository
	roleSvc  service.RoleService
	RoleCtrl controller.RoleController
	authSvc  service.AuthService
	AuthCtrl controller.AuthController
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	roleService service.RoleService,
	roleCtrl controller.RoleController,
	authService service.AuthService,
	authCtrl controller.AuthController,
) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		roleRepo: roleRepo,
		roleSvc:  roleService,
		RoleCtrl: roleCtrl,
		authSvc:  authService,
		AuthCtrl: authCtrl,
	}
}
