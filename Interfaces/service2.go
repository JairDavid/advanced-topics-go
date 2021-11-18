package interfaces

type Service2 struct {
}

func (s *Service2) Create(a string) string {
	return "Create from service 2 " + a
}

func (s *Service2) Update(a string) string {
	return "Update from service 2 " + a
}
