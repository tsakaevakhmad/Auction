package category

import "Auction/internal/core/domain/entity"

type CategoryDto struct {
	ID       string
	Name     string
	ParentID *string
	Parent   *CategoryDto
	Children []CategoryDto
}

func (c CategoryDto) MapFromCategory(category *entity.Category) *CategoryDto {
	if category == nil {
		return nil
	}
	var children []CategoryDto
	for _, child := range category.Children {
		children = append(children, *CategoryDto{}.MapFromCategory(&child))
	}
	return &CategoryDto{
		ID:       category.ID,
		Name:     category.Name,
		ParentID: category.ParentID,
		Parent:   CategoryDto{}.MapFromCategory(category.Parent),
		Children: children,
	}
}

func (c CategoryDto) MapToCategory() *entity.Category {
	children := make([]entity.Category, len(c.Children))
	for i, child := range c.Children {
		children[i] = *child.MapToCategory()
	}

	var parent *entity.Category
	if c.Parent != nil {
		parent = c.Parent.MapToCategory()
	}

	return &entity.Category{
		ID:       c.ID,
		Name:     c.Name,
		ParentID: c.ParentID,
		Parent:   parent,
		Children: children,
	}
}
