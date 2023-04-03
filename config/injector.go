//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-gin-gorm-wire/app/controller"
	"go-gin-gorm-wire/app/repository"
	"go-gin-gorm-wire/app/service"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleServiceSet = wire.NewSet(service.RoleServiceInit,
	wire.Bind(new(service.RoleService), new(*service.RoleServiceImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

var roleCtrlSet = wire.NewSet(controller.RoleControllerInit,
	wire.Bind(new(controller.RoleController), new(*controller.RoleControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleCtrlSet, roleServiceSet, roleRepoSet)
	return nil
}
