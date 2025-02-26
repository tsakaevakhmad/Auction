package category

import (
	"Auction/comands/category"
	"Auction/services/dbcontext"
	"gorm.io/gorm"
)

type DeleteCategoryHandler struct {
	db gorm.DB
}

func NewDeleteCategoryHandler(ctx *dbcontext.PgContext) *DeleteCategoryHandler {
	return &DeleteCategoryHandler{
		db: *ctx.Context(),
	}
}

func (dch *DeleteCategoryHandler) Handler(command category.DeleteCategoryCommand) {
	dch.db.Exec("DELETE FROM categories WHERE id = ?", command.ID)
}
