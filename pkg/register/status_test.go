package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusRegister(t *testing.T) {
	var st = new(StatusRegister)

	// Carry
	*st = 0xFFFFFFFF
	assert.True(t, st.Carry())
	st.ClearCarry()
	assert.False(t, st.Carry())
	assert.Equal(t, StatusRegister(0x7FFFFFFF), *st)
	st.SetCarry()
	assert.True(t, st.Carry())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Overflow
	assert.True(t, st.Overflow())
	st.ClearOverflow()
	assert.False(t, st.Overflow())
	assert.Equal(t, StatusRegister(0xBFFFFFFF), *st)
	st.SetOverflow()
	assert.True(t, st.Overflow())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Zero
	assert.True(t, st.Zero())
	st.ClearZero()
	assert.False(t, st.Zero())
	assert.Equal(t, StatusRegister(0xDFFFFFFF), *st)
	st.SetZero()
	assert.True(t, st.Zero())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Negative
	assert.True(t, st.Negative())
	st.ClearNegative()
	assert.False(t, st.Negative())
	assert.Equal(t, StatusRegister(0xEFFFFFFF), *st)
	st.SetNegative()
	assert.True(t, st.Negative())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Address
	assert.Equal(t, PtrPtrIxOfs, st.AddressMode())
	st.SelectAddressMode(Ptr)
	assert.Equal(t, Ptr, st.AddressMode())
	assert.Equal(t, StatusRegister(0xF1FFFFFF), *st)
	st.SelectAddressMode(PtrOfs)
	assert.Equal(t, PtrOfs, st.AddressMode())
	assert.Equal(t, StatusRegister(0xF3FFFFFF), *st)
	st.SelectAddressMode(PtrIx)
	assert.Equal(t, PtrIx, st.AddressMode())
	assert.Equal(t, StatusRegister(0xF5FFFFFF), *st)
	st.SelectAddressMode(PtrIxOfs)
	assert.Equal(t, PtrIxOfs, st.AddressMode())
	assert.Equal(t, StatusRegister(0xF7FFFFFF), *st)
	st.SelectAddressMode(PtrPtr)
	assert.Equal(t, PtrPtr, st.AddressMode())
	assert.Equal(t, StatusRegister(0xF9FFFFFF), *st)
	st.SelectAddressMode(PtrPtrOfs)
	assert.Equal(t, PtrPtrOfs, st.AddressMode())
	assert.Equal(t, StatusRegister(0xFBFFFFFF), *st)
	st.SelectAddressMode(PtrPtrIx)
	assert.Equal(t, PtrPtrIx, st.AddressMode())
	assert.Equal(t, StatusRegister(0xFDFFFFFF), *st)
	st.SelectAddressMode(PtrPtrIxOfs)
	assert.Equal(t, PtrPtrIxOfs, st.AddressMode())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// InterruptDisable
	assert.True(t, st.InterruptDisable())
	st.ClearInterruptDisable()
	assert.False(t, st.InterruptDisable())
	assert.Equal(t, StatusRegister(0xFEFFFFFF), *st)
	st.SetInterruptDisable()
	assert.True(t, st.InterruptDisable())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Register
	assert.Equal(t, RegisterR3, st.Register())
	st.SelectRegister(RegisterR2)
	assert.Equal(t, RegisterR2, st.Register())
	assert.Equal(t, StatusRegister(0xFFBFFFFF), *st)
	st.SelectRegister(RegisterR1)
	assert.Equal(t, RegisterR1, st.Register())
	assert.Equal(t, StatusRegister(0xFF7FFFFF), *st)
	st.SelectRegister(RegisterR0)
	assert.Equal(t, RegisterR0, st.Register())
	assert.Equal(t, StatusRegister(0xFF3FFFFF), *st)
	st.SelectRegister(RegisterR3)
	assert.Equal(t, RegisterR3, st.Register())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Pointer Register Set
	assert.False(t, st.PointerRegisterSet0())
	st.SelectPointerRegisterSet0()
	assert.True(t, st.PointerRegisterSet0())
	assert.Equal(t, StatusRegister(0xFFDFFFFF), *st)
	st.SelectPointerRegisterSet1()
	assert.False(t, st.PointerRegisterSet0())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Counter Register Set
	assert.False(t, st.CounterRegisterSet0())
	st.SelectCounterRegisterSet0()
	assert.True(t, st.CounterRegisterSet0())
	assert.Equal(t, StatusRegister(0xFFEFFFFF), *st)
	st.SelectCounterRegisterSet1()
	assert.False(t, st.CounterRegisterSet0())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Operand Size
	assert.Equal(t, Operand64, st.OperandSize())
	st.SelectOperandSize(Operand32)
	assert.Equal(t, Operand32, st.OperandSize())
	assert.Equal(t, StatusRegister(0xFFFBFFFF), *st)
	st.SelectOperandSize(Operand16)
	assert.Equal(t, Operand16, st.OperandSize())
	assert.Equal(t, StatusRegister(0xFFF7FFFF), *st)
	st.SelectOperandSize(Operand8)
	assert.Equal(t, Operand8, st.OperandSize())
	assert.Equal(t, StatusRegister(0xFFF3FFFF), *st)
	st.SelectOperandSize(Operand64)
	assert.Equal(t, Operand64, st.OperandSize())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// Math Mode
	assert.Equal(t, MathFloat, st.MathMode())
	st.SelectMathMode(MathFixed)
	assert.Equal(t, MathFixed, st.MathMode())
	assert.Equal(t, StatusRegister(0xFFFEFFFF), *st)
	st.SelectMathMode(MathFractional)
	assert.Equal(t, MathFractional, st.MathMode())
	assert.Equal(t, StatusRegister(0xFFFDFFFF), *st)
	st.SelectMathMode(MathInteger)
	assert.Equal(t, MathInteger, st.MathMode())
	assert.Equal(t, StatusRegister(0xFFFCFFFF), *st)
	st.SelectMathMode(MathFloat)
	assert.Equal(t, MathFloat, st.MathMode())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// System bits
	assert.Equal(t, uint8(0xFF), st.System())
	st.SetSystem(0x00)
	assert.Equal(t, uint8(0x00), st.System())
	assert.Equal(t, StatusRegister(0xFFFF00FF), *st)
	st.SetSystem(0xFF)
	assert.Equal(t, uint8(0xFF), st.System())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)

	// User bits
	assert.Equal(t, uint8(0xFF), st.User())
	st.SetUser(0x00)
	assert.Equal(t, uint8(0x00), st.User())
	assert.Equal(t, StatusRegister(0xFFFFFF00), *st)
	st.SetUser(0xFF)
	assert.Equal(t, uint8(0xFF), st.User())
	assert.Equal(t, StatusRegister(0xFFFFFFFF), *st)
}
