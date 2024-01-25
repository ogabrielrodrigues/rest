package rest

import (
	"encoding/json"
	"io"
)

// Bind the src JSON request to dst struct.
// Returns a rest.Err if an error occurs in the binding process
func Bind(src io.ReadCloser, dst interface{}) *Err {
	if err := json.NewDecoder(src).Decode(&dst); err != nil {
		return NewBadRequestErr("error decoding request body", nil)
	}

	if err := Validate.Struct(dst); err != nil {
		return ValidateStructErr(err)
	}

	return nil
}
