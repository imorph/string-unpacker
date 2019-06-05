package unpkr

// Unpacker sees something packed, time to unpack!
type Unpacker interface {
	Unpack() string
}

// PackedString is string with repeated runes changed for numbers "aaaab" => "a4b"
type PackedString string

// Unpack implementation
func (s PackedString) Unpack() string {
	return string(s)
}
