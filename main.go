package main

import (
	"github.com/joho/godotenv"
	"go-gin-gorm-wire/app/router"
	"go-gin-gorm-wire/config"
	_ "go-gin-gorm-wire/docs"
	"os"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

// @title Customers API
// @version 1.0
// @description This is description
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
