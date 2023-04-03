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
		user.GET("", init.UserCtrl.GetAllUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		role := api.Group("/role", service.AuthMiddleware)
		role.GET("", init.RoleCtrl.GetAllRole)
		role.POST("", init.RoleCtrl.AddRoleData)
		role.GET("/:roleID", init.RoleCtrl.GetRoleById)
		role.PUT("/:roleID", init.RoleCtrl.UpdateRole)
		role.DELETE("/:roleID", init.RoleCtrl.DeleteRole)

		auth := api.Group("/auth")
		auth.POST("/signup", init.UserCtrl.AddUserData)
		auth.POST("/login", init.AuthCtrl.Login)
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
