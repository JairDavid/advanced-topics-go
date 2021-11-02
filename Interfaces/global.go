package interfaces

type Repository interface {
	Create()
	Update()
}

func Create(r Repository) {
	r.Create()
}

func Update(r Repository) {
	r.Update()
}
