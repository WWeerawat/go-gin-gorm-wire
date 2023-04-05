package config

import (
	"go-gin-gorm-wire/app/module"
)

type Initialization struct {
	UserModule module.UserModule
	RoleModule module.RoleModule
	AuthModule module.AuthModule
}

func NewInitialization(
	userModule module.UserModule,
	roleModule module.RoleModule,
	authModule module.AuthModule,

) *Initialization {
	return &Initialization{
		UserModule: userModule,
		RoleModule: roleModule,
		AuthModule: authModule,
	}
}
