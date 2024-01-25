package rest

import (
	"encoding/json"
	"io"
)

func Bind(src io.ReadCloser, dst interface{}) *Err {
	if err := json.NewDecoder(src).Decode(&dst); err != nil {
		return NewBadRequestErr("error decoding request body", nil)
	}

	if err := Validate.Struct(dst); err != nil {
		return ValidateErr(err)
	}

	return nil
}
