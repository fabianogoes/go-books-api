package controller

import (
	"fmt"
	"net/http"

	"github.com/fabianogoes/dev-books-api/adapter/input/controller/dto"
	"github.com/fabianogoes/dev-books-api/application/domain/usecase"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type bookController struct {
	createBookUseCase *usecase.CreateBookUseCase
	listBookUseCase   *usecase.ListBookUseCase
	findBookUseCase   *usecase.FindBookUseCase
	updateBookUseCase *usecase.UpdateBookUseCase
	deleteBookUseCase *usecase.DeleteBookUseCase
}

func NewBookController(
	cuc *usecase.CreateBookUseCase,
	luc *usecase.ListBookUseCase,
	fuc *usecase.FindBookUseCase,
	uuc *usecase.UpdateBookUseCase,
	duc *usecase.DeleteBookUseCase,
) *bookController {
	return &bookController{
		createBookUseCase: cuc,
		listBookUseCase:   luc,
		findBookUseCase:   fuc,
		updateBookUseCase: uuc,
		deleteBookUseCase: duc,
	}
}

func (bc *bookController) Create(c *gin.Context) {
	var err error
	var payload dto.CreateBookRequest

	if err = c.BindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err = bc.createBookUseCase.Create(payload.Title, payload.Description, payload.Author); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Book %s created successfully", payload.Title)})
}

func (bc *bookController) List(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bc.listBookUseCase.List())
}

func (bc *bookController) FindById(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if book := bc.findBookUseCase.FindById(id); book != nil {
		c.IndentedJSON(http.StatusOK, bc.listBookUseCase.List())
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Book %s not found!", id)})
}

func (bc *bookController) Update(c *gin.Context) {
	var err error
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	var payload dto.UpdateBookRequest

	if err = c.BindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err = bc.updateBookUseCase.Update(id, payload.Title, payload.Description, payload.Author); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Book %s updated successFully!", payload.Title)})
}

func (bc *bookController) Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if b := bc.findBookUseCase.FindById(id); b == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Book %s not found!", id)})
		return
	}

	if err := bc.deleteBookUseCase.Delete(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
