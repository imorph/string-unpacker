package unpkr

import "testing"

func TestUnpack001(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "aaaab"
	want = "a4b"
	got = ps.Unpack()
	
	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack002(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = ps.Unpack()
	
	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack003(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "abcd"
	want = "abcd"
	got = ps.Unpack()
	
	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack004(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "45"
	want = ""
	got = ps.Unpack()
	
	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
