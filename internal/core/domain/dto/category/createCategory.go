package category

type CreateCategory struct {
	Name     string  `json:"name"`
	ParentId *string `json:"parentId"`
}
