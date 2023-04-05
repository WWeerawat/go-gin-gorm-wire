package module

import (
	"github.com/google/wire"
	"go-gin-gorm-wire/app/controller"
	"go-gin-gorm-wire/app/repository"
	"go-gin-gorm-wire/app/service"
)

type RoleModule struct {
	Ctrl controller.RoleController
	Svc  service.RoleService
	Repo repository.RoleRepository
}

var roleServiceSet = wire.NewSet(service.RoleServiceInit,
	wire.Bind(new(service.RoleService), new(*service.RoleServiceImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

var roleCtrlSet = wire.NewSet(controller.RoleControllerInit,
	wire.Bind(new(controller.RoleController), new(*controller.RoleControllerImpl)),
)

var RoleModuleSet = wire.NewSet(
	wire.Struct(new(RoleModule), "*"),
	roleCtrlSet,
	roleServiceSet,
	roleRepoSet,
)
