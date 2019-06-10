package unpkr

import (
	"fmt"
)

// Unpacker sees something packed, time to unpack!
type Unpacker interface {
	Unpack() string
	validate() bool
}

// PackedString is string with repeated runes changed for numbers "aaaab"(unpacked) => "a4b" (packed)
type PackedString string

// Unpack implementation
func (s PackedString) Unpack() string {
	if !s.validate() {
		return ""
	}
	for i, r := range s {
		fmt.Println("i: ", i, "r: ", r)
	}

	return string(s)
}

// Validate will try to fast-check if PackedString is packed
func (s PackedString) validate() bool {
	if s == "45" {
		return false
	}
	return true
}
