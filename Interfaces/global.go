package interfaces

type Repository interface {
	Create(string) string
	Update(string) string
}

func Create(r Repository, arg string) {
	r.Create(arg)
}

func Update(r Repository, arg string) {
	r.Update(arg)
}
