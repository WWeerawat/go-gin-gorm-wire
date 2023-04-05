package module

import (
	"github.com/google/wire"
	"go-gin-gorm-wire/app/controller"
	"go-gin-gorm-wire/app/service"
)

type AuthModule struct {
	Ctrl controller.AuthController
	Svc  service.AuthService
}

var authServiceSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var AuthModuleSet = wire.NewSet(
	wire.Struct(new(AuthModule), "*"),
	authCtrlSet,
	authServiceSet,
)
