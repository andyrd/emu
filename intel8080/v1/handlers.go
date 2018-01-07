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

func (v *v1) LXI_H_D16() {
	v.state.L = v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.H = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 10
}

func (v *v1) SHLD_A16() {
	lo := v.state.Memory[v.state.PC]
	v.state.PC++
	hi := v.state.Memory[v.state.PC]
	v.state.PC++

	a16 := uint16(hi)<<8 | uint16(lo)
	v.state.Memory[a16] = v.state.L
	v.state.Memory[a16+1] = v.state.H

	v.cycles -= 16
}

func (v *v1) INX_H() {
	result := (uint16(v.state.H)<<8 | uint16(v.state.L)) + 1
	v.state.H = byte(result >> 8)
	v.state.L = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_H() {
	result := v.state.H + 1
	v.setHalfCarryAdd(v.state.H, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.H = result
	v.cycles -= 5
}

func (v *v1) DCR_H() {
	result := v.state.H - 1
	v.setHalfCarrySub(v.state.H, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.H = result
	v.cycles -= 5
}

func (v *v1) MVI_H_D8() {
	v.state.H = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) DAA() {
	if (v.state.A&0x0F) > 9 || v.halfCarrySet() {
		v.setHalfCarryAdd(v.state.A, 6)
		a := uint16(v.state.A) + uint16(6)
		v.setCarry(a)
		v.state.A = byte(a)
	}

	if ((v.state.A&0xF0)>>4) > 9 || v.carrySet() {
		a := uint16(v.state.A) + uint16(0x60)
		v.setCarry(a)
		v.state.A = byte(a)
	}

	v.setSign(v.state.A)
	v.setParity(v.state.A)
	v.setZero(v.state.A)

	v.cycles -= 4
}

func (v *v1) DAD_H() {
	a := uint32(v.state.H)<<8 | uint32(v.state.L)
	r32 := a + a
	v.setCarry16(r32)
	r := uint16(r32)
	v.state.H = uint8(r >> 8)
	v.state.L = uint8(r & 0xFF)
	v.cycles -= 10
}

func (v *v1) LHLD_A16() {
	lo := v.state.Memory[v.state.PC]
	v.state.PC++
	hi := v.state.Memory[v.state.PC]
	v.state.PC++

	a16 := uint16(hi)<<8 | uint16(lo)
	v.state.L = v.state.Memory[a16]
	v.state.H = v.state.Memory[a16+1]

	v.cycles -= 16
}

func (v *v1) DCX_H() {
	result := (uint16(v.state.H)<<8 | uint16(v.state.L)) - 1
	v.state.H = byte(result >> 8)
	v.state.L = byte(result)
	v.cycles -= 5
}

func (v *v1) INR_L() {
	result := v.state.L + 1
	v.setHalfCarryAdd(v.state.L, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.L = result
	v.cycles -= 5
}

func (v *v1) DCR_L() {
	result := v.state.L - 1
	v.setHalfCarrySub(v.state.L, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.L = result
	v.cycles -= 5
}

func (v *v1) MVI_L_D8() {
	v.state.L = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) CMA() {
	v.state.A = ^v.state.A
	v.cycles -= 4
}

func (v *v1) LXI_SP_D16() {
	lo := v.state.Memory[v.state.PC]
	v.state.PC++
	hi := v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.SP = uint16(hi)<<8 | uint16(lo)
	v.cycles -= 10
}

func (v *v1) STA_A16() {
	lo := v.state.Memory[v.state.PC]
	v.state.PC++
	hi := v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.Memory[uint16(hi)<<8|uint16(lo)] = v.state.A
	v.cycles -= 13
}

func (v *v1) INX_SP() {
	v.state.SP++
	v.cycles -= 5
}

func (v *v1) INR_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	memval := v.state.Memory[memloc]
	result := memval + 1
	v.setHalfCarryAdd(memval, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.Memory[memloc] = result
	v.cycles -= 10
}

func (v *v1) DCR_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	memval := v.state.Memory[memloc]
	result := memval - 1
	v.setHalfCarrySub(memval, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.Memory[memloc] = result
	v.cycles -= 10
}

func (v *v1) MVI_M_D8() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	v.state.Memory[memloc] = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 10
}

func (v *v1) STC() {
	v.state.Flags |= carryFlag
	v.cycles -= 4
}

func (v *v1) DAD_SP() {
	a := uint32(v.state.SP)
	b := uint32(v.state.H)<<8 | uint32(v.state.L)
	r32 := a + b
	v.setCarry16(r32)
	r := uint16(r32)
	v.state.H = uint8(r >> 8)
	v.state.L = uint8(r & 0xFF)
	v.cycles -= 10
}

func (v *v1) LDA_A16() {
	lo := v.state.Memory[v.state.PC]
	v.state.PC++
	hi := v.state.Memory[v.state.PC]
	v.state.PC++
	v.state.A = v.state.Memory[uint16(hi)<<8|uint16(lo)]
	v.cycles -= 13
}

func (v *v1) DCX_SP() {
	v.state.SP--
	v.cycles -= 5
}

func (v *v1) INR_A() {
	result := v.state.A + 1
	v.setHalfCarryAdd(v.state.A, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.A = result
	v.cycles -= 5
}

func (v *v1) DCR_A() {
	result := v.state.A - 1
	v.setHalfCarrySub(v.state.A, 1)
	v.setParity(result)
	v.setZero(result)
	v.setSign(result)
	v.state.A = result
	v.cycles -= 5
}

func (v *v1) MVI_A_D8() {
	v.state.A = v.state.Memory[v.state.PC]
	v.state.PC++
	v.cycles -= 7
}

func (v *v1) CMC() {
	v.state.Flags ^= carryFlag
	v.cycles -= 4
}

func (v *v1) MOV_B_B() {
	v.cycles -= 5
}

func (v *v1) MOV_B_C() {
	v.state.B = v.state.C
	v.cycles -= 5
}

func (v *v1) MOV_B_D() {
	v.state.B = v.state.D
	v.cycles -= 5
}

func (v *v1) MOV_B_E() {
	v.state.B = v.state.E
	v.cycles -= 5
}

func (v *v1) MOV_B_H() {
	v.state.B = v.state.H
	v.cycles -= 5
}

func (v *v1) MOV_B_L() {
	v.state.B = v.state.L
	v.cycles -= 5
}

func (v *v1) MOV_B_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	v.state.B = v.state.Memory[memloc]
	v.cycles -= 7
}

func (v *v1) MOV_B_A() {
	v.state.B = v.state.A
	v.cycles -= 5
}

func (v *v1) MOV_C_B() {
	v.state.C = v.state.B
	v.cycles -= 5
}

func (v *v1) MOV_C_C() {
	v.cycles -= 5
}

func (v *v1) MOV_C_D() {
	v.state.C = v.state.D
	v.cycles -= 5
}

func (v *v1) MOV_C_E() {
	v.state.C = v.state.E
	v.cycles -= 5
}

func (v *v1) MOV_C_H() {
	v.state.C = v.state.H
	v.cycles -= 5
}

func (v *v1) MOV_C_L() {
	v.state.C = v.state.L
	v.cycles -= 5
}

func (v *v1) MOV_C_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	v.state.C = v.state.Memory[memloc]
	v.cycles -= 7
}

func (v *v1) MOV_C_A() {
	v.state.C = v.state.A
	v.cycles -= 5
}

func (v *v1) MOV_D_B() {
	v.state.D = v.state.B
	v.cycles -= 5
}

func (v *v1) MOV_D_C() {
	v.state.D = v.state.C
	v.cycles -= 5
}

func (v *v1) MOV_D_D() {
	v.cycles -= 5
}

func (v *v1) MOV_D_E() {
	v.state.D = v.state.E
	v.cycles -= 5
}

func (v *v1) MOV_D_H() {
	v.state.D = v.state.H
	v.cycles -= 5
}

func (v *v1) MOV_D_L() {
	v.state.D = v.state.L
	v.cycles -= 5
}

func (v *v1) MOV_D_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	v.state.D = v.state.Memory[memloc]
	v.cycles -= 7
}

func (v *v1) MOV_D_A() {
	v.state.D = v.state.A
	v.cycles -= 5
}

func (v *v1) MOV_E_B() {
	v.state.E = v.state.B
	v.cycles -= 5
}

func (v *v1) MOV_E_C() {
	v.state.E = v.state.C
	v.cycles -= 5
}

func (v *v1) MOV_E_D() {
	v.state.E = v.state.D
	v.cycles -= 5
}

func (v *v1) MOV_E_E() {
	v.cycles -= 5
}

func (v *v1) MOV_E_H() {
	v.state.E = v.state.H
	v.cycles -= 5
}

func (v *v1) MOV_E_L() {
	v.state.E = v.state.L
	v.cycles -= 5
}

func (v *v1) MOV_E_M() {
	memloc := uint16(v.state.H)<<8 | uint16(v.state.L)
	v.state.E = v.state.Memory[memloc]
	v.cycles -= 7
}

func (v *v1) MOV_E_A() {
	v.state.E = v.state.A
	v.cycles -= 5
}
