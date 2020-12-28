package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralRegister(t *testing.T) {
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

	r.SetUint8(uint8(0x7F))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0x7F), r.Uint8())
	assert.Equal(t, uint16(0x007F), r.Uint16())
	assert.Equal(t, uint32(0x0000007F), r.Uint32())
	assert.Equal(t, uint64(0x000000000000007F), r.Uint64())
	assert.Equal(t, int8(0x7F), r.Int8())
	assert.Equal(t, int16(0x7F), r.Int16())
	assert.Equal(t, int32(0x7F), r.Int32())
	assert.Equal(t, int64(0x7F), r.Int64())

	r.SetUint8(uint8(0xFF))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint16(uint16(0x7FFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0x7FFF), r.Uint16())
	assert.Equal(t, uint32(0x00007FFF), r.Uint32())
	assert.Equal(t, uint64(0x0000000000007FFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(0x7FFF), r.Int16())
	assert.Equal(t, int32(0x00007FFF), r.Int32())
	assert.Equal(t, int64(0x0000000000007FFF), r.Int64())

	r.SetUint16(uint16(0xFFFF))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint32(uint32(0x7FFFFFFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0x7FFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x000000007FFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(0x7FFFFFFF), r.Int32())
	assert.Equal(t, int64(0x000000007FFFFFFF), r.Int64())

	r.SetUint32(uint32(0xFFFFFFFF))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetUint64(uint64(0x7FFFFFFFFFFFFFFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x7FFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFFFFFFFFFF), r.Int64())

	r.SetUint64(uint64(0xFFFFFFFFFFFFFFFF))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt8(int8(0x7F))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0x7F), r.Uint8())
	assert.Equal(t, uint16(0x007F), r.Uint16())
	assert.Equal(t, uint32(0x0000007F), r.Uint32())
	assert.Equal(t, uint64(0x000000000000007F), r.Uint64())
	assert.Equal(t, int8(0x7F), r.Int8())
	assert.Equal(t, int16(0x7F), r.Int16())
	assert.Equal(t, int32(0x7F), r.Int32())
	assert.Equal(t, int64(0x7F), r.Int64())

	r.SetInt8(int8(-1))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt16(int16(0x7FFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0x7FFF), r.Uint16())
	assert.Equal(t, uint32(0x00007FFF), r.Uint32())
	assert.Equal(t, uint64(0x0000000000007FFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(0x7FFF), r.Int16())
	assert.Equal(t, int32(0x7FFF), r.Int32())
	assert.Equal(t, int64(0x7FFF), r.Int64())

	r.SetInt16(int16(-1))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt32(int32(0x7FFFFFFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0x7FFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x000000007FFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(0x7FFFFFFF), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFF), r.Int64())

	r.SetInt32(int32(-1))
	assert.True(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(-1), r.Int64())

	r.SetInt64(int64(0x7FFFFFFFFFFFFFFF))
	assert.False(t, r.Negative())
	assert.Equal(t, uint8(0xFF), r.Uint8())
	assert.Equal(t, uint16(0xFFFF), r.Uint16())
	assert.Equal(t, uint32(0xFFFFFFFF), r.Uint32())
	assert.Equal(t, uint64(0x7FFFFFFFFFFFFFFF), r.Uint64())
	assert.Equal(t, int8(-1), r.Int8())
	assert.Equal(t, int16(-1), r.Int16())
	assert.Equal(t, int32(-1), r.Int32())
	assert.Equal(t, int64(0x7FFFFFFFFFFFFFFF), r.Int64())

	r.SetInt64(int64(-1))
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