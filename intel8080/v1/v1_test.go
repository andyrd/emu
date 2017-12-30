package v1

import (
	"testing"
	"time"

	ops "github.com/andyrd/emu/intel8080"
)

// repurpose an unused opcode to termiate the test
const terminateOp = 0x08

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

func TestDAD_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.DAD_B,
		terminateOp,
	})

	cpu.state.B = 0xFF
	cpu.state.C = 0xFE
	cpu.state.H = 0x00
	cpu.state.L = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x00 {
		t.Fatal("Invalid value in register H")
	}
	if cpu.state.L != 0x01 {
		t.Fatal("Invalid value in register L")
	}
	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid Flags value")
	}
}

func TestLDAX_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.LDAX_B,
		terminateOp,
		0xAA,
	})

	cpu.state.B = 0x00
	cpu.state.C = 0x02

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0xAA {
		t.Fatal("Invalid value in register A")
	}
}

func TestDCX_B(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCX_B,
		terminateOp,
	})

	cpu.state.B = 0x98
	cpu.state.C = 0x00

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.B != 0x97 || cpu.state.C != 0xFF {
		t.Fatal("Invalid value in register pair BC")
	}
}

func TestINR_C(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_C,
		terminateOp,
	})

	cpu.state.C = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.C != 0x10 {
		t.Fatal("Invalid value in register C")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_C(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_C,
		terminateOp,
	})

	cpu.state.C = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.C != 0x0E {
		t.Fatal("Invalid value in register C")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_C_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_C_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.C != 0xF5 {
		t.Fatal("Invalid value in register C")
	}
}

func TestRRC(t *testing.T) {
	cpu := initTest([]byte{
		ops.RRC,
		terminateOp,
	})

	cpu.state.A = 0xAA

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x02 {
		t.Fatal("Invalid Flags value")
	}
	if cpu.state.A != 0x55 {
		t.Fatal("Invalid value in register A")
	}
}

func TestLXI_D_D16(t *testing.T) {
	cpu := initTest([]byte{
		ops.LXI_D_D16, 0xCD, 0xAB,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0xAB {
		t.Fatal("Invalid value for register D")
	}
	if cpu.state.E != 0xCD {
		t.Fatal("Invalid value for register E")
	}
}

func TestSTAX_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.STAX_D,
		ops.NOP,
		terminateOp,
		0x00,
	})

	cpu.state.A = 0x07
	cpu.state.D = 0x00
	cpu.state.E = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Memory[0x03] != 7 {
		t.Fatal("Invalid value at memory address")
	}
}

func TestINX_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.INX_D,
		terminateOp,
	})

	cpu.state.D = 0x01
	cpu.state.E = 0xFF

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0x02 || cpu.state.E != 0x00 {
		t.Fatal("Invalid value in register pair DE")
	}
}

func TestINR_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_D,
		terminateOp,
	})

	cpu.state.D = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0x10 {
		t.Fatal("Invalid value in register D")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_D,
		terminateOp,
	})

	cpu.state.D = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0x0E {
		t.Fatal("Invalid value in register D")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_D_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_D_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0xF5 {
		t.Fatal("Invalid value in register D")
	}
}

func TestRAL(t *testing.T) {
	cpu := initTest([]byte{
		ops.RAL,
		terminateOp,
	})

	cpu.state.A = 0xAA

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid Flags value")
	}
	if cpu.state.A != 0x54 {
		t.Fatal("Invalid value in register A")
	}
}

func TestDAD_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.DAD_D,
		terminateOp,
	})

	cpu.state.D = 0xFF
	cpu.state.E = 0xFE
	cpu.state.H = 0x00
	cpu.state.L = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x00 {
		t.Fatal("Invalid value in register H")
	}
	if cpu.state.L != 0x01 {
		t.Fatal("Invalid value in register L")
	}
	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid Flags value")
	}
}

func TestLDAX_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.LDAX_D,
		terminateOp,
		0xAA,
	})

	cpu.state.D = 0x00
	cpu.state.E = 0x02

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0xAA {
		t.Fatal("Invalid value in register A")
	}
}

func TestDCX_D(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCX_D,
		terminateOp,
	})

	cpu.state.D = 0x98
	cpu.state.E = 0x00

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.D != 0x97 || cpu.state.E != 0xFF {
		t.Fatal("Invalid value in register pair DE")
	}
}

func TestINR_E(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_E,
		terminateOp,
	})

	cpu.state.E = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.E != 0x10 {
		t.Fatal("Invalid value in register E")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_E(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_E,
		terminateOp,
	})

	cpu.state.E = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.E != 0x0E {
		t.Fatal("Invalid value in register E")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_E_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_E_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.E != 0xF5 {
		t.Fatal("Invalid value in register E")
	}
}

func TestRAR(t *testing.T) {
	cpu := initTest([]byte{
		ops.RAR,
		terminateOp,
	})

	cpu.state.A = 0xAA

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x02 {
		t.Fatal("Invalid Flags value")
	}
	if cpu.state.A != 0x55 {
		t.Fatal("Invalid value in register A")
	}
}

func TestLXI_H_D16(t *testing.T) {
	cpu := initTest([]byte{
		ops.LXI_H_D16, 0xCD, 0xAB,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0xAB {
		t.Fatal("Invalid value for register D")
	}
	if cpu.state.L != 0xCD {
		t.Fatal("Invalid value for register E")
	}
}

func TestSHLD_A16(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.SHLD_A16
	cpuMem[1] = 0x0A
	cpuMem[2] = 0x01
	cpuMem[3] = terminateOp

	cpu := initTest(cpuMem)
	cpu.state.H = 0xAE
	cpu.state.L = 0x29

	cpu.PowerOn()
	<-cpu.done

	if cpuMem[0x010A] != 0x29 || cpuMem[0x010B] != 0xAE {
		t.Fatal("Invalid memory values")
	}
}

func TestINX_H(t *testing.T) {
	cpu := initTest([]byte{
		ops.INX_H,
		terminateOp,
	})

	cpu.state.H = 0x01
	cpu.state.L = 0xFF

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x02 || cpu.state.L != 0x00 {
		t.Fatal("Invalid value in register pair HL")
	}
}

func TestINR_H(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_H,
		terminateOp,
	})

	cpu.state.H = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x10 {
		t.Fatal("Invalid value in register H")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_H(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_H,
		terminateOp,
	})

	cpu.state.H = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x0E {
		t.Fatal("Invalid value in register H")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_H_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_H_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0xF5 {
		t.Fatal("Invalid value in register H")
	}
}

func TestDAA(t *testing.T) {
	cpu := initTest([]byte{
		ops.DAA,
		terminateOp,
	})

	cpu.state.A = 0x9B

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0x01 {
		t.Fatal("Invalid value in register A")
	}

	if cpu.state.Flags != 0x13 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDAD_H(t *testing.T) {
	cpu := initTest([]byte{
		ops.DAD_H,
		terminateOp,
	})

	cpu.state.H = 0x00
	cpu.state.L = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x00 {
		t.Fatal("Invalid value in register H")
	}
	if cpu.state.L != 0x06 {
		t.Fatal("Invalid value in register L")
	}
	if cpu.state.Flags != 0x02 {
		t.Fatal("Invalid Flags value")
	}
}

func TestLHLD_A16(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.LHLD_A16
	cpuMem[1] = 0x00
	cpuMem[2] = 0x30
	cpuMem[3] = terminateOp
	cpuMem[0x3000] = 0x4E
	cpuMem[0x3001] = 0x06

	cpu := initTest(cpuMem)
	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x06 || cpu.state.L != 0x4E {
		t.Fatal("Invalid memory values")
	}
}

func TestDCX_H(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCX_H,
		terminateOp,
	})

	cpu.state.H = 0x98
	cpu.state.L = 0x00

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x97 || cpu.state.L != 0xFF {
		t.Fatal("Invalid value in register pair HL")
	}
}

func TestINR_L(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_L,
		terminateOp,
	})

	cpu.state.L = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.L != 0x10 {
		t.Fatal("Invalid value in register L")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_L(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_L,
		terminateOp,
	})

	cpu.state.L = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.L != 0x0E {
		t.Fatal("Invalid value in register L")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_L_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_L_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.L != 0xF5 {
		t.Fatal("Invalid value in register L")
	}
}

func TestCMA(t *testing.T) {
	cpu := initTest([]byte{
		ops.CMA,
		terminateOp,
	})

	cpu.state.A = 0x51

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0xAE {
		t.Fatal("Invalid value in register A")
	}
}

func TestLXI_SP_D16(t *testing.T) {
	cpu := initTest([]byte{
		ops.LXI_SP_D16, 0xCD, 0xAB,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.SP != 0xABCD {
		t.Fatal("Invalid value for SP")
	}
}

func TestSTA_A16(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.STA_A16
	cpuMem[1] = 0xB3
	cpuMem[2] = 0x05
	cpuMem[3] = terminateOp

	cpu := initTest(cpuMem)
	cpu.state.A = 0x55

	cpu.PowerOn()
	<-cpu.done

	if cpuMem[0x05B3] != 0x55 {
		t.Fatal("Invalid value at memory location")
	}
}

func TestINX_SP(t *testing.T) {
	cpu := initTest([]byte{
		ops.INX_SP,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.SP != 0x01 {
		t.Fatal("Invalid value in SP")
	}
}

func TestINR_M(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.INR_M
	cpuMem[1] = terminateOp
	cpuMem[0x05B3] = 0x0F

	cpu := initTest(cpuMem)
	cpu.state.H = 0x05
	cpu.state.L = 0xB3

	cpu.PowerOn()
	<-cpu.done

	if cpuMem[0x05B3] != 0x10 {
		t.Fatal("Invalid value at memory location")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_M(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.DCR_M
	cpuMem[1] = terminateOp
	cpuMem[0x05B3] = 0x0F

	cpu := initTest(cpuMem)
	cpu.state.H = 0x05
	cpu.state.L = 0xB3

	cpu.PowerOn()
	<-cpu.done

	if cpuMem[0x05B3] != 0x0E {
		t.Fatal("Invalid value at memory location")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_M_D8(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.MVI_M_D8
	cpuMem[1] = 0x0F
	cpuMem[2] = terminateOp

	cpu := initTest(cpuMem)
	cpu.state.H = 0x05
	cpu.state.L = 0xB3

	cpu.PowerOn()
	<-cpu.done

	if cpuMem[0x05B3] != 0x0F {
		t.Fatal("Invalid value at memory location")
	}
}

func TestSTC(t *testing.T) {
	cpu := initTest([]byte{
		ops.STC,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid value in flags")
	}
}

func TestDAD_SP(t *testing.T) {
	cpu := initTest([]byte{
		ops.DAD_SP,
		terminateOp,
	})

	cpu.state.SP = 0xFFFE
	cpu.state.H = 0x00
	cpu.state.L = 0x03

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.H != 0x00 {
		t.Fatal("Invalid value in register H")
	}
	if cpu.state.L != 0x01 {
		t.Fatal("Invalid value in register L")
	}
	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid Flags value")
	}
}

func TestLDA_A16(t *testing.T) {
	cpuMem := make([]byte, 0xFFFF)
	cpuMem[0] = ops.LDA_A16
	cpuMem[1] = 0xB3
	cpuMem[2] = 0x05
	cpuMem[3] = terminateOp
	cpuMem[0x05B3] = 0x55

	cpu := initTest(cpuMem)
	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0x55 {
		t.Fatal("Invalid value in register A")
	}
}

func TestDCX_SP(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCX_SP,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.SP != 0xFFFF {
		t.Fatal("Invalid value in SP")
	}
}

func TestINR_A(t *testing.T) {
	cpu := initTest([]byte{
		ops.INR_A,
		terminateOp,
	})

	cpu.state.A = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0x10 {
		t.Fatal("Invalid value in register A")
	}
	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestDCR_A(t *testing.T) {
	cpu := initTest([]byte{
		ops.DCR_A,
		terminateOp,
	})

	cpu.state.A = 0x0F

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0x0E {
		t.Fatal("Invalid value in register A")
	}

	if cpu.state.Flags != 0x12 {
		t.Fatal("Invalid value in Flags")
	}
}

func TestMVI_A_D8(t *testing.T) {
	cpu := initTest([]byte{
		ops.MVI_A_D8,
		0xF5,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.A != 0xF5 {
		t.Fatal("Invalid value in register A")
	}
}

func TestCMC(t *testing.T) {
	cpu := initTest([]byte{
		ops.CMC,
		terminateOp,
	})

	cpu.PowerOn()
	<-cpu.done

	if cpu.state.Flags != 0x03 {
		t.Fatal("Invalid value in Flags")
	}
}
