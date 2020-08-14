package dto

import (
	"encoding/json"
	"io"
)

// Sample represents an incoming request
type Sample struct {
	Content string `json:"content"`
}

// FromJSON generates a Sample from given buffer
func (s *Sample) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	err := d.Decode(s)
	if err != nil {
		return err
	}

	return nil
}

// ToJSON writes JSON dump data
func (s *Sample) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	err := e.Encode(s)
	if err != nil {
		return err
	}

	return nil
}
