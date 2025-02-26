package category

type UpdateCategoryCommand struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	ParentId *string `json:"parentId"`
}
