// SPDX-License-Identifier: Apache-2.0

package register

const (
	// DefaultSB is the default stack base
	DefaultSB uint32 = 0xFFFE0000

	// DefaultSP is the default stack pointer
	DefaultSP uint16 = 0xFFFF
	
	// Bottom of stack for pushing an 8 bit value
	StackBottom8 int32 = 0x0000
	
	// Bottom of stack for pushing a 16 bit value
	StackBottom16 int32 = 0x0001
	
	// Bottom of stack for pushing a 32 bit value
	StackBottom32 int32 = 0x0003
	
	// Bottom of stack for pushing a 64 bit value
	StackBottom64 int32 = 0x0007
)

// StackError represents an error performing a stack operation
type StackError string

func (e StackError) Error() string {
    return string(e)
}

var (
    ErrStackOverflow = StackError("Stack Overflow")
    ErrStackUnderflow = StackError("Staack Underflow") 
)

// Stack represents a 64K block of memory that is the current stack space.
// Stack does not allocate space, it manages an allocated space.
type Stack struct {
    stack []uint8
    ptr int32
}

// NewStack constructs a Stack from allocated space and initial pointer value.
// 
func NewStack(
    stack []uint8,
    ptr uint16,
) *Stack {
    return &Stack{
        stack: stack,
        ptr: int32(ptr),
    }
}

// Push8 pushes an 8 bit value to the stack.
// Throws ErrStackOverflow is the stack does not have at least 8 bits left.
func (s *Stack) Push8(op uint8) {
    if (s.ptr == StackBottom8) {
        panic(ErrStackOverflow)
    }
    
    s.stack[s.ptr] = op
    s.ptr--
}

// Push16 pushes a 16 bit value to the stack.
// Throws ErrStackOverflow is the stack does not have at least 16 bits left.
func (s *Stack) Push16(op uint16) {
    if (s.ptr <= StackBottom16) {
        panic(ErrStackOverflow)
    }
    
    s.stack[s.ptr] = (uint8)(op & 0x00FF)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0xFF00)
    s.ptr--
}

// Push32 pushes a 32 bit value to the stack.
// Throws ErrStackOverflow is the stack does not have at least 32 bits left.
func (s *Stack) Push32(op uint32) {
    if (s.ptr <= StackBottom32) {
        panic(ErrStackOverflow)
    }
    
    s.stack[s.ptr] = (uint8)(op & 0x000000FF)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x0000FF00)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x00FF0000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0xFF000000)
    s.ptr--
}

// Push64 pushes a 64 bit value to the stack.
// Throws ErrStackOverflow is the stack does not have at least 64 bits left.
func (s *Stack) Push64(op uint64) {
    if (s.ptr <= StackBottom64) {
        panic(ErrStackOverflow)
    }
    
    s.stack[s.ptr] = (uint8)(op & 0x00000000000000FF)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x000000000000FF00)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x0000000000FF0000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x00000000FF000000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x000000FF00000000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x0000FF0000000000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0x00FF000000000000)
    s.ptr--
    s.stack[s.ptr] = (uint8)(op & 0xFF00000000000000)
    s.ptr--
}
