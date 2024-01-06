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
	_ runtime.Validatable        = &Float64{}
	_ runtime.ContextValidatable = &Float64{}
	_ runtime.Validatable        = &Float64Slice{}
	_ runtime.ContextValidatable = &Float64Slice{}
)

// Float64 represents a float64 that may be null
type Float64 struct {
	Valid   bool // Valid is true if value is not null and valid float64
	Value   float64
}

// NewFloat64 creates a new Float64 object
func NewFloat64(val *float64) Float64 {
    if val == nil {
        return Float64{}
    }
    return Float64{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float64) Get() *float64 {
    if !f.Valid {
        return nil
    }
    return &f.Value
}

// Set sets the value and valid flag to true
func (f *Float64) Set(val float64) {
    f.Valid = true
    f.Value = val
}

func (f *Float64) UnmarshalJSON(data []byte) error {
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
func (f *Float64) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Float64Slice represents a []float64 that may be null
type Float64Slice struct {
	Valid   bool // Valid is true if value is not null
	Value   []float64
}


// NewFloat64Slice creates a new Float64Slice object
func NewFloat64Slice(val []float64) Float64Slice {
    if val == nil {
        return Float64Slice{}
    }
    return Float64Slice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float64Slice) UnmarshalJSON(data []byte) error {
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
func (f *Float64Slice) Get() *[]float64 {
    if !f.Valid {
        return nil
    }
    return &f.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (f *Float64Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float64Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}