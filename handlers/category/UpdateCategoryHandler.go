package category

import (
	"Auction/comands/category"
	"Auction/services/dbcontext"
	"gorm.io/gorm"
)

type UpdateCategoryHandler struct {
	db *gorm.DB
}

func NewUpdateCategoryHandler(db *dbcontext.PgContext) *UpdateCategoryHandler {
	return &UpdateCategoryHandler{
		db: db.Context(),
	}
}

func (gch UpdateCategoryHandler) Handler(query category.UpdateCategoryCommand) {
	gch.db.Exec("UPDATE categories SET name = ?, parent_id = ? WHERE id = ?", query.Name, query.ParentId, query.Id)
}
