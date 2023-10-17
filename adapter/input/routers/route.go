package routers

import (
	"net/http"

	"github.com/fabianogoes/dev-books-api/adapter/input/controller"
	"github.com/fabianogoes/dev-books-api/adapter/output/database"
	"github.com/fabianogoes/dev-books-api/application/domain/usecase"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/health", health)

	repository := database.NewInMemoryRepository()
	createBookUseCase := usecase.NewCreateBookUseCase(repository)
	listBookUseCase := usecase.NewListBookUseCase(repository)
	findBookUseCase := usecase.NewFindBookUseCase(repository)
	updateBookUseCase := usecase.NewUpdateBookUseCase(repository)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(repository)
	controller := controller.NewBookController(
		createBookUseCase,
		listBookUseCase,
		findBookUseCase,
		updateBookUseCase,
		deleteBookUseCase,
	)

	router.POST("/books", controller.Create)
	router.GET("/books", controller.List)
	router.GET("/books/:id", controller.FindById)
	router.PUT("/books/:id", controller.Update)
	router.DELETE("/books/:id", controller.Delete)
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to Dev Books API"})
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "UP"})
}
