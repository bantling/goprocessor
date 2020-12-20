package proc

import (
	"github.com/bantling/gofuncs"
)

const (
	// DefaultSB is the default stack base
	DefaultSB uint32 = 0xFFFE0000

	// DefaultSP is the default stack pointer
	DefaultSP uint16 = 0xFFFF

	// STCarry is ST filter for carry flag
	STCarry uint32 = 0x80000000

	// STOverflow is ST filter for overflow flag
	STOverflow uint32 = 0x40000000

	// STZero is ST filter for zero flag
	STZero uint32 = 0x20000000

	// STNegative is ST filter for negative flag
	STNegative uint32 = 0x10000000

	// STAddress is ST filter for address mode
	STAddress uint32 = 0x0E000000

	// STAddressShift is ST shift for address
	STAddressShift int = 25

	// STInterruptDisable is ST filter for interrupt disable
	STInterruptDisable uint32 = 0x01000000

	// STRegister is ST filter for register
	STRegister uint32 = 0x00C00000

	// STRegisterShift is the ST shift for register
	STRegisterShift int = 22

	// STPointerRegisterSet is ST filter for pointer register set
	STPointerRegisterSet uint32 = 0x00200000

	// STCounterRegisterSet is ST filter for counter register set
	STCounterRegisterSet uint32 = 0x00100000

	// STOperand is ST filter for operand size
	STOperand uint32 = 0x000C0000

	// STOperandShift is the ST shift for operand
	STOperandShift int = 18

	// STMath is ST filter for math mode
	STMath uint32 = 0x00030000

	// STMathShift is the ST shift for math mode
	STMathShift int = 16

	// STSystem is ST filter for system bits
	STSystem uint32 = 0x0000FF00

	// STSystemShift is the ST shift for system bits
	STSystemShift int = 8

	// STUser is ST filter for user bits
	STUser uint32 = 0x000000FF
)

// Address modes
const (
	Ptr uint8 = iota
	PtrOfs
	PtrIx
	PtrIxOfs
	PtrPtr
	PtrPtrOfs
	PtrPtrIx
	PtrPtrIxOfs
)

// STAddressModeErr ifaAddress mode is invalid
const STAddressModeErr = "Address mode must be <= 7"

// Register selection
const (
	RegisterR0 uint8 = iota
	RegisterR1
	RegisterR2
	RegisterR3
)

// STRegisterErr if register is invalid
const STRegisterErr = "Register must be <= 3"

// Operand sizes
const (
	Operand8 uint8 = iota
	Operand16
	Operand32
	Operand64
)

// STOperandModeErr if operand mode is invalid
const STOperandModeErr = "Operand mode must be <= 3"

// Math modes
const (
	MathInteger uint8 = iota
	MathFractional
	MathFixed
	MathFloat
)

// STMathModeErr if math mode is invalid
const STMathModeErr = "Math mode must be <= 3"

// StatusRegister defines the status register
// CVZNAAAI RRPTOOMM SSSSSSSS UUUUUUUU
// Carry, oVerflow, Zero, Negative, Address mode, Interrupt Disable,
// Register, Pointer register, counTer register, Operand size, Math mode,
// 8 bits reserved for system use, and 8 bits reserved for users.
type StatusRegister uint32

// Carry returns true if the carry flag is set
func (st StatusRegister) Carry() bool {
	return (uint32(st) & STCarry) == STCarry
}

// SetCarry sets the carry flag
func (st StatusRegister) SetCarry() {
	st |= StatusRegister(STCarry)
}

// ClearCarry clears the carry flag
func (st StatusRegister) ClearCarry() {
	st = StatusRegister((uint32(st) | STCarry) - STCarry)
}

// Overflow returns true if the overflow flag is set
func (st StatusRegister) Overflow() bool {
	return (uint32(st) & STOverflow) == STOverflow
}

// SetOverflow sets the overflow flag
func (st StatusRegister) SetOverflow() {
	st |= StatusRegister(STOverflow)
}

// ClearOverflow clears the overflow flag
func (st StatusRegister) ClearOverflow() {
	st = StatusRegister((uint32(st) | STOverflow) - STOverflow)
}

// Negative returns true if the negative flag is set
func (st StatusRegister) Negative() bool {
	return (uint32(st) & STNegative) == STNegative
}

// SetNegative sets the negative flag
func (st StatusRegister) SetNegative() {
	st |= StatusRegister(STNegative)
}

// ClearNegative clears the negative flag
func (st StatusRegister) ClearNegative() {
	st = StatusRegister((uint32(st) | STNegative) - STNegative)
}

// AddressMode returns the address mode
func (st StatusRegister) AddressMode() uint8 {
	return uint8((uint32(st) & STAddress) >> STAddressShift)
}

// SetAddressMode sets the address mode
func (st StatusRegister) SetAddressMode(am uint8) {
	gofuncs.PanicBM(am <= PtrPtrIxOfs, STAddressModeErr)

	st = StatusRegister(((uint32(st) | STAddress) - STAddress) + (uint32(am) << STAddressShift))
}

// InterruptDisable returns true if the interrupt disable flag is set
func (st StatusRegister) InterruptDisable() bool {
	return (uint32(st) & STInterruptDisable) == STInterruptDisable
}

// SetInterruptDisable sets the interrupt disable flag
func (st StatusRegister) SetInterruptDisable() {
	st |= StatusRegister(STInterruptDisable)
}

// ClearInterruptDisable clears the interrupt disable flag
func (st StatusRegister) ClearInterruptDisable() {
	st = StatusRegister((uint32(st) | STInterruptDisable) - STInterruptDisable)
}

// RegisterMode returns the register mode
func (st StatusRegister) RegisterMode() uint8 {
	return uint8((uint32(st) & STRegister) >> STRegisterShift)
}

// SetRegisterMode sets the register mode
func (st StatusRegister) SetRegisterMode(rm uint8) {
	gofuncs.PanicBM(rm <= Operand64, STRegisterErr)

	st = StatusRegister(((uint32(st) | STRegister) - STRegister) + (uint32(rm) << STRegisterShift))
}

// PointerRegisterSet0 returns true if pointer register set 0 is selected
func (st StatusRegister) PointerRegisterSet0() bool {
	return (uint32(st) & STPointerRegisterSet) == 0
}

// PointerRegisterSet1 returns true if pointer register set 1 is selected
func (st StatusRegister) PointerRegisterSet1() bool {
	return (uint32(st) & STPointerRegisterSet) == STPointerRegisterSet
}

// SetPointerRegisterSet0 selects pointer register set 0
func (st StatusRegister) SetPointerRegisterSet0() {
	st = StatusRegister((uint32(st) | STPointerRegisterSet) - STPointerRegisterSet)
}

// SetPointerRegisterSet1 selects pointer register set 1
func (st StatusRegister) SetPointerRegisterSet1() {
	st |= StatusRegister(STPointerRegisterSet)
}

// CounterRegisterSet0 returns true if counter register set 0 is selected
func (st StatusRegister) CounterRegisterSet0() bool {
	return (uint32(st) & STCounterRegisterSet) == 0
}

// CounterRegisterSet1 returns true if counter register set 1 is selected
func (st StatusRegister) CounterRegisterSet1() bool {
	return (uint32(st) & STCounterRegisterSet) == STCounterRegisterSet
}

// SetCounterRegisterSet0 selects counter register set 0
func (st StatusRegister) SetCounterRegisterSet0() {
	st = StatusRegister((uint32(st) | STCounterRegisterSet) - STCounterRegisterSet)
}

// SetCounterRegisterSet1 selects counter register set 1
func (st StatusRegister) SetCounterRegisterSet1() {
	st |= StatusRegister(STCounterRegisterSet)
}

// OperandSize returns the operand size
func (st StatusRegister) OperandSize() uint8 {
	return uint8((uint32(st) & STOperand) >> STOperandShift)
}

// SetOperandSize sets the operand size
func (st StatusRegister) SetOperandSize(om uint8) {
	gofuncs.PanicBM(om <= Operand64, STOperandModeErr)

	st = StatusRegister(((uint32(st) | STOperand) - STOperand) + (uint32(om) << STOperandShift))
}

// MathMode returns the math mode
func (st StatusRegister) MathMode() uint8 {
	return uint8((uint32(st) & STMath) >> STMathShift)
}

// SetMathMode sets the math mode
func (st StatusRegister) SetMathMode(mm uint8) {
	gofuncs.PanicBM(mm <= MathFloat, STMathModeErr)

	st = StatusRegister(((uint32(st) | STMath) - STMath) + (uint32(mm) << STMathShift))
}

// System returns the system defined ST bits
func (st StatusRegister) System() uint8 {
	return uint8((uint32(st) & STSystem) >> STSystemShift)
}

// SetSystem sets the system defined ST bits
func (st StatusRegister) SetSystem(sb uint8) {
	st = StatusRegister(((uint32(st) | STSystem) - STSystem) + (uint32(sb) << STSystemShift))
}

// User returns the user defined ST bits
func (st StatusRegister) User() uint8 {
	return uint8(uint32(st) & STUser)
}

// SetUser sets the user defined ST bits
func (st StatusRegister) SetUser(sb uint8) {
	st = StatusRegister(((uint32(st) | STUser) - STUser) + uint32(sb))
}

// Registers contains the processor registers
type Registers struct {
	// General Purpose
	R0 uint64
	R1 uint64
	R2 uint64
	R3 uint64

	// Ptr
	PTR0 uint32
	PTR1 uint32

	// Offset
	OFS0 uint16
	OFS1 uint16

	// Index
	IX0 uint16
	IX1 uint16

	// Index Step
	IS0 uint16
	IS1 uint16

	// Counter
	CTR0 uint32
	CTR1 uint32

	// Counter Step
	CS0 uint16
	CS1 uint16

	// Program counter
	PC uint32

	// Status
	st StatusRegister

	// Stack base
	SB uint32

	// Stack pointer
	SP uint16
}

// OfRegisters creates a new Registers where all registers = 0 except:
// SB = DefaultSB
// SP = DefaultSP
func OfRegisters() Registers {
	var reg Registers
	reg.SB = DefaultSB
	reg.SP = DefaultSP

	return reg
}

// ST returns current status register value
func (r Registers) ST() StatusRegister {
	return r.st
}
