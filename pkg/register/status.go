package register

import (
	"github.com/bantling/gofuncs"
)

const (
	// STCarrySet is ST filter for setting carry flag
	STCarrySet uint32 = 0x80000000

	// STCarryClear is ST filter for clearing carry flag
	STCarryClear uint32 = 0xFFFFFFFF - STCarrySet

	// STOverflowSet is ST filter for setting overflow flag
	STOverflowSet uint32 = 0x40000000

	// STOverflowClear is ST filter for clearing overflow flag
	STOverflowClear uint32 = 0xFFFFFFFF - STOverflowSet

	// STZeroSet is ST filter for setting zero flag
	STZeroSet uint32 = 0x20000000

	// STZeroClear is ST filter for clearing zero flag
	STZeroClear uint32 = 0xFFFFFFFF - STZeroSet

	// STNegativeSet is ST filter for setting negative flag
	STNegativeSet uint32 = 0x10000000

	// STNegativeClear is ST filter for clearing negative flag
	STNegativeClear uint32 = 0xFFFFFFFF - STNegativeSet

	// STAddressRead is ST filter for reading address mode
	STAddressRead uint32 = 0x0E000000

	// STAddressSet is ST filter for setting address mode
	STAddressSet uint32 = 0xFFFFFFFF - STAddressRead

	// STAddressShift is ST shift for address
	STAddressShift int = 25

	// STInterruptDisableSet is ST filter for setting interrupt disable
	STInterruptDisableSet uint32 = 0x01000000

	// STInterruptDisableClear is ST filter for clearing interrupt disable
	STInterruptDisableClear uint32 = 0xFFFFFFFF - STInterruptDisableSet

	// STRegisterRead is ST filter for reading register mode
	STRegisterRead uint32 = 0x00C00000

	// STRegisterSet is ST filter for setting register mode
	STRegisterSet uint32 = 0xFFFFFFFF - STRegisterRead

	// STRegisterShift is the ST shift for register
	STRegisterShift int = 22

	// STPointerRegisterSet is ST filter for pointer register set
	STPointerRegisterSet uint32 = 0x00200000

	// STPointerRegisterClear is ST filter for pointer register clear
	STPointerRegisterClear uint32 = 0xFFFFFFFF - STPointerRegisterSet

	// STCounterRegisterSet is ST filter for counter register set
	STCounterRegisterSet uint32 = 0x00100000

	// STCounterRegisterClear is ST filter for counter register clear
	STCounterRegisterClear uint32 = 0xFFFFFFFF - STCounterRegisterSet

	// STOperandRead is ST filter for reading operand size
	STOperandRead uint32 = 0x000C0000

	// STOperandSet is ST filter for setting operand size
	STOperandSet uint32 = 0xFFFFFFFF - STOperandRead

	// STOperandShift is the ST shift for operand
	STOperandShift int = 18

	// STMathRead is ST filter for reading math mode
	STMathRead uint32 = 0x00030000

	// STMathSet is ST filter for setting math mode
	STMathSet uint32 = 0xFFFFFFFF - STMathRead

	// STMathShift is the ST shift for math mode
	STMathShift int = 16

	// STSystemRead is ST filter for reading system bits
	STSystemRead uint32 = 0x0000FF00

	// STSystemSet is ST filter for setting system bits
	STSystemSet uint32 = 0xFFFFFFFF - STSystemRead

	// STSystemShift is the ST shift for system bits
	STSystemShift int = 8

	// STUserRead is ST filter for reading user bits
	STUserRead uint32 = 0x000000FF

	// STUserSet is ST filter for setting user bits
	STUserSet uint32 = 0xFFFFFFFF - STUserRead
)

// Address modes
const (
	Ptr         uint8 = iota // 000 1 = 1
	PtrOfs                   // 001 1 = 3
	PtrIx                    // 010 1 = 5
	PtrIxOfs                 // 011 1 = 7
	PtrPtr                   // 100 1 = 9
	PtrPtrOfs                // 101 1 = B
	PtrPtrIx                 // 110 1 = D
	PtrPtrIxOfs              // 111 1 = F
)

// STAddressModeErr ifaAddress mode is invalid
const STAddressModeErr = "Address mode must be <= 7"

// Register selection
const (
	RegisterR0 uint8 = iota // 00 11 = 3
	RegisterR1              // 01 11 = 7
	RegisterR2              // 10 11 = B
	RegisterR3              // 11 11 = F
)

// STRegisterErr if register is invalid
const STRegisterErr = "Register must be <= 3"

// Operand sizes
const (
	Operand8  uint8 = iota // 00 11 = 3
	Operand16              // 01 11 = 7
	Operand32              // 10 11 = B
	Operand64              // 11 11 = F
)

// STOperandModeErr if operand mode is invalid
const STOperandModeErr = "Operand mode must be <= 3"

// Math modes
const (
	MathInteger    uint8 = iota // 11 00 = C
	MathFractional              // 11 01 = D
	MathFixed                   // 11 10 = E
	MathFloat                   // 11 11 = F
)

// STMathModeErr if math mode is invalid
const STMathModeErr = "Math mode must be <= 3"

// StatusRegister defines the status register.
// CVZNAAAI RRPTOOMM SSSSSSSS UUUUUUUU.
// Carry, oVerflow, Zero, Negative, Address mode, Interrupt Disable,
// Register, Pointer register set, counTer register set, Operand size, Math mode.
// 8 bits reserved for system use, and 8 bits reserved for users.
type StatusRegister uint32

// Carry returns true if the carry flag is set
func (st StatusRegister) Carry() bool {
	return (uint32(st) & STCarrySet) == STCarrySet
}

// SetCarry sets the carry flag
func (st *StatusRegister) SetCarry() {
	*st |= StatusRegister(STCarrySet)
}

// ClearCarry clears the carry flag
func (st *StatusRegister) ClearCarry() {
	*st &= StatusRegister(STCarryClear)
}

// Overflow returns true if the overflow flag is set
func (st StatusRegister) Overflow() bool {
	return (uint32(st) & STOverflowSet) == STOverflowSet
}

// SetOverflow sets the overflow flag
func (st *StatusRegister) SetOverflow() {
	*st |= StatusRegister(STOverflowSet)
}

// ClearOverflow clears the overflow flag
func (st *StatusRegister) ClearOverflow() {
	*st &= StatusRegister(STOverflowClear)
}

// Zero returns true if the overflow flag is set
func (st StatusRegister) Zero() bool {
	return (uint32(st) & STZeroSet) == STZeroSet
}

// SetZero sets the overflow flag
func (st *StatusRegister) SetZero() {
	*st |= StatusRegister(STZeroSet)
}

// ClearZero clears the overflow flag
func (st *StatusRegister) ClearZero() {
	*st &= StatusRegister(STZeroClear)
}

// Negative returns true if the negative flag is set
func (st StatusRegister) Negative() bool {
	return (uint32(st) & STNegativeSet) == STNegativeSet
}

// SetNegative sets the negative flag
func (st *StatusRegister) SetNegative() {
	*st |= StatusRegister(STNegativeSet)
}

// ClearNegative clears the negative flag
func (st *StatusRegister) ClearNegative() {
	*st &= StatusRegister(STNegativeClear)
}

// AddressMode returns the address mode
func (st StatusRegister) AddressMode() uint8 {
	return uint8((uint32(st) & STAddressRead) >> STAddressShift)
}

// SelectAddressMode selects an address mode
func (st *StatusRegister) SelectAddressMode(am uint8) {
	gofuncs.PanicBM(am <= PtrPtrIxOfs, STAddressModeErr)

	*st = StatusRegister((uint32(*st) & STAddressSet) + (uint32(am) << STAddressShift))
}

// InterruptDisable returns true if the interrupt disable flag is set
func (st StatusRegister) InterruptDisable() bool {
	return (uint32(st) & STInterruptDisableSet) == STInterruptDisableSet
}

// SetInterruptDisable sets the interrupt disable flag
func (st *StatusRegister) SetInterruptDisable() {
	*st |= StatusRegister(STInterruptDisableSet)
}

// ClearInterruptDisable clears the interrupt disable flag
func (st *StatusRegister) ClearInterruptDisable() {
	*st = StatusRegister(STInterruptDisableClear)
}

// Register returns the selected register
func (st StatusRegister) Register() uint8 {
	return uint8((uint32(st) & STRegisterRead) >> STRegisterShift)
}

// SelectRegister selects a register
func (st *StatusRegister) SelectRegister(rm uint8) {
	gofuncs.PanicBM(rm <= Operand64, STRegisterErr)

	*st = StatusRegister((uint32(*st) & STRegisterSet) + (uint32(rm) << STRegisterShift))
}

// PointerRegisterSet0 returns true if pointer register set 0 is selected
func (st StatusRegister) PointerRegisterSet0() bool {
	return (uint32(st) & STPointerRegisterSet) == 0
}

// SelectPointerRegisterSet0 selects pointer register set 0
func (st *StatusRegister) SelectPointerRegisterSet0() {
	*st &= StatusRegister(STPointerRegisterClear)
}

// SelectPointerRegisterSet1 selects pointer register set 1
func (st *StatusRegister) SelectPointerRegisterSet1() {
	*st |= StatusRegister(STPointerRegisterSet)
}

// CounterRegisterSet0 returns true if counter register set 0 is selected
func (st StatusRegister) CounterRegisterSet0() bool {
	return (uint32(st) & STCounterRegisterSet) == 0
}

// SelectCounterRegisterSet0 selects counter register set 0
func (st *StatusRegister) SelectCounterRegisterSet0() {
	*st &= StatusRegister(STCounterRegisterClear)
}

// SelectCounterRegisterSet1 selects counter register set 1
func (st *StatusRegister) SelectCounterRegisterSet1() {
	*st |= StatusRegister(STCounterRegisterSet)
}

// OperandSize returns the operand size
func (st StatusRegister) OperandSize() uint8 {
	return uint8((uint32(st) & STOperandRead) >> STOperandShift)
}

// SelectOperandSize selects the operand size
func (st *StatusRegister) SelectOperandSize(om uint8) {
	gofuncs.PanicBM(om <= Operand64, STOperandModeErr)

	*st = StatusRegister((uint32(*st) & STOperandSet) + (uint32(om) << STOperandShift))
}

// MathMode returns the math mode
func (st StatusRegister) MathMode() uint8 {
	return uint8((uint32(st) & STMathRead) >> STMathShift)
}

// SelectMathMode sets the math mode
func (st *StatusRegister) SelectMathMode(mm uint8) {
	gofuncs.PanicBM(mm <= MathFloat, STMathModeErr)

	*st = StatusRegister((uint32(*st) & STMathSet) + (uint32(mm) << STMathShift))
}

// System returns the system defined ST bits
func (st StatusRegister) System() uint8 {
	return uint8((uint32(st) & STSystemRead) >> STSystemShift)
}

// SetSystem sets the system defined ST bits
func (st *StatusRegister) SetSystem(sb uint8) {
	*st = StatusRegister((uint32(*st) & STSystemSet) + (uint32(sb) << STSystemShift))
}

// User returns the user defined ST bits
func (st StatusRegister) User() uint8 {
	return uint8(uint32(st) & STUserRead)
}

// SetUser sets the user defined ST bits
func (st *StatusRegister) SetUser(sb uint8) {
	*st = StatusRegister((uint32(*st) & STUserSet) + uint32(sb))
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
