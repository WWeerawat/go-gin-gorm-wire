package module

import (
	"github.com/google/wire"
	"go-gin-gorm-wire/app/controller"
	"go-gin-gorm-wire/app/repository"
	"go-gin-gorm-wire/app/service"
)

type UserModule struct {
	Ctrl controller.UserController
	Svc  service.UserService
	Repo repository.UserRepository
}

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var UserModuleSet = wire.NewSet(
	wire.Struct(new(UserModule), "*"),
	userCtrlSet,
	userServiceSet,
	userRepoSet,
)
