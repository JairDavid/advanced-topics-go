package interfaces

type Service2 struct {
}

func (s *Service2) Create(a string) interface{} {
	return "Create from service 2 " + a
}

func (s *Service2) Update(a string) interface{} {
	return "Update from service 2 " + a
}
