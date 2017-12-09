package v1

import (
	"testing"
	"time"

	ops "github.com/andyrd/emu/intel8080"
)

// repurpose an unused opcode to termiate the test
const terminateOp = 0x08

func TestLXI_B_D16(t *testing.T) {
	cpu := initTest([]byte{
		ops.LXI_B_D16, 0xCD, 0xAB,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0xAB {
		t.Fatal("Invalid value for register B")
	}

	if cpu.state.C != 0xCD {
		t.Fatal("Invalid value for register C")
	}
}

func TestSTAX_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.STAX_B,
		ops.NOP,
		terminateOp,
		0x00,
	})

	cpu.state.A = 0x07
	cpu.state.B = 0x00
	cpu.state.C = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Memory[0x03] != 7 {
		t.Fatal("Invalid value at memory address")
	}
}

func TestINX_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.INX_B,
		terminateOp,
	})

	cpu.state.B = 0x01
	cpu.state.C = 0xFF

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0x02 || cpu.state.C != 0x00 {
		t.Fatal("Invalid value in register pair BC")
	}
}

func TestINR_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_B,
		terminateOp,
	})

	cpu.state.B = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0x10 {
		t.Fatal("Invalid value in register B")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_B,
		terminateOp,
	})

	cpu.state.B = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0x0E {
		t.Fatal("Invalid value in register B")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_B_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_B_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0xF5 {
		t.Fatal("Invalid value in register B")
	}
}

func TestRLC(t *testing.T) {
	cpu := initTest([]byte{
		ops.RLC,
		terminateOp,
	})

	cpu.state.A = 0xAA

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid Flags value")
	}

	if cpu.state.A != 0x55 {
		t.Fatal("Invalid value in register A")
	}
}

// func TestDAD_B(t *testing.T) {
// 	cpu := initTest([]byte{
// 		ops.DAD_B,
// 		terminateOp,
// 	})

// 	cpu.state.B = 0xFF

// }

func initTest(memory []byte) *v1 {
	s := ops.State{
		Memory: memory,
		Flags:  0x02,
	}

	cpu := NewV1(s).(*v1)
	cpu.handlers[terminateOp] = func() {
		cpu.PowerOff()
	}

	go func() {
		clock := cpu.Clock()
		for {
			time.Sleep(10 * time.Microsecond)
			clock <- 20
		}
	}()

	return cpu
}
