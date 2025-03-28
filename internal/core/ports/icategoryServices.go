package ports

import (
	"Auction/internal/core/domain/dto/category"
)

type IСategoryServices interface {
	CreateCategory(data category.CreateCategory) error
	DeleteCategory(id string) error
	GetCategory(id string) (*category.CategoryDto, error)
	GetCategories() ([]category.CategoryDto, error)
	UpdateCategory(data category.UpdateCategory) error
}
