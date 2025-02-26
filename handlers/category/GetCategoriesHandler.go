package category

import (
	"Auction/domain/entity"
	"Auction/queries/category"
	"Auction/services/dbcontext"
	"gorm.io/gorm"
)

type GetCategoryHandler struct {
	db *gorm.DB
}

func NewGetCategoryHandler(db *dbcontext.PgContext) *GetCategoryHandler {
	return &GetCategoryHandler{db: db.Context()}
}

func (gch GetCategoryHandler) Handler(query category.GetCategoryQuery) *[]entity.Category {
	var categories []entity.Category
	gch.db.Find(&categories)
	categories = group(categories, nil, 0)
	return &categories
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
