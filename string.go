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
	_ runtime.Validatable        = &String{}
	_ runtime.ContextValidatable = &String{}
	_ runtime.Validatable        = &StringSlice{}
	_ runtime.ContextValidatable = &StringSlice{}
)

// String represents a string that may be null
type String struct {
	Valid   bool // Valid is true if value is not null and valid string
	Value   string
}

// NewString creates a new String object
func NewString(val *string) String {
    if val == nil {
        return String{}
    }
    return String{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (s *String) Get() *string {
    if !s.Valid {
        return nil
    }
    return &s.Value
}

// Set sets the value and valid flag to true
func (s *String) Set(val string) {
    s.Valid = true
    s.Value = val
}

func (s *String) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) || bytes.Equal(data, empty) || bytes.Equal(data, []byte("")) {
        return nil
    }
    if err := json.Unmarshal(data, &s.Value); err != nil {
        return err
    }
    s.Valid = true
    return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (s *String) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (s *String) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// StringSlice represents a []string that may be null
type StringSlice struct {
	Valid   bool // Valid is true if value is not null
	Value   []string
}


// NewStringSlice creates a new StringSlice object
func NewStringSlice(val []string) StringSlice {
    if val == nil {
        return StringSlice{}
    }
    return StringSlice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (s *StringSlice) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        s.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &s.Value); err != nil {
        return err
    }
    s.Valid = true
    return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (s *StringSlice) Get() *[]string {
    if !s.Valid {
        return nil
    }
    return &s.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (s *StringSlice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (s *StringSlice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}