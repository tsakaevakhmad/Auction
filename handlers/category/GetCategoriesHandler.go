package category

import (
	"Auction/domain/entity"
	"Auction/queries/queries"
	"Auction/services/dbcontext"
	"gorm.io/gorm"
)

type GetCategoryHandler struct {
	db *gorm.DB
}

func NewGetCategoryHandler(db *dbcontext.PgContext) *GetCategoryHandler {
	return &GetCategoryHandler{db: db.Context()}
}

func (gch GetCategoryHandler) Handler(query queries.GetCategoryQuery) *[]entity.Category {
	var categories []entity.Category
	gch.db.Find(&categories)
	return &categories
}
