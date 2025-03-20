package category

type CategoryDto struct {
	ID       string
	Name     string
	ParentID *string
	Parent   *CategoryDto
	Children []CategoryDto
}
