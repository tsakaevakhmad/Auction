package ports

import (
	"Auction/internal/core/domain/entity"
)

type ICategoryRepository interface {
	CreateCategory(name string, parentId *string, childs ...string) error
	DeleteCategory(id string) error
	FindCategory(id string) (*entity.Category, error)
	FindCategories() ([]entity.Category, error)
	UpdateCategory(category entity.Category) error
}
