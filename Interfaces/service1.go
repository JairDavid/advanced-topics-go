package interfaces

type Service1 struct {
}

func (s *Service1) Create(a string) interface{} {
	return "Create from service 1 " + a
}

func (s *Service1) Update(a string) interface{} {
	return "Update from service 1" + a
}
