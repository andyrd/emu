package v1

import (
	"github.com/andyrd/emu/intel8080"
)

const (
	carryFlag     byte = 0x01
	parityFlag    byte = 0x04
	halfCarryFlag byte = 0x10
	zeroFlag      byte = 0x40
	signFlag      byte = 0x80
)

type opCodeHandler func()

type v1 struct {
	state      intel8080.State
	dataBus    chan byte
	addressBus chan uint16
	reset      chan intel8080.Signal
	int        chan intel8080.Signal
	inte       chan intel8080.Signal
	clock      chan int
	ready      chan intel8080.Signal
	wait       chan intel8080.State
	done       chan struct{}
	cycles     int
	handlers   [0xFF]opCodeHandler
	poweredOn  bool
}

func NewV1(initialState intel8080.State) intel8080.Intel8080 {
	nv := &v1{
		state:      initialState,
		dataBus:    make(chan byte),
		addressBus: make(chan uint16),
		reset:      make(chan intel8080.Signal),
		int:        make(chan intel8080.Signal),
		inte:       make(chan intel8080.Signal),
		clock:      make(chan int),
		ready:      make(chan intel8080.Signal),
		wait:       make(chan intel8080.State),
		done:       make(chan struct{}),
		poweredOn:  false,
	}

	nv.initHandlers()

	return nv
}

func (v *v1) PowerOn() {
	if !v.poweredOn {
		v.poweredOn = true
		go v.mainLoop()
	}
}

func (v *v1) PowerOff() {
	if v.poweredOn {
		close(v.done)
		v.poweredOn = false
	}
}

func (v *v1) mainLoop() {
	for {
		select {
		case <-v.done:
			return
		case <-v.int:
			// handle interrupts
		case <-v.reset:
			// handle reset
		case <-v.ready:
			// handle ready
		default:
			if v.cycles >= 0 {
				op := v.state.Memory[v.state.PC]
				v.state.PC++
				v.handlers[op]()
			} else {
				v.cycles += <-v.clock
			}
		}
	}
}

func (v *v1) DataBus() chan byte {
	return v.dataBus
}

func (v *v1) AddressBus() <-chan uint16 {
	return v.addressBus
}

func (v *v1) RESET() chan<- intel8080.Signal {
	return v.reset
}

func (v *v1) INT() chan<- intel8080.Signal {
	return v.int
}

func (v *v1) INTE() <-chan intel8080.Signal {
	return v.inte
}

func (v *v1) Clock() chan<- int {
	return v.clock
}

func (v *v1) READY() chan<- intel8080.Signal {
	return v.ready
}

func (v *v1) WAIT() <-chan intel8080.State {
	return v.wait
}
