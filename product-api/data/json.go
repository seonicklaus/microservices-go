package data

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from JSON strinng
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// ToJSON serializes the given interface into a string based on JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
