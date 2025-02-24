package controllers

import (
	category2 "Auction/comands/category"
	"Auction/handlers/category"
	"Auction/queries/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryControler struct {
	createHandler *category.CreateCategoryHandler
	getHandler    *category.GetCategoryHandler
}

func NewCategoryControler(createHandler *category.CreateCategoryHandler, getHandler *category.GetCategoryHandler) *CategoryControler {
	return &CategoryControler{
		createHandler: createHandler,
		getHandler:    getHandler,
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

func (cc *CategoryControler) Get(c *gin.Context) {
	var query queries.GetCategoryQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	categories := cc.getHandler.Handler(query)
	c.JSON(http.StatusOK, categories)
}
