// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable


import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Compile time validation that our types implement the expected interfaces
var (
	_ runtime.Validatable        = &Float32{}
	_ runtime.ContextValidatable = &Float32{}
	_ runtime.Validatable        = &Float32Slice{}
	_ runtime.ContextValidatable = &Float32Slice{}
)

// Float32 represents a float32 that may be null
type Float32 struct {
	Valid   bool // Valid is true if value is not null and valid float32
	Value   float32
}

// NewFloat32 creates a new Float32 object
func NewFloat32(val *float32) Float32 {
    if val == nil {
        return Float32{}
    }
    return Float32{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float32) Get() *float32 {
    if !f.Valid {
        return nil
    }
    return &f.Value
}

// Set sets the value and valid flag to true
func (f *Float32) Set(val float32) {
    f.Valid = true
    f.Value = val
}

func (f *Float32) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) || bytes.Equal(data, empty) || bytes.Equal(data, []byte("")) {
        return nil
    }
    if err := json.Unmarshal(data, &f.Value); err != nil {
        return err
    }
    f.Valid = true
    return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (f *Float32) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float32) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Float32Slice represents a []float32 that may be null
type Float32Slice struct {
	Valid   bool // Valid is true if value is not null
	Value   []float32
}


// NewFloat32Slice creates a new Float32Slice object
func NewFloat32Slice(val []float32) Float32Slice {
    if val == nil {
        return Float32Slice{}
    }
    return Float32Slice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float32Slice) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        f.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &f.Value); err != nil {
        return err
    }
    f.Valid = true
    return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float32Slice) Get() *[]float32 {
    if !f.Valid {
        return nil
    }
    return &f.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (f *Float32Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float32Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}