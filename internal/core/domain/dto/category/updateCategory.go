package category

type UpdateCategory struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	ParentId *string `json:"parentId"`
}
