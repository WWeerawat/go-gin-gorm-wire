//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-gin-gorm-wire/app/module"
)

var db = wire.NewSet(ConnectToDB)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		module.UserModuleSet,
		module.RoleModuleSet,
		module.AuthModuleSet,
	)
	return nil
}
