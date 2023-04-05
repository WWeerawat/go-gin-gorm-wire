package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin-gorm-wire/app/service"
	"go-gin-gorm-wire/config"
	_ "go-gin-gorm-wire/docs"
	"net/http"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", HealthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		user := api.Group("/user", service.AuthMiddleware)
		user.GET("", init.UserModule.Ctrl.GetAllUserData)
		user.GET("/:userID", init.UserModule.Ctrl.GetUserById)
		user.PUT("/:userID", init.UserModule.Ctrl.UpdateUserData)
		user.DELETE("/:userID", init.UserModule.Ctrl.DeleteUser)

		role := api.Group("/role", service.AuthMiddleware)
		role.GET("", init.RoleModule.Ctrl.GetAllRole)
		role.POST("", init.RoleModule.Ctrl.AddRoleData)
		role.GET("/:roleID", init.RoleModule.Ctrl.GetRoleById)
		role.PUT("/:roleID", init.RoleModule.Ctrl.UpdateRole)
		role.DELETE("/:roleID", init.RoleModule.Ctrl.DeleteRole)

		auth := api.Group("/auth")
		auth.POST("/signup", init.UserModule.Ctrl.AddUserData)
		auth.POST("/login", init.AuthModule.Ctrl.Login)
	}

	return router
}

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /health [get]
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
