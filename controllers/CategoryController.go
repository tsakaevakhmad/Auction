package controllers

import (
	categoryCommands "Auction/comands/category"
	categoryHandler "Auction/handlers/category"
	categoryQuery "Auction/queries/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryControler struct {
	createHandler *categoryHandler.CreateCategoryHandler
	getHandler    *categoryHandler.GetCategoryHandler
	updateHandler *categoryHandler.UpdateCategoryHandler
}

func NewCategoryControler(
	createHandler *categoryHandler.CreateCategoryHandler,
	getHandler *categoryHandler.GetCategoryHandler,
	updateHandler *categoryHandler.UpdateCategoryHandler) *CategoryControler {
	return &CategoryControler{
		createHandler: createHandler,
		getHandler:    getHandler,
		updateHandler: updateHandler,
	}
}

// @Summary Добавить категорию
// @Description Добавляет категорию
// @Tags category
// @Accept json
// @Param request body category.CreateCategoryCommand true "Тело запроса"
// @Router /api/v1/category/create [post]
func (cc *CategoryControler) Create(c *gin.Context) {
	var cmd categoryCommands.CreateCategoryCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	cc.createHandler.Handle(cmd)
}

// @Summary Получить список категорий
// @Tags category
// @Accept json
// @Param request body category.GetCategoryQuery true "Тело запроса"
// @Router /api/v1/category/getall [post]
func (cc *CategoryControler) GetAll(c *gin.Context) {
	var query categoryQuery.GetCategoryQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	categories := cc.getHandler.Handler(query)
	c.JSON(http.StatusOK, categories)
}

// @Summary Обновить категорию
// @Tags category
// @Accept json
// @Param request body category.UpdateCategoryCommand true "Тело запроса"
// @Router /api/v1/category/update [put]
func (cc *CategoryControler) Update(c *gin.Context) {
	var command categoryCommands.UpdateCategoryCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	cc.updateHandler.Handler(command)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated category"})
}
