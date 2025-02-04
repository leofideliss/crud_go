package category

type repositoryMemory struct {
    categories map[string]*Category
}

func NewRepository() *repositoryMemory{
    return &repositoryMemory{
        make(map[string]*Category),
    }
}

func (r *repositoryMemory) Save (data interface{}) bool{
    return true
}
