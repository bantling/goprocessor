// SPDX-License-Identifier: Apache-2.0

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

	// (23 << 31) * (2 << 32) = (23 << 31) * (1 << 33) = 23 << 64
	// 23 = 0001 0111 = 17
	// 23 << 31 = 1011 | 1000 0000 0000 0000 0000 0000 0000 0000 = 00 | 00 00 00 0B 80 00 00 00
	//  1 << 33 = 0010 | 0000 0000 0000 0000 0000 0000 0000 0000 = 00 | 00 00 00 02 00 00 00 00
	// 23 << 64 = 0001 0111 63 0s                                = 17 | 00 00 00 00 00 00 00 00
	r0.SetUint64(0x0000000B80000000)
	r1.SetUint64(0x0000000200000000)
	st.SelectOperandSize(Operand64)
	st.SetZero()
	st.SetNegative()
	r0.MultiplyIntegerSigned(r1, st)

	assert.Equal(t, uint64(0x0000000000000000), r0.Uint64())
	assert.Equal(t, uint64(0x0000000000000017), r1.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
}

func TestGeneralRegisterOr(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)

	// F0 | 0F
	r0.SetUint8(0xF0)
	r1.SetUint8(0x0F)
	r0.Or(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// F0 | 1F
	r0.SetUint8(0xF0)
	r1.SetUint8(0x1F)
	r0.Or(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 79 | AA = 0111 1001 | 1010 1010 = 1111 1011 = FB
	r0.SetUint8(0x79)
	r1.SetUint8(0xAA)
	r0.Or(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFB), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 24 | 42 = 0010 0100 | 0100 0010 = 0110 0110 = 66  
	r0.SetUint8(0x24)
	r1.SetUint8(0x42)
	r0.Or(*r1, st)

	assert.Equal(t, uint64(0x66), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 99 | 88
	r0.SetUint8(0x99)
	r1.SetUint8(0x88)
	r0.Or(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFF99), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
}

func TestGeneralRegisterShiftRightArithmetic(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)
	
	// 85 >>> 2 = 1000 0101 >>> 2 = 1110 0001 = E1 
	r0.SetUint8(0x85)
	r1.SetUint8(2)
	r0.ShiftRightArithmetic(*r1, st)
	
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFE1), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
	
	// 7F >>> 3 = 0111 1111 >>> 3 = 0000 1111 = 0F 
	r0.SetUint8(0x7F)
	r1.SetUint8(3)
	r0.ShiftRightArithmetic(*r1, st)
	
	assert.Equal(t, uint64(0x0F), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// 7F >>> 100 = 0
	r0.SetUint8(0x7F)
	r1.SetUint64(100)
	r0.ShiftRightArithmetic(*r1, st);
	
	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// F7 >>> 2 = 1111 0111 >>> 2 = 1111 1101 = FD
	r0.SetUint8(0xF7)
	r1.SetUint32(2)
	r0.ShiftRightArithmetic(*r1, st);
	
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFD), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
	
	// 85 >>> 1,000,000,000 = FF
	r0.SetUint8(0x85)
	r1.SetUint32(1000000000)
	r0.ShiftRightArithmetic(*r1, st);
	
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
}

func TestGeneralRegisterShiftLeft(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)
	
	// 85 << 2 = 1000 0101 << 2 = 1110 0001 0100 = E14 = 14 in 8 bit mode
	r0.SetUint8(0x85)
	r1.SetUint8(2)
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFF85), r0.Uint64())
	r0.ShiftLeft(*r1, st)
	
	assert.Equal(t, uint64(0x14), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// 7F << 3 = 0111 1111 << 3 = 0011 1111 1000 = 3F8 = F8 in 8 bit mode 
	r0.SetUint8(0x7F)
	r1.SetUint8(3)
	r0.ShiftLeft(*r1, st)
	
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFF8), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())
	
	// 7F << 100 = 0
	r0.SetUint8(0x7F)
	r1.SetUint64(100)
	r0.ShiftLeft(*r1, st);
	
	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())
}

func TestGeneralRegisterShiftRight(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)
	
	// 85 >> 2 = 1000 0101 >> 2 = 0010 0001 = 21 
	r0.SetUint8(0x85)
	r1.SetUint8(2)
	r0.ShiftRight(*r1, st)
	
	assert.Equal(t, uint64(0x21), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// 7F >> 3 = 0111 1111 >> 3 = 0000 1111 = 0F 
	r0.SetUint8(0x7F)
	r1.SetUint8(3)
	r0.ShiftRight(*r1, st)
	
	assert.Equal(t, uint64(0x0F), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// 7F >> 100 = 0
	r0.SetUint8(0x7F)
	r1.SetUint64(100)
	r0.ShiftRight(*r1, st);
	
	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// F7 >> 2 = 1111 0111 >> 2 = 0011 1101 = 3D
	r0.SetUint8(0xF7)
	r1.SetUint32(2)
	r0.ShiftRight(*r1, st);
	
	assert.Equal(t, uint64(0x3D), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
	
	// 85 >> 1,000,000,000 = 00
	r0.SetUint8(0x85)
	r1.SetUint32(1000000000)
	r0.ShiftRight(*r1, st);
	
	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())
}

func TestGeneralRegisterSubtractInteger(t *testing.T) {
	var (
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
		st = new(StatusRegister)
	)

	// 0 - 1
	r0.SetUint64(0)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, MaxUint64, r0.Uint64())
	assert.True(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// MaxUint64 - 1
	r0.SetUint64(MaxUint64)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, uint64(MaxUint64 - 1), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 7F - 1
	r0.SetUint8(0x7F)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, uint64(0x7E), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 80 - 1
	r0.SetUint8(0x80)
	r1.SetUint8(1)
	st.ClearCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, uint64(0x7F), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 80 - 1
	r0.SetUint8(0x80)
	r1.SetUint64(1)
	st.ClearCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, uint64(0x7F), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.True(t, st.IsOverflow())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 0x01 - 0 - C
	r0.SetUint16(0x01)
	r1.SetUint8(0)
	st.SetCarry()
	r0.SubtractInteger(*r1, st)

	assert.Equal(t, uint64(0x00), r0.Uint64())
	assert.False(t, st.IsCarry())
	assert.False(t, st.IsOverflow())
	assert.True(t, st.IsZero())
	assert.False(t, st.IsNegative())
}

func TestGeneralRegisterXor(t *testing.T) {
	var (
		st = new(StatusRegister)
		r0 = new(GeneralRegister)
		r1 = new(GeneralRegister)
	)

	// F0 ^ 0F = FFFFFFFFFFFFFFFF
	r0.SetUint8(0xF0)
	r1.SetUint8(0x0F)
	r0.Xor(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// F0 ^ 1F = FFFFFFFFFFFFFFEF
	r0.SetUint8(0xF0)
	r1.SetUint8(0x1F)
	r0.Xor(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFEF), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 79 ^ AA = 0111 1001 ^ 1010 1010 = 1101 0011 = D3
	r0.SetUint8(0x79)
	r1.SetUint8(0xAA)
	r0.Xor(*r1, st)

	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFD3), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.True(t, st.IsNegative())

	// 24 ^ 42 = 0010 0100 ^ 0100 0010 = 0110 0110 = 66  
	r0.SetUint8(0x24)
	r1.SetUint8(0x42)
	r0.Xor(*r1, st)

	assert.Equal(t, uint64(0x66), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())

	// 99 ^ 88 = 1001 1001 ^ 1000 1000 = 0001 0001
	r0.SetUint8(0x99)
	r1.SetUint8(0x88)
	r0.Xor(*r1, st)

	assert.Equal(t, uint64(0x11), r0.Uint64())
	assert.False(t, st.IsZero())
	assert.False(t, st.IsNegative())
}
