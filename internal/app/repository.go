package app

type Repository interface {
    Create(data *Category) bool
    Read(id string) (*Category , error)
    Update(data *Category , id string) bool
    Delete(id string) bool
    List() interface{}
}

type Category struct {
    Name string `json:"name" bson:"name"`
    Tag string `json:"tag" bson:"tag"`
}
