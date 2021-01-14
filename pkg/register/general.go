package register

// GeneralRegisterError represents an error performing operations on general registers
type GeneralRegisterError string

func (e GeneralRegisterError) Error() string {
	return string(e)
}

const (
	// SignBit8 is the sign bit for an 8 bit value
	SignBit8 uint64 = 0x0000000000000080

	// SignBit16 is the sign bit for a 16 bit value
	SignBit16 uint64 = 0x0000000000008000

	// SignBit32 is the sign bit for a 32 bit value
	SignBit32 uint64 = 0x0000000080000000

	// SignBit64 is the sign bit for a 64 bit value
	SignBit64 uint64 = 0x8000000000000000

	// SignExtend8 is the bit pattern to sign extend an 8 bit value
	SignExtend8 uint64 = 0xFFFFFFFFFFFFFF80

	// SignExtend16 is the bit pattern to sign extend a 16 bit value
	SignExtend16 uint64 = 0xFFFFFFFFFFFF8000

	// SignExtend32 is the bit pattern to sign extend a 32 bit value
	SignExtend32 uint64 = 0xFFFFFFFF80000000

	// MaxUint8 is largest 8 bit unsigned value within a uint64, and a mask for only lowest 8 bits
	MaxUint8 uint64 = 0x00000000000000FF

	// MaxUint16 is largest 16 bit unsigned value within a uint64, and a mask for only lowest 16 bits
	MaxUint16 uint64 = 0x000000000000FFFF

	// MaxUint32 is largest 32 bit unsiged value within a uint64, and a mask for only lowest 32 bits
	MaxUint32 uint64 = 0x00000000FFFFFFFF

	// MaxUint64 is largest 64 bit unsiged value
	MaxUint64 uint64 = 0xFFFFFFFFFFFFFFFF

	// ErrDivisionByZero is returned by Divide method if the denomoinator is zero
	ErrDivisionByZero = GeneralRegisterError("Division By Zero")
)

// GeneralRegister defines a general register
// Can operate as a signed or unsigned 8, 16, 32, or 64 bit value.
type GeneralRegister uint64

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
		*r = GeneralRegister(SignExtend8 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint16 sets the register to the given uint16 value
func (r *GeneralRegister) SetUint16(val uint16) {
	val64 := uint64(val)
	if (val64 & SignBit16) == SignBit16 {
		*r = GeneralRegister(SignExtend16 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint32 sets the register to the given uint32 value
func (r *GeneralRegister) SetUint32(val uint32) {
	val64 := uint64(val)
	if (val64 & SignBit32) == SignBit32 {
		*r = GeneralRegister(SignExtend32 | val64)
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

// Negative returns true if the register is negative
func (r GeneralRegister) Negative() bool {
	return (uint64(r) & SignBit64) == SignBit64
}

// ExtendSign extends the sign of the vregister, considering the current operand size
func (r *GeneralRegister) ExtendSign(st StatusRegister) {
	val := uint64(*r)
	switch st.OperandSize() {
	case STOperand8:
		if (val & SignBit8) == SignBit8 {
			*r = GeneralRegister(val | SignExtend8)
		} else {
			*r = GeneralRegister(val & MaxUint8)
		}

	case STOperand16:
		if (val & SignBit16) == SignBit16 {
			*r = GeneralRegister(val | SignExtend16)
		} else {
			*r = GeneralRegister(val & MaxUint16)
		}

	case STOperand32:
		if (val & SignBit32) == SignBit32 {
			*r = GeneralRegister(val | SignExtend32)
		} else {
			*r = GeneralRegister(val & MaxUint32)
		}
	}
}

// AddInteger sets r = r + op + Carry using integer math, with the following side effects:
// Result is sign extended
// Carry is true if unsigned result < original value
// Overflow is true if the two operands are the same sign and the result is the opposite sign
// Zero is true if result is zero
// Negative is true if result is negative
func (r *GeneralRegister) AddInteger(op GeneralRegister, st *StatusRegister) {
	oldVal := *r
	oldNeg := oldVal.Negative()

	*r += op
	if st.IsCarry() {
		*r++
	}

	r.ExtendSign(*st)
	newNeg := r.Negative()

	st.Carry(*r < oldVal)
	st.Overflow((oldNeg == op.Negative()) && (oldNeg != newNeg))
	st.Zero(*r == 0)
	st.Negative(newNeg)
}

// And r and op, with the following side effects:
// Result is sign extended
// Zero is true if result is zero
// Negative is true if result is negative
func (r *GeneralRegister) And(op GeneralRegister, st *StatusRegister) {
	*r &= op

	r.ExtendSign(*st)

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}

// Compare r with op, with the following side effects:
// Carry is true if r >= op unsigned
// Overflow is true if r >= op signed
// Zero is true if r == op
func (r *GeneralRegister) Compare(op GeneralRegister, st *StatusRegister) {
	// Unsigned >= is simple
	st.Carry(*r >= op)

	// Signed >= is trickier (using 8 bit examples):
	// 2 > 1
	// 1 < 255, but 255 is -1, so 1 > 255
	// 127 < 128, but 128 is -128, so 127 > 128
	// r > op if:
	// a) r + and op -
	// b) r and op are same sign and r > op
	rNeg := r.Negative()
	opNeg := op.Negative()
	st.Overflow((!rNeg && opNeg) || ((rNeg == opNeg) && (*r >= op)))

	// Unsigned and signed == is simple
	st.Zero(*r == op)
}

// DivideInteger sets r = r / op and op = r % op using integer math, with the following side effects:
// Results in r and op are sign extended
// Returns ErrDivisionByZero if op = 0
func (r *GeneralRegister) DivideInteger(op *GeneralRegister, st *StatusRegister) error {
	if uint64(*op) == 0 {
		return ErrDivisionByZero
	}

	return nil
}
