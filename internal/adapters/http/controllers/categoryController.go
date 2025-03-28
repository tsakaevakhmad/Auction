package controllers

import (
	"Auction/internal/core/domain/dto/category"
	"Auction/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	categoryServices ports.IСategoryServices
}

func NewCategoryController(categoryServices ports.IСategoryServices) *CategoryController {
	return &CategoryController{
		categoryServices,
	}
}

// @Summary		Добавить категорию
// @Description	Добавляет категорию
// @Tags			category
// @Accept			json
// @Param			request	body	category.CreateCategory	true	"Тело запроса"
// @Router			/api/v1/category/create [post]
func (cc *CategoryController) Create(c *gin.Context) {
	var cmd category.CreateCategory
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	cc.categoryServices.CreateCategory(cmd)
}

// @Summary	Получить список категорий
// @Tags		category
// @Accept		json
// @Router		/api/v1/category/getall [post]
func (cc *CategoryController) GetAll(c *gin.Context) {
	categories, err := cc.categoryServices.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary	Получить категорию
// @Tags		category
// @Accept		json
// @Param		id	path	string	true	"Идентификатор категории"
// @Router		/api/v1/category/Get/{id} [get]
func (cc *CategoryController) Get(c *gin.Context) {
	id := c.Param("id")
	result, err := cc.categoryServices.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
	}
	c.JSON(http.StatusOK, result)
}

// @Summary	Обновить категорию
// @Tags		category
// @Accept		json
// @Param		request	body	category.UpdateCategory	true	"Тело запроса"
// @Router		/api/v1/category/update [put]
func (cc *CategoryController) Update(c *gin.Context) {
	var command category.UpdateCategory
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	cc.categoryServices.UpdateCategory(command)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated category"})
}

// @Summary	Обновить категорию
// @Tags		category
// @Accept		json
// @Param		id	path	string	true	"Идентификатор категории"
// @Router		/api/v1/category/delete/{id} [delete]
func (cc *CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	cc.categoryServices.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted category"})
}
