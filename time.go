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
	_ runtime.Validatable        = &Time{}
	_ runtime.ContextValidatable = &Time{}
	_ runtime.Validatable        = &TimeSlice{}
	_ runtime.ContextValidatable = &TimeSlice{}
)

// Time represents a time that may be null
type Time struct {
	Valid   bool // Valid is true if value is not null and valid time
	Value   time
}

// NewTime creates a new Time object
func NewTime(val *time) Time {
    if val == nil {
        return Time{}
    }
    return Time{
        Valid: true,
        Value: *val,
    }
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (t *Time) Get() *time {
    if !t.Valid {
        return nil
    }
    return &t.Value
}

// Set sets the value and valid flag to true
func (t *Time) Set(val time) {
    t.Valid = true
    t.Value = val
}

func (t *Time) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) || bytes.Equal(data, empty) || bytes.Equal(data, []byte("")) {
        return nil
    }
    if err := json.Unmarshal(data, &t.Value); err != nil {
        return err
    }
    t.Valid = true
    return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (t *Time) Validate(formats strfmt.Registry) error {
    return nil 
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (t *Time) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// TimeSlice represents a []time that may be null
type TimeSlice struct {
	Valid   bool // Valid is true if value is not null
	Value   []time
}


// NewTimeSlice creates a new TimeSlice object
func NewTimeSlice(val []time) TimeSlice {
    if val == nil {
        return TimeSlice{}
    }
    return TimeSlice{
        Valid: true,
        Value: val,
    }
}

// UnmarshalJSON implements json.Marshaler interface.
func (t *TimeSlice) UnmarshalJSON(data []byte) error {
    if bytes.Equal(data, null) {
        t.Valid = false
        return nil
    }
    if err := json.Unmarshal(data, &t.Value); err != nil {
        return err
    }
    t.Valid = true
    return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (t *TimeSlice) Get() *[]time {
    if !t.Valid {
        return nil
    }
    return &t.Value
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (t *TimeSlice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (t *TimeSlice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}