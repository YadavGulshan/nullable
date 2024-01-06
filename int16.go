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
	_ runtime.Validatable        = &Int16{}
	_ runtime.ContextValidatable = &Int16{}
	_ runtime.Validatable        = &Int16Slice{}
	_ runtime.ContextValidatable = &Int16Slice{}
)

// Int16 represents a int16 that may be null
type Int16 struct {
	Valid   bool // Valid is true if value is not null and valid int16
	Value   int16
}

// NewInt16 creates a new Int16 object
func NewInt16(val *int16) Int16 {
    if val == nil {
        return Int16{}
    }
    return Int16{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int16) Get() *int16 {
    if !i.Valid {
        return nil
    }
    return &i.Value
}

// Set sets the value and valid flag to true
func (i *Int16) Set(val int16) {
    i.Valid = true
    i.Value = val
}

func (i *Int16) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        i.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &i.Value); err != nil {
        return err
    }
    i.Valid = true
    return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int16) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int16) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Int16Slice represents a []int16 that may be null
type Int16Slice struct {
	Valid   bool // Valid is true if value is not null
	Value   []int16
}


// NewInt16Slice creates a new Int16Slice object
func NewInt16Slice(val []int16) Int16Slice {
    if val == nil {
        return Int16Slice{}
    }
    return Int16Slice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int16Slice) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        i.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &i.Value); err != nil {
        return err
    }
    i.Valid = true
    return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int16Slice) Get() *[]int16 {
    if !i.Valid {
        return nil
    }
    return &i.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int16Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int16Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}