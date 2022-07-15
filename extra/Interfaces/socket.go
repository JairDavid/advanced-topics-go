package interfaces

type Socket interface {
	Conexion(bool)
}

func Connect(s Socket) {
	s.Conexion(true)
}

func Disconnect(s Socket) {
	s.Conexion(false)
}
