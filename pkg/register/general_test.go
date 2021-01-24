package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralRegisterBasics(t *testing.T) {
	var r = new(GeneralRegister)

	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0x00), r.Uint8())
	assert.Equal(t, uint16(0x0000), r.Uint16())
	assert.Equal(t, uint32(0x00000000), r.Uint32())
	assert.Equal(t, uint64(0x0000000000000000), r.Uint64())
	assert.Equal(t, int8(0), r.Int8())
	assert.Equal(t, int16(0), r.Int16())
	assert.Equal(t, int32(0), r.Int32())
	assert.Equal(t, int64(0), r.Int64())

	r.SetUint8(0x7F)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0x7F), r.Uint8())
	assert.Equal(t, uint16(0x007F), r.Uint16())
	assert.Equal(t, uint32(0x0000007F), r.Uint32())
	assert.Equal(t, uint64(0x000000000000007F), r.Uint64())
	assert.Equal(t, int8(0x7F), r.Int8())
	assert.Equal(t, int16(0x7F), r.Int16())
	assert.Equal(t, int32(0x7F), r.Int32())
	assert.Equal(t, int64(0x7F), r.Int64())

	r.SetUint8(0xFF)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint16(0x7FFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0x7FFF), r.Uint16())
	assert.Equal(t, uint32(0x00007FFF), r.Uint32())
	assert.Equal(t, uint64(0x0000000000007FFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(0x7FFF), r.Int16())
	assert.Equal(t, int32(0x00007FFF), r.Int32())
	assert.Equal(t, int64(0x0000000000007FFF), r.Int64())

	r.SetUint16(0xFFFF)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint32(0x7FFFFFFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0x7FFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x000000007FFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(0x7FFFFFFF), r.Int32())
	assert.Equal(t, int64(0x000000007FFFFFFF), r.Int64())

	r.SetUint32(0xFFFFFFFF)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint64(0x7FFFFFFFFFFFFFFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x7FFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFFFFFFFFFF), r.Int64())

	r.SetUint64(0xFFFFFFFFFFFFFFFF)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt8(0x7F)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0x7F), r.Uint8())
	assert.Equal(t, uint16(0x007F), r.Uint16())
	assert.Equal(t, uint32(0x0000007F), r.Uint32())
	assert.Equal(t, uint64(0x000000000000007F), r.Uint64())
	assert.Equal(t, int8(0x7F), r.Int8())
	assert.Equal(t, int16(0x7F), r.Int16())
	assert.Equal(t, int32(0x7F), r.Int32())
	assert.Equal(t, int64(0x7F), r.Int64())

	r.SetInt8(-1)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt16(0x7FFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0x7FFF), r.Uint16())
	assert.Equal(t, uint32(0x00007FFF), r.Uint32())
	assert.Equal(t, uint64(0x0000000000007FFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(0x7FFF), r.Int16())
	assert.Equal(t, int32(0x7FFF), r.Int32())
	assert.Equal(t, int64(0x7FFF), r.Int64())

	r.SetInt16(-1)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt32(0x7FFFFFFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0x7FFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x000000007FFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(0x7FFFFFFF), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFF), r.Int64())

	r.SetInt32(-1)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt64(0x7FFFFFFFFFFFFFFF)
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x7FFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFFFFFFFFFF), r.Int64())

	r.SetInt64(-1)
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())
}

func TestGeneralRegisterAddInteger(t *testing.T) {
	var (
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
		st = new(StatusRegister)
	)

	// MaxUint64 - 1 + 1
	r0.SetUint64(MaxUint64 - 1)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, MaxUint64, r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// MaxUint64 + 1
	r0.SetUint64(MaxUint64)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, uint64(0), r0.Uint64())
	assert.True(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 7E + 1
	r0.SetUint8(0x7E)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, uint64(0x7F), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 7F + 1
	r0.SetUint8(0x7F)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFF80), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 80 + -1
	r0.SetUint8(0x80)
	r1.SetUint64(MaxUint64)
	st.ClearCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, uint64(0x7F), r0.Uint64())
	assert.True(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 0x01 + 0 + C
	r0.SetUint16(0x01)
	r1.SetUint8(0)
	st.SetCarry()
	r0.AddInteger(*r1, st)

	assert.Equal(t, uint64(0x02), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
}

func TestGeneralRegisterAnd(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)

	// F0 ^ 0F
	r0.SetUint8(0xF0)
	r1.SetUint8(0x0F)
	r0.And(*r1, st)

	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// F0 ^ 1F
	r0.SetUint8(0xF0)
	r1.SetUint8(0x1F)
	r0.And(*r1, st)

	assert.Equal(t, uint64(0x10), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 79 ^ AA
	r0.SetUint8(0x79)
	r1.SetUint8(0xAA)
	r0.And(*r1, st)

	assert.Equal(t, uint64(0x28), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 24 ^ 42
	r0.SetUint8(0x24)
	r1.SetUint8(0x42)
	r0.And(*r1, st)

	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 99 ^ 88
	r0.SetUint8(0x99)
	r1.SetUint8(0x88)
	r0.And(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFF88), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
}

func TestGeneralRegisterCompare(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)

	// 0x01 <=> 0x00
	r0.SetUint8(0x01)
	r1.SetUint8(0x00)
	r0.Compare(*r1, st)

	assert.True(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())

	// 0x01 <=> 0x01
	r0.SetUint8(0x01)
	r1.SetUint8(0x01)
	r0.Compare(*r1, st)

	assert.True(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.True(t, st.IsZero())

	// 0x01 <=> 0x02
	r0.SetUint8(0x01)
	r1.SetUint8(0x02)
	r0.Compare(*r1, st)

	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())

	// 0x01 <=> 0x80
	r0.SetUint8(0x01)
	r1.SetUint8(0x80)
	r0.Compare(*r1, st)

	assert.False(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())

	// 0x7F <=> 0x7E
	r0.SetUint8(0x7F)
	r1.SetUint8(0x7E)
	r0.Compare(*r1, st)

	assert.True(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())

	// 0x7F <=> 0x80
	r0.SetUint8(0x7F)
	r1.SetUint8(0x80)
	r0.Compare(*r1, st)

	assert.False(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())

	// 0x80 <=> 0x7F
	r0.SetUint8(0x80)
	r1.SetUint8(0x7F)
	r0.Compare(*r1, st)

	assert.True(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
}

func TestGeneralRegisterDivideInteger(t *testing.T) {
	var (
		r0  = new(GeneralRegister)
		r1  = new(GeneralRegister)
		st  = new(StatusRegister)
		err error
	)

	// 7 / 2 = 3 r 1
	r0.SetUint8(0x07)
	r1.SetUint8(0x02)
	st.ClearNegative()
	st.ClearOverflow()
	err = r0.DivideIntegerSigned(r1, st)

	assert.Nil(t, err)
	assert.Equal(t, uint64(0x0000000000000003), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000001), r1.Uint64())
	assert.False(t, st.IsNegative())
	assert.False(t, st.IsOverflow())

	// -7 / 2 = -3 r -1
	r0.SetUint8(0xF9)
	r1.SetUint8(0x02)
	st.ClearNegative()
	st.ClearOverflow()
	err = r0.DivideIntegerSigned(r1, st)

	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFD), r0.Uint64())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r1.Uint64())
	assert.True(t, st.IsNegative())
	assert.True(t, st.IsOverflow())

	// 7 / -2 = -3 r 1
	r0.SetUint8(0x07)
	r1.SetUint8(0xFE)
	st.ClearNegative()
	st.ClearOverflow()
	err = r0.DivideIntegerSigned(r1, st)

	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFD), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000001), r1.Uint64())
	assert.True(t, st.IsNegative())
	assert.False(t, st.IsOverflow())

	// -7 / -2 = 3 r -1
	r0.SetUint8(0xF9)
	r1.SetUint8(0xFE)
	st.ClearNegative()
	st.ClearOverflow()
	err = r0.DivideIntegerSigned(r1, st)

	assert.Nil(t, err)
	assert.Equal(t, uint64(0x0000000000000003), r0.Uint64())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r1.Uint64())
	assert.False(t, st.IsNegative())
	assert.True(t, st.IsOverflow())

	// 7 / 0 = division by zero error
	r0.SetUint8(0x07)
	r1.SetUint8(0x00)
	st.ClearNegative()
	st.ClearOverflow()
	err = r0.DivideIntegerSigned(r1, st)

	assert.Equal(t, ErrDivisionByZero, err)
	assert.Equal(t, uint64(0x0000000000000007), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000000), r1.Uint64())
	assert.False(t, st.IsNegative())
	assert.False(t, st.IsOverflow())
}

func TestGeneralRegisterMultiplyInteger(t *testing.T) {
	var (
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
		st = new(StatusRegister)
	)

	// 7 * 0 = 0
	r0.SetUint8(0x07)
	r1.SetUint8(0x00)
	st.SelectOperandSize(Operand8)
	st.ClearZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x0000000000000000), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000000), r1.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 7 * 2 = 14
	r0.SetUint8(0x07)
	r1.SetUint8(0x02)
	st.SelectOperandSize(Operand8)
	st.SetZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x000000000000000E), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000002), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// -7 * 2 = -14
	r0.SetUint8(0xF9)
	r1.SetUint8(0x02)
	st.SelectOperandSize(Operand8)
	st.SetZero()
	st.ClearNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFF2), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000002), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 7 * -2 = -14
	r0.SetUint8(0x07)
	r1.SetUint8(0xFE)
	st.SelectOperandSize(Operand8)
	st.SetZero()
	st.ClearNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFF2), r0.Uint64())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFE), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// -7 * -2 = 14
	r0.SetUint8(0xF9)
	r1.SetUint8(0xFE)
	st.SelectOperandSize(Operand8)
	st.SetZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x000000000000000E), r0.Uint64())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFE), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// (7 << 8) * (2 << 8) = 14 << 16
	r0.SetUint64(0x0000000000000700)
	r1.SetUint64(0x0000000000000200)
	st.SelectOperandSize(Operand64)
	st.SetZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x00000000000E0000), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000000), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// (23 << 31) * (2 << 32) = 46 << 63
	// 23 = 0001 0111 = 17 
	// 46 = 0010 1110 = 2E
	//            0000 | 0000 0000 0000 0000 0000 0000 0000 0000
	// 23 << 31 = 1011 | 1000 0000 0000 0000 0000 0000 0000 0000 = B 80 00 00 00
	//  2 << 32 = 0010 | 0000 0000 0000 0000 0000 0000 0000 0000 = 2 00 00 00 00
	// 46 << 63 = 0010 1110 63 0s = 0101 1100 60 0s              = 
	r0.SetUint64(0x00000005C0000000)
	r1.SetUint64(0x0000000200000000)
	st.SelectOperandSize(Operand64)
	st.SetZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x0000000000000000), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000000), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
}
