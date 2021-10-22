package interfaces

type Repository interface {
	Create(string) (interface{}, error)
}
