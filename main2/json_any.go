package main

import (
	"encoding/json"
	"errors"
)

// Any is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type Any []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m Any) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return json.Marshal(m)
}

// UnmarshalJSON sets *m to a copy of data.
func (m *Any) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func MarshalJSONAny(value interface{}) (*Any, error) {
	bs, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	a := Any(bs)
	return &a, nil
}

func UnmarshalJSONAny(any *Any, value interface{}) error {
	return json.Unmarshal(*any, value)
}

var _ json.Marshaler = (*Any)(nil)
var _ json.Unmarshaler = (*Any)(nil)
