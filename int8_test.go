// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"encoding/json"
	"testing"
)


func TestInt8_UnmarshalJSON(t *testing.T) {
    tests := []struct {
        name string
        buf *bytes.Buffer
        expect Int8
        expectErr error
    }{
        {
            name: "null value",
            buf: bytes.NewBufferString(`{"value": null}`),
            expect: Int8{ Valid: false },
            expectErr: nil,
        },
        {
            name: "empty",
            buf: bytes.NewBufferString(`{}`),
            expect: Int8{},
            expectErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            str := struct {
                Value Int8 `json:"value"`
            }{}

            if err := json.Unmarshal(tt.buf.Bytes(), &str); !typeMatch(tt.expectErr, err) {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
            if got.Valid != tt.expect.Valid {
                t.Errorf("expected Valid to be %#v, got %#v", tt.expect.Valid, got.Valid)
            }
        })
    }
}
