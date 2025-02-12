package app

import "github.com/google/uuid"

type repositoryInMemory struct {
	categories map[string]*Category
}

func CreateRepository() *repositoryInMemory {
	return &repositoryInMemory{
		make(map[string]*Category),
	}
}

func (r *repositoryInMemory) Create(c *Category) bool {
    id := uuid.New()
    r.categories[id.String()] = c
	return true
}

func (r *repositoryInMemory) Delete (id string) bool {
    delete(r.categories , id)
    return true
}

func (r *repositoryInMemory) Read(id string) (*Category , error) {
    if len(r.categories) != 0 {
        return r.categories[id] , nil
    }
    return nil ,  nil
}

func (r *repositoryInMemory) Update(data *Category , id string) bool {
    if r.categories[id] != nil{
        r.categories[id] = data
        return true
    }
    return false
}

func (r *repositoryInMemory) List() interface{}{
    return r.categories
}

