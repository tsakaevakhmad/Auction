package ports

import (
	"Auction/internal/core/domain/entity"
)

type icategoryService interface {
	CreateCategory(name string, parentId *string, childs ...[]string) error
	DeleteCategory(id string) error
	GetCategory(id string) (entity.Category, error)
	GetCategories() ([]entity.Category, error)
}
