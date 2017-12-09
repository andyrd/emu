package v1

var parityLookup [0xFF]byte

func init() {
	initParityLookup()
}

func halfCarryAdd(a byte, b byte) byte {
	if (a&0x0F)+(b&0x0F) > 0x0F {
		return 1
	}
	return 0
}

func halfCarrySub(a byte, b byte) byte {
	if (a&0x0F)+(^b&0x0F)+1 > 0x0F {
		return 1
	}
	return 0
}

func sign(r byte) byte {
	return (r & 0x80) >> 7
}

func parity(r byte) byte {
	return parityLookup[r]
}

func zero(r byte) byte {
	if r == 0 {
		return 1
	}
	return 0
}

func carry16(r uint32) byte {
	if r > 0xFFFF {
		return 1
	}
	return 0
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
