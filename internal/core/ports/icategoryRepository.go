package ports

import "Auction/internal/core/domain/dto/category"

type icategoryRepository interface {
	CreateCategory(name string, parentId *string, childs ...[]string) error
	DeleteCategory(id string) error
	GetCategory(id string) (*category.CategoryDto, error)
	GetCategories() ([]category.CategoryDto, error)
}
