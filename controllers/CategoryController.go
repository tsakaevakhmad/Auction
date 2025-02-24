package controllers

import (
	category2 "Auction/comands/category"
	"Auction/handlers/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryControler struct {
	createHandler *category.CreateCategoryHandler
}

func NewCategoryControler(createHandler *category.CreateCategoryHandler) *CategoryControler {
	return &CategoryControler{
		createHandler: createHandler,
	}
}

func (cc *CategoryControler) Create(c *gin.Context) {
	var cmd category2.CreateCategoryCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	cc.createHandler.Handle(cmd)
}
