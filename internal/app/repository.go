package app

type Repository interface {
	Create(data *Category) bool
	Read(id int) (*Category , error)
	// Update(data *Category , id int) bool
	// Delete(id int) bool

}

type Category struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
