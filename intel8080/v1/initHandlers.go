package v1

import ops "github.com/andyrd/emu/intel8080"

func (v *v1) initHandlers() {
	v.handlers[ops.NOP] = v.NOP
	v.handlers[ops.LXI_B_D16] = v.LXI_B_D16
	v.handlers[ops.STAX_B] = v.STAX_B
	v.handlers[ops.INX_B] = v.INX_B
	v.handlers[ops.INR_B] = v.INR_B
	v.handlers[ops.DCR_B] = v.DCR_B
	v.handlers[ops.MVI_B_D8] = v.MVI_B_D8
	v.handlers[ops.RLC] = v.RLC
	v.handlers[ops.DAD_B] = v.DAD_B
	v.handlers[ops.LDAX_B] = v.LDAX_B
	v.handlers[ops.DCX_B] = v.DCX_B
	v.handlers[ops.INR_C] = v.INR_C
	v.handlers[ops.DCR_C] = v.DCR_C
	v.handlers[ops.MVI_C_D8] = v.MVI_C_D8
	v.handlers[ops.RRC] = v.RRC
}
