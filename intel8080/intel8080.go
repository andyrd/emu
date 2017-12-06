package intel8080

type Signal struct{}

type State struct {
	Memory []byte
	A      byte
	B      byte
	C      byte
	D      byte
	E      byte
	H      byte
	L      byte
	SP     uint16
	PC     uint16
	Flags  byte
}

type Intel8080 interface {
	DataBus() chan byte
	AddressBus() <-chan uint16
	RESET() chan<- Signal
	INT() chan<- Signal
	INTE() <-chan Signal
	Clock() chan<- int
	READY() chan<- Signal
	WAIT() <-chan State

	PowerOn()
	PowerOff()
}
