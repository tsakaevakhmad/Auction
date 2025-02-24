package category

type CreateCategoryCommand struct {
	Name     string  `json:"name"`
	ParentId *string `json:"parentId"`
}
