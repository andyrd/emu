package v1

var parityLookup [0xFF]byte

func init() {
	initParityLookup()
}

func (v *v1) halfCarryAdd(a byte, b byte) {
	if (a&0x0F)+(b&0x0F) > 0x0F {
		v.state.Flags |= halfCarryFlag
	} else {
		v.state.Flags &= ^halfCarryFlag
	}
}

func (v *v1) halfCarrySub(a byte, b byte) {
	if (a&0x0F)+(^b&0x0F)+1 > 0x0F {
		v.state.Flags |= halfCarryFlag
	} else {
		v.state.Flags &= ^halfCarryFlag
	}
}

func (v *v1) sign(r byte) {
	if r&0x80 == 0x80 {
		v.state.Flags |= signFlag
	} else {
		v.state.Flags &= ^signFlag
	}
}

func (v *v1) parity(r byte) {
	if parityLookup[r] == 1 {
		v.state.Flags |= parityFlag
	} else {
		v.state.Flags &= ^parityFlag
	}
}

func (v *v1) zero(r byte) {
	if r == 0 {
		v.state.Flags |= zeroFlag
	} else {
		v.state.Flags &= ^zeroFlag
	}
}

func (v *v1) carry16(r uint32) {
	if r > 0xFFFF {
		v.state.Flags |= carryFlag
	} else {
		v.state.Flags &= ^carryFlag
	}
}

func initParityLookup() {
	p2 := func(n byte) []byte {
		return []byte{n, n ^ 1, n ^ 1, n}
	}
	p4 := func(n byte) []byte {
		return append(append(append(p2(n), p2(n^1)...), p2(n^1)...), p2(n)...)
	}
	p6 := func(n byte) []byte {
		return append(append(append(p4(n), p4(n^1)...), p4(n^1)...), p4(n)...)
	}

	lookupSlice := append(append(append(p6(1), p6(0)...), p6(0)...), p6(1)...)
	copy(parityLookup[:], lookupSlice)
}
