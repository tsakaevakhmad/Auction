package services

import (
	"Auction/internal/core/domain/dto/category"
	"Auction/internal/core/domain/entity"
	"Auction/internal/core/ports"
)

type CategoryServices struct {
	repo ports.ICategoryRepository
}

func NewCategoryServices(repo ports.ICategoryRepository) ports.IСategoryServices {
	return &CategoryServices{repo}
}

func (c CategoryServices) CreateCategory(data category.CreateCategory) error {
	return c.repo.CreateCategory(data.Name, data.ParentId, data.Childs...)
}

func (c CategoryServices) DeleteCategory(id string) error {
	return c.repo.DeleteCategory(id)
}

func (c CategoryServices) GetCategory(id string) (*category.CategoryDto, error) {
	result, err := c.repo.FindCategory(id)
	return category.CategoryDto{}.MapFromCategory(result), err
}

func (c CategoryServices) GetCategories() ([]category.CategoryDto, error) {
	result, err := c.repo.FindCategories()
	var categories []category.CategoryDto
	for cat := range result {
		categories = append(
			categories, *category.CategoryDto{}.MapFromCategory(&result[cat]),
		)
	}
	return categories, err
}

func (c CategoryServices) UpdateCategory(data category.UpdateCategory) error {
	return c.repo.UpdateCategory(entity.Category{ID: data.Id, Name: data.Name})
}
