package main

import (
	"github.com/arbeeorlar/initializers"
	"github.com/arbeeorlar/migrations"
	"github.com/arbeeorlar/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
}

func main() {
	migrations.AutoMigrate()
	mainRouter := gin.Default()
	routes.AuthRoutes(mainRouter)
	mainRouter.Run()
}
