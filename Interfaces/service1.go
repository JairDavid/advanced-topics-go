package interfaces

type Service1 struct {
}

func (s *Service1) Create(a string) string {
	return "Create from service 1 " + a
}

func (s *Service1) Update(a string) string {
	return "Update from service 1" + a
}
