package repositories

import (
	dbcontext "Auction/internal/adapters/db"
	"Auction/internal/core/domain/entity"
	"Auction/internal/core/ports"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(pgdb *dbcontext.PgContext) ports.ICategoryRepository {
	return &CategoryRepository{db: pgdb.Context()}
}

func (r CategoryRepository) CreateCategory(name string, parentId *string, childs ...string) error {
	category := &entity.Category{
		Name:     name,
		ParentID: parentId,
	}
	appendChilds(category, childs...)
	return r.db.Create(category).Error
}

func (r CategoryRepository) DeleteCategory(id string) error {
	return r.db.Exec("DELETE FROM categories WHERE id = ?", id).Error
}

func (r CategoryRepository) FindCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return &category, err
}

func (r CategoryRepository) FindCategories() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	categories = group(categories, nil, 0)
	return categories, err
}

func (r CategoryRepository) UpdateCategory(category entity.Category) error {
	var entity entity.Category
	err := r.db.Where("id = ?", category.ID).First(&entity).Error
	if err != nil {
		return err
	}
	entity.ParentID = category.ParentID
	entity.Name = category.Name
	return r.db.Save(&entity).Error
}

func group(categories []entity.Category, parrent *entity.Category, iteration int) []entity.Category {
	var groupedCategories []entity.Category
	if parrent == nil && iteration == 0 {
		for i := range categories {
			if categories[i].ParentID == nil {
				categories[i].Children = group(categories, &categories[i], iteration+1)
				groupedCategories = append(groupedCategories, categories[i])
			}
		}
	} else if parrent != nil {
		for i := range categories {
			if categories[i].ParentID != nil && *categories[i].ParentID == parrent.ID {
				categories[i].Children = group(categories, &categories[i], iteration+1)
				groupedCategories = append(groupedCategories, categories[i])
			}
		}
	}
	return groupedCategories
}

func appendChilds(category *entity.Category, childs ...string) {
	if childs == nil || len(childs) == 0 {
		return
	}
	for i := range childs {
		category.Children = append(category.Children, entity.Category{Name: childs[i]})
	}
}
