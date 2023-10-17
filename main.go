package main

import (
	"github.com/fabianogoes/dev-books-api/adapter/input/routers"
	"github.com/fabianogoes/dev-books-api/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Application it's Running!")
	router := gin.Default()

	routers.InitRoutes(router)

	router.Run(":3000")
}
