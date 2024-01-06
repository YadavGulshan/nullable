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
	_ runtime.Validatable        = &Int64{}
	_ runtime.ContextValidatable = &Int64{}
	_ runtime.Validatable        = &Int64Slice{}
	_ runtime.ContextValidatable = &Int64Slice{}
)

// Int64 represents a int64 that may be null
type Int64 struct {
	Valid   bool // Valid is true if value is not null and valid int64
	Value   int64
}

// NewInt64 creates a new Int64 object
func NewInt64(val *int64) Int64 {
    if val == nil {
        return Int64{}
    }
    return Int64{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int64) Get() *int64 {
    if !i.Valid {
        return nil
    }
    return &i.Value
}

// Set sets the value and valid flag to true
func (i *Int64) Set(val int64) {
    i.Valid = true
    i.Value = val
}

func (i *Int64) UnmarshalJSON(data []byte) error {
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
func (i *Int64) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Int64Slice represents a []int64 that may be null
type Int64Slice struct {
	Valid   bool // Valid is true if value is not null
	Value   []int64
}


// NewInt64Slice creates a new Int64Slice object
func NewInt64Slice(val []int64) Int64Slice {
    if val == nil {
        return Int64Slice{}
    }
    return Int64Slice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int64Slice) UnmarshalJSON(data []byte) error {
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
func (i *Int64Slice) Get() *[]int64 {
    if !i.Valid {
        return nil
    }
    return &i.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int64Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int64Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}