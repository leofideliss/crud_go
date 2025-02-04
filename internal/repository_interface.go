package category

type Repository interface{
    Save(data interface{}) bool
}
