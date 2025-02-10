package app


type repositoryInMemory struct {
	categories map[int]*Category
}

func CreateRepository() *repositoryInMemory {
	return &repositoryInMemory{
		make(map[int]*Category),
	}
}

func (r *repositoryInMemory) Create(c *Category) bool {
	r.categories[len(r.categories)] = c
	return true
}

func (r *repositoryInMemory) Delete (id int) bool {
    delete(r.categories , id)
    return true
}

func (r *repositoryInMemory) Read(id int) (*Category , error) {
    if len(r.categories) != 0 || id <= len(r.categories) {
        return r.categories[id] , nil
    }
    return nil ,  nil
}

func (r *repositoryInMemory) Update(c *Category , id int) bool{
    if r.categories[id] != nil {
        r.categories[id] = c
        return true
    }
    return false
}

