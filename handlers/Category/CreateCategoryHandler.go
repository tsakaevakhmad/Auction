package Category

import (
	"Auction/comands/category"
	"Auction/domain/entity"
	"Auction/services/dbcontext"
	"gorm.io/gorm"
)

type CreateCategoryHandler struct {
	pgContext *gorm.DB
}

func NewCreateCategoryHandler(ctx dbcontext.PgContext) *CreateCategoryHandler {
	return &CreateCategoryHandler{
		pgContext: ctx.Context(),
	}
}

func (cch *CreateCategoryHandler) Handle(command category.CreateCategoryCommand) error {
	category := &entity.Category{
		Name:     command.Name,
		ParentID: command.ParentId,
	}
	db := cch.pgContext
	return db.Create(category).Error
}
