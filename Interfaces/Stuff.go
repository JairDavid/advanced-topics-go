package interfaces

import "fmt"

type TV struct {
	works bool
}

type Radio struct {
	works bool
}

type CoffeMaker struct {
	works bool
}

func newCoffeMaker() *CoffeMaker {
	return &CoffeMaker{works: false}
}

func newTV() *TV {
	return &TV{works: false}
}

func newRadio() *Radio {
	return &Radio{works: false}
}

func (t *TV) Conexion(electricity bool) {
	t.works = electricity
}

func (r *Radio) Conexion(electricity bool) {
	r.works = electricity
}

func (cm *CoffeMaker) Conexion(electricity bool) {
	cm.works = electricity
}

func TestingImplementation() {
	myTv := newTV()
	myRadio := newRadio()
	myCoffeMaker := newCoffeMaker()

	Connect(myTv)
	Connect(myRadio)
	Connect(myCoffeMaker)

	fmt.Println(myTv.works)
	fmt.Println(myRadio.works)
	fmt.Println(myCoffeMaker.works)

	Disconnect(myTv)
	Disconnect(myRadio)
	Disconnect(myCoffeMaker)

	fmt.Println(myTv.works)
	fmt.Println(myRadio.works)
	fmt.Println(myCoffeMaker.works)
}
