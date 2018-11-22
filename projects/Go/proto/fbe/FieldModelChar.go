// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "time"
import "github.com/google/uuid"

// Workaround for Go unused imports issue
var _ = time.Unix(0, 0)
var _ = uuid.Nil

// Fast Binary Encoding rune field model class
type FieldModelChar struct {
    buffer *Buffer  // Field model buffer
    offset int      // Field model buffer offset
}

// Get the field size
func (fm FieldModelChar) FBESize() int { return 1 }
// Get the field extra size
func (fm FieldModelChar) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelChar) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelChar) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelChar) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelChar(buffer *Buffer, offset int) *FieldModelChar {
    return &FieldModelChar{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelChar) Verify() bool { return true }

// Get the value
func (fm FieldModelChar) Get() rune {
    return fm.GetDefault('\000')
}

// Get the value with provided default value
func (fm FieldModelChar) GetDefault(defaults rune) rune {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return defaults
    }

    return ReadChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())
}

// Set the value
func (fm *FieldModelChar) Set(value rune) {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return
    }

    WriteChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
}
