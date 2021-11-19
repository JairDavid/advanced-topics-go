package interfaces

type Repository interface {
	Create(string) interface{}
	Update(string) interface{}
}

func Create(r Repository, arg string) {
	r.Create(arg)
}

func Update(r Repository, arg string) {
	r.Update(arg)
}
