// SPDX-License-Identifier: Apache-2.0

package register

import (
	"math/big"
)

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
	
	// ZeroHigherBits8 is the bit pattern to zero out upper 56 bits
	ZeroHigherBits8 uint64 = 0x00000000000000FF
	
	// ZeroHigherBits16 is the bit pattern to zero out upper 48 bits
	ZeroHigherBits16 uint64 = 0x000000000000FFFF
	
	// ZeroHigherBits32 is the bit pattern to zero out upper 32 bits
	ZeroHigherBits32 uint64 = 0x00000000FFFFFFFF

	// MaxUint8 is largest 8 bit unsigned value within a uint64, and a mask for only lowest 8 bits
	MaxUint8 uint64 = 0x00000000000000FF

	// MaxUint16 is largest 16 bit unsigned value within a uint64, and a mask for only lowest 16 bits
	MaxUint16 uint64 = 0x000000000000FFFF

	// MaxUint32 is largest 32 bit unsigned value within a uint64, and a mask for only lowest 32 bits
	MaxUint32 uint64 = 0x00000000FFFFFFFF

	// MaxUint64 is largest 64 bit unsigned value
	MaxUint64 uint64 = 0xFFFFFFFFFFFFFFFF

	// ErrDivisionByZero is returned by Divide method if the denominator is zero
	ErrDivisionByZero = GeneralRegisterError("Division By Zero")
)

// GeneralRegister defines a general register
// Can operate as a signed or unsigned 8, 16, 32, or 64 bit value
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
// The sign is extended by copying bit 7 to bits 8 thru 63.
func (r *GeneralRegister) SetUint8(val uint8) {
	val64 := uint64(val)
	if (val64 & SignBit8) == SignBit8 {
		*r = GeneralRegister(SignExtend8 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint16 sets the register to the given uint16 value.
// The sign is extended by copying bit 15 to bits 16 thru 63.
func (r *GeneralRegister) SetUint16(val uint16) {
	val64 := uint64(val)
	if (val64 & SignBit16) == SignBit16 {
		*r = GeneralRegister(SignExtend16 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint32 sets the register to the given uint32 value.
// The sign is extended by copying bit 31 to bits 32 thru 63.
func (r *GeneralRegister) SetUint32(val uint32) {
	val64 := uint64(val)
	if (val64 & SignBit32) == SignBit32 {
		*r = GeneralRegister(SignExtend32 | val64)
	} else {
		*r = GeneralRegister(val64)
	}
}

// SetUint64 sets the register to the given uint64 value.
// The sign does not need to be extended.
func (r *GeneralRegister) SetUint64(val uint64) {
	*r = GeneralRegister(val)
}

// SetInt8 sets the register to the given int8 value.
// The sign does not need to be extended.
func (r *GeneralRegister) SetInt8(val int8) {
	*r = GeneralRegister(val)
}

// SetInt16 sets the register to the given int16 value.
// The sign does not need to be extended.
func (r *GeneralRegister) SetInt16(val int16) {
	*r = GeneralRegister(val)
}

// SetInt32 sets the register to the given int32 value.
// The sign does not need to be extended.
func (r *GeneralRegister) SetInt32(val int32) {
	*r = GeneralRegister(val)
}

// SetInt64 sets the register to the given int64 value.
// The sign does not need to be extended.
func (r *GeneralRegister) SetInt64(val int64) {
	*r = GeneralRegister(val)
}

// Negative returns true if the register is negative
func (r GeneralRegister) Negative() bool {
	return (uint64(r) & SignBit64) == SignBit64
}

// ExtendSign extends the sign of the register, considering the current operand size
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
	
	// No need to do anything for 64 bit operands
	}
}

// ZeroHigherBits zeroes our bits higher than the current operand size
func (r *GeneralRegister) ZeroHigherBits(st StatusRegister) {
	switch st.OperandSize() {
	case STOperand8:
    	*r &= GeneralRegister(ZeroHigherBits8)

	case STOperand16:
    	*r &= GeneralRegister(ZeroHigherBits16)

	case STOperand32:
    	*r &= GeneralRegister(ZeroHigherBits32)
	
	// No need to do anything for 64 bit operands
	}
}

// AddInteger sets r = r + op + Carry using integer math, with the following side effects:
// - Result is sign extended
// - Carry is true if unsigned result < original value
// - Overflow is true if the two operands are the same sign and the result is the opposite sign
// - Zero is true if result is zero
// - Negative is true if result is negative
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
// - Result is sign extended
// - Zero is true if result is zero
// - Negative is true if result is negative
func (r *GeneralRegister) And(op GeneralRegister, st *StatusRegister) {
	*r &= op

	r.ExtendSign(*st)

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}

// Compare r with op, with the following side effects:
// - Carry is true if r >= op unsigned
// - Overflow is true if r >= op signed
// - Zero is true if r == op
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

// DivideIntegerSigned sets r = r / op and op = r % op using signed integer math, with the following side effects:
// - Zero is true if the result is zero
// - Negative is true if the quotient is negative
// - Overflow is true if the remainder is negative
//
// Returns ErrDivisionByZero if op = 0
func (r *GeneralRegister) DivideIntegerSigned(op *GeneralRegister, st *StatusRegister) error {
	if uint64(*op) == 0 {
		return ErrDivisionByZero
	}

	var (
		dividend  = r.Int64()
		divisor   = op.Int64()
		quotient  = dividend / divisor
		remainder = dividend % divisor
	)

	// Don't need to extend sign as we used int64
	*r = GeneralRegister(quotient)
	*op = GeneralRegister(remainder)

    st.Zero(*r == 0)
	st.Negative(r.Negative())
	st.Overflow(op.Negative())

	return nil
}

// MultiplyIntegerSigned sets r = r * op, with the following side effects:
// - Sign is extended
// - Zero is true if the result is zero
// - Negative is true if the result is negative
//
// If the operand size is 64 bits, r contains the lower 64 bits of the result,
// and op contains the upper 64 bits. Otherwise, only r contains the complete result.
func (r *GeneralRegister) MultiplyIntegerSigned(op *GeneralRegister, st *StatusRegister) {
	var (
		rVal  = r.Int64()
		opVal = op.Int64()
	)

	// If operand size is at most 32 bits, the result fits in 64 bits
	// Only r is affected
	if st.OperandSize() <= STOperand32 {
		r.SetInt64(rVal * opVal)
		st.Zero(*r == 0)
		st.Negative(r.Negative())
	} else {
		// Otherwise, the result may not fit into 64 bits
		// Both r and op are affected
		// r is lower 64 bits, op is upper 64 bits
		var (
			rInt  = big.NewInt(rVal)
			opInt = big.NewInt(opVal)
		)

		rInt = rInt.Mul(rInt, opInt)

		if rInt.IsInt64() {
			// Turns out result does fit into int64
			int64Val := rInt.Int64()

			r.SetInt64(int64Val)
			// Sign extend positive to upper 64
			op.SetInt64(0)

			isNeg := int64Val < 0
			if isNeg {
			    // Sign extend negative to upper 64
				op.SetInt64(-1)
			}

			st.Zero(int64Val == 0)
			st.Negative(isNeg)
		} else {
			// Result is at least 65 bits
			// Get absolute value bytes in order from most to least significant (little endian), contrary to docs
			bytes := rInt.Bytes()

			var (
				lastIndex = len(bytes) - 1
				low64Val  = uint64(bytes[lastIndex])
			)

			// Grab last 8 bytes into r
			for i := 1; i <= 7; i++ {
				low64Val = (low64Val << 8) + uint64(bytes[lastIndex-i])
			}

			// Grab remaining bytes into op (can't be more than 8 remaining)
			high64Val := uint64(bytes[lastIndex-8])
			for i := 9; lastIndex-i >= 0; i++ {
				high64Val = (high64Val << 8) + uint64(bytes[lastIndex-i])
			}

			// Check the sign and adjust if negative
			if rInt.Sign() < 0 {
				// Flip bits,using XOR (go has no logical NOT operator)
				low64Val ^= MaxUint64
				high64Val ^= MaxUint64

				// Increment lower64 bits.
				// If the result is 0, it wrapped around, increment upper 64 bits.
				low64Val++
				if low64Val == 0 {
					high64Val++
				}
			}

			*r = GeneralRegister(low64Val)
			*op = GeneralRegister(high64Val)
			st.Zero((*r == 0) && (*op == 0))
			st.Negative(high64Val >= SignBit64)
		}
	}
}

// Or r and op, with the following side effects:
// - Result is sign extended
// - Zero is true if result is zero
// - Negative is true if result is negative
func (r *GeneralRegister) Or(op GeneralRegister, st *StatusRegister) {
	*r |= op

	r.ExtendSign(*st)

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}

// Shift right arithmetic r and op, with the following side effects:
// - Result is sign extended
// - Zero is true if result is zero
// - Negative is true if result is negative
func (r *GeneralRegister) ShiftRightArithmetic(op GeneralRegister, st *StatusRegister) {
    if (r.Negative()) {
        // Go doesn't have arithmetic shift right, so manually shift, setting sign bit to 1 after each step
        // As an optimization, if the number of shifts is >= 64, the result must be -1
        numShift := uint64(op)
        if (numShift >= 64) {
            *r = GeneralRegister(MaxUint64) 
        } else {
            for ; numShift > 0; numShift-- {
                *r >>= 1
                *r |= GeneralRegister(SignBit64)
            }
        }
    } else {
        *r >>= op
    }

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}

// Shift left r and op, with the following side effects:
// - Result is sign extended
// - Zero is true if result is zero
// - Negative is true if result is negative
func (r *GeneralRegister) ShiftLeft(op GeneralRegister, st *StatusRegister) {
    *r <<= op
    r.ExtendSign(*st)

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}

// Shift right r and op, with the following side effects:
// - Result is sign extended
// - Zero is true if result is zero
// - Negative is true if result is negative
func (r *GeneralRegister) ShiftRight(op GeneralRegister, st *StatusRegister) {
    r.ZeroHigherBits(*st)
    *r >>= op
    r.ExtendSign(*st)

	st.Zero(*r == 0)
	st.Negative(r.Negative())
}
