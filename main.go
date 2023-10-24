package main

import (
	"github.com/fabianogoes/dev-books-api/adapter/input/routers"
	"github.com/fabianogoes/dev-books-api/configuration/logger"
	"github.com/fabianogoes/dev-books-api/configuration/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Application it's Running!")
	DB := postgres.Init()

	router := gin.Default()
	routers.Init(router, DB)

	router.Run(":3000")
}
