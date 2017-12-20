package v1

func (v *v1) NOP() {
	v.cycles -= 4
}

func (v *v1) LXI_B_D16() {
	v.state.C = v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.B = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 10
}

func (v *v1) STAX_B() {
	v.state.Memory[uint16(v.state.B)<<8|uint16(v.state.C)] = v.state.A
	v.cycles -= 7
}

func (v *v1) INX_B() {
	result := (uint16(v.state.B)<<8 | uint16(v.state.C)) + 1
	v.state.B = byte(result >> 8)
	v.state.C = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_B() {
	result := v.state.B + 1
	v.setHalfCarryAdd(v.state.B, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.B = result
	v.cycles -= 5
}

func (v *v1) DCR_B() {
	result := v.state.B - 1
	v.setHalfCarrySub(v.state.B, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.B = result
	v.cycles -= 5
}

func (v *v1) MVI_B_D8() {
	v.state.B = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) RLC() {
	if (v.state.A & 0x80) == 0x80 {
		v.state.Flags |= carryFlag
	} else {
		v.state.Flags &= ^carryFlag
	}
	v.state.A = (v.state.A << 1) | (v.state.A >> 7)
	v.cycles -= 4
}

func (v *v1) DAD_B() {
	a := uint32(v.state.B)<<8 | uint32(v.state.C)
	b := uint32(v.state.H)<<8 | uint32(v.state.L)
	r32 := a + b
	v.setCarry16(r32)
	r := uint16(r32)
	v.state.H = uint8(r >> 8)
	v.state.L = uint8(r & 0xFF)
	v.cycles -= 10
}

func (v *v1) LDAX_B() {
	v.state.A = v.state.Memory[uint16(v.state.B)<<8|uint16(v.state.C)]
	v.cycles -= 7
}

func (v *v1) DCX_B() {
	result := (uint16(v.state.B)<<8 | uint16(v.state.C)) - 1
	v.state.B = byte(result >> 8)
	v.state.C = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_C() {
	result := v.state.C + 1
	v.setHalfCarryAdd(v.state.C, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.C = result
	v.cycles -= 5
}

func (v *v1) DCR_C() {
	result := v.state.C - 1
	v.setHalfCarrySub(v.state.C, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.C = result
	v.cycles -= 5
}

func (v *v1) MVI_C_D8() {
	v.state.C = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) RRC() {
	if (v.state.A & 0x01) == 0x01 {
		v.state.Flags |= carryFlag
	} else {
		v.state.Flags &= ^carryFlag
	}
	v.state.A = (v.state.A >> 1) | (v.state.A << 7)
	v.cycles -= 4
}

func (v *v1) LXI_D_D16() {
	v.state.E = v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.D = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 10
}

func (v *v1) STAX_D() {
	v.state.Memory[uint16(v.state.D)<<8|uint16(v.state.E)] = v.state.A
	v.cycles -= 7
}

func (v *v1) INX_D() {
	result := (uint16(v.state.D)<<8 | uint16(v.state.E)) + 1
	v.state.D = byte(result >> 8)
	v.state.E = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_D() {
	result := v.state.D + 1
	v.setHalfCarryAdd(v.state.D, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.D = result
	v.cycles -= 5
}

func (v *v1) DCR_D() {
	result := v.state.D - 1
	v.setHalfCarrySub(v.state.D, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.D = result
	v.cycles -= 5
}

func (v *v1) MVI_D_D8() {
	v.state.D = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) RAL() {
	carryValue := v.state.Flags & carryFlag

	if (v.state.A & 0x80) == 0x80 {
		v.state.Flags |= carryFlag
	} else {
		v.state.Flags &= ^carryFlag
	}

	v.state.A = (v.state.A << 1) | carryValue
	v.cycles -= 4
}

func (v *v1) DAD_D() {
	a := uint32(v.state.D)<<8 | uint32(v.state.E)
	b := uint32(v.state.H)<<8 | uint32(v.state.L)
	r32 := a + b
	v.setCarry16(r32)
	r := uint16(r32)
	v.state.H = uint8(r >> 8)
	v.state.L = uint8(r & 0xFF)
	v.cycles -= 10
}

func (v *v1) LDAX_D() {
	v.state.A = v.state.Memory[uint16(v.state.D)<<8|uint16(v.state.E)]
	v.cycles -= 7
}

func (v *v1) DCX_D() {
	result := (uint16(v.state.D)<<8 | uint16(v.state.E)) - 1
	v.state.D = byte(result >> 8)
	v.state.E = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_E() {
	result := v.state.E + 1
	v.setHalfCarryAdd(v.state.E, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.E = result
	v.cycles -= 5
}

func (v *v1) DCR_E() {
	result := v.state.E - 1
	v.setHalfCarrySub(v.state.E, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.E = result
	v.cycles -= 5
}

func (v *v1) MVI_E_D8() {
	v.state.E = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) RAR() {
	carryValue := v.state.Flags & carryFlag

	if (v.state.A & 0x01) == 0x01 {
		v.state.Flags |= carryFlag
	} else {
		v.state.Flags &= ^carryFlag
	}

	v.state.A = (carryValue << 7) | (v.state.A >> 1)
	v.cycles -= 4
}
