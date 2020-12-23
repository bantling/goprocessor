package register

const (
	// SignBit8 is the sign bit for an 8 bit value
	SignBit8 uint64 = 0x0000000000000080

	// SignBit16 is the sign bit for a 16 bit value
	SignBit16 uint64 = 0x0000000000008000

	// SignBit32 is the sign bit for a 32 bit value
	SignBit32 uint64 = 0x0000000080000000

	// SignBit64 is the sign bit for a 64 bit value
	SignBit64 uint64 = 0x8000000000000000

	// SignBitExtend8 is the bit pattern to sign extend an 8 bit value
	SignBitExtend8 uint64 = 0xFFFFFFFFFFFFFF80

	// SignBitExtend16 is the bit pattern to sign extend a 16 bit value
	SignBitExtend16 uint64 = 0xFFFFFFFFFFFF8000

	// SignBitExtend32 is the bit pattern to sign extend a 32 bit value
	SignBitExtend32 uint64 = 0xFFFFFFFF80000000
)

// GeneralRegister defines a general register
// Can operate as a signed or unsigned 8, 16, 32, or 64 bit value.
type GeneralRegister uint64

// Negative returns true if the register is negative
func (r GeneralRegister) Negative() bool {
	return (uint64(r) & SignBit64) == SignBit64
}

// Uint8 returns register value as a uint8
func (r GeneralRegister) Uint8() uint8 {
	return uint8(r)
}

// Uint16 returns register value as a uint16
func (r GeneralRegister) Uint16() uint16 {
	return uint16(r)
}

// Uint32 returns register value as a uint32
func (r GeneralRegister) Uint32() uint32 {
	return uint32(r)
}

// Uint64 returns register value as a uint64
func (r GeneralRegister) Uint64() uint64 {
	return uint64(r)
}

// Int8 returns register value as an int8
func (r GeneralRegister) Int8() int8 {
	return int8(r)
}

// Int16 returns register value as an int16
func (r GeneralRegister) Int16() int16 {
	return int16(r)
}

// Int32 returns register value as an int32
func (r GeneralRegister) Int32() int32 {
	return int32(r)
}

// Int64 returns register value as an int64
func (r GeneralRegister) Int64() int64 {
	return int64(r)
}

// SetUint8 sets the register to the uint8 value.
// The sign is extended by copying bit 7 to bits 8 thru 63
func (r *GeneralRegister) SetUint8(val uint8) {
	val64 := uint64(val)
	if (val64 & SignBit8) == SignBit8 {
		*r = GeneralRegister(SignBitExtend8 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint16 sets the register to the given uint16 value
func (r *GeneralRegister) SetUint16(val uint16) {
	val64 := uint64(val)
	if (val64 & SignBit16) == SignBit16 {
		*r = GeneralRegister(SignBitExtend16 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint32 sets the register to the given uint32 value
func (r *GeneralRegister) SetUint32(val uint32) {
	val64 := uint64(val)
	if (val64 & SignBit32) == SignBit32 {
		*r = GeneralRegister(SignBitExtend32 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint64 sets the register to the given uint64 value
func (r *GeneralRegister) SetUint64(val uint64) {
	*r = GeneralRegister(val)
}

// SetInt8 sets the register to the given int8 value
func (r *GeneralRegister) SetInt8(val int8) {
	*r = GeneralRegister(val)
}

// SetInt16 sets the register to the given int16 value
func (r *GeneralRegister) SetInt16(val int16) {
	*r = GeneralRegister(val)
}

// SetInt32 sets the register to the given int32 value
func (r *GeneralRegister) SetInt32(val int32) {
	*r = GeneralRegister(val)
}

// SetInt64 sets the register to the given int64 value
func (r *GeneralRegister) SetInt64(val int64) {
	*r = GeneralRegister(val)
}
