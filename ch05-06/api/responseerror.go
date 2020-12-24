package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ResponseError is the error for the JSON response.
type ResponseError struct {
	Err error
}

// MarshalJSON implements JSON Marshaler for ResponseError
func (re ResponseError) MarshalJSON() ([]byte, error) {
	if re.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, re.Err)), nil
}

// UnmarshalJSON implements JSON Unmarshaler for ResponseError
func (re *ResponseError) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	if v == nil {
		re.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		re.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError.UnmarshalJSON: unmarshal failed")
	}
}
