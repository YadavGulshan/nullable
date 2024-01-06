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
	_ runtime.Validatable        = &Bool{}
	_ runtime.ContextValidatable = &Bool{}
	_ runtime.Validatable        = &BoolSlice{}
	_ runtime.ContextValidatable = &BoolSlice{}
)

// Bool represents a bool that may be null
type Bool struct {
	Valid   bool // Valid is true if value is not null and valid bool
	Value   bool
}

// NewBool creates a new Bool object
func NewBool(val *bool) Bool {
    if val == nil {
        return Bool{}
    }
    return Bool{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (b *Bool) Get() *bool {
    if !b.Valid {
        return nil
    }
    return &b.Value
}

// Set sets the value and valid flag to true
func (b *Bool) Set(val bool) {
    b.Valid = true
    b.Value = val
}

func (b *Bool) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        b.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &b.Value); err != nil {
        return err
    }
    b.Valid = true
    return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (b *Bool) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (b *Bool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// BoolSlice represents a []bool that may be null
type BoolSlice struct {
	Valid   bool // Valid is true if value is not null
	Value   []bool
}


// NewBoolSlice creates a new BoolSlice object
func NewBoolSlice(val []bool) BoolSlice {
    if val == nil {
        return BoolSlice{}
    }
    return BoolSlice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (b *BoolSlice) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        b.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &b.Value); err != nil {
        return err
    }
    b.Valid = true
    return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (b *BoolSlice) Get() *[]bool {
    if !b.Valid {
        return nil
    }
    return &b.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (b *BoolSlice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (b *BoolSlice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}