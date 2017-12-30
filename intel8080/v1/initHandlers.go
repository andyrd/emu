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
	v.handlers[ops.LXI_D_D16] = v.LXI_D_D16
	v.handlers[ops.STAX_D] = v.STAX_D
	v.handlers[ops.INX_D] = v.INX_D
	v.handlers[ops.INR_D] = v.INR_D
	v.handlers[ops.DCR_D] = v.DCR_D
	v.handlers[ops.MVI_D_D8] = v.MVI_D_D8
	v.handlers[ops.RAL] = v.RAL
	v.handlers[ops.DAD_D] = v.DAD_D
	v.handlers[ops.LDAX_D] = v.LDAX_D
	v.handlers[ops.DCX_D] = v.DCX_D
	v.handlers[ops.INR_E] = v.INR_E
	v.handlers[ops.DCR_E] = v.DCR_E
	v.handlers[ops.MVI_E_D8] = v.MVI_E_D8
	v.handlers[ops.RAR] = v.RAR
	v.handlers[ops.LXI_H_D16] = v.LXI_H_D16
	v.handlers[ops.SHLD_A16] = v.SHLD_A16
	v.handlers[ops.INX_H] = v.INX_H
	v.handlers[ops.INR_H] = v.INR_H
	v.handlers[ops.DCR_H] = v.DCR_H
	v.handlers[ops.MVI_H_D8] = v.MVI_H_D8
	v.handlers[ops.DAA] = v.DAA
	v.handlers[ops.DAD_H] = v.DAD_H
	v.handlers[ops.LHLD_A16] = v.LHLD_A16
	v.handlers[ops.DCX_H] = v.DCX_H
	v.handlers[ops.INR_L] = v.INR_L
	v.handlers[ops.DCR_L] = v.DCR_L
	v.handlers[ops.MVI_L_D8] = v.MVI_L_D8
	v.handlers[ops.CMA] = v.CMA
	v.handlers[ops.LXI_SP_D16] = v.LXI_SP_D16
	v.handlers[ops.STA_A16] = v.STA_A16
	v.handlers[ops.INX_SP] = v.INX_SP
	v.handlers[ops.INR_M] = v.INR_M
	v.handlers[ops.DCR_M] = v.DCR_M
	v.handlers[ops.MVI_M_D8] = v.MVI_M_D8
	v.handlers[ops.STC] = v.STC
	v.handlers[ops.DAD_SP] = v.DAD_SP
}
