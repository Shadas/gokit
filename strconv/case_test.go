package strconv

import "testing"

type caseTest struct {
	in  byte
	out byte
}

func TestByteLower(t *testing.T) {
	cases := []caseTest{
		{in: 'a', out: 'a'},
		{in: 'b', out: 'b'},
		{in: 'c', out: 'c'},
		{in: 'A', out: 'a'},
		{in: 'B', out: 'b'},
		{in: 'C', out: 'c'},
		{in: '#', out: '#'},
	}
	for _, c := range cases {
		if out := ByteLower(c.in); out != c.out {
			t.Errorf("TestByteLower failed with in=%c, want_out=%c, get_out=%c", c.in, c.out, out)
		}
	}
}

func TestByteUpper(t *testing.T) {
	cases := []caseTest{
		{in: 'a', out: 'A'},
		{in: 'b', out: 'B'},
		{in: 'c', out: 'C'},
		{in: 'A', out: 'A'},
		{in: 'B', out: 'B'},
		{in: 'C', out: 'C'},
		{in: '#', out: '#'},
	}
	for _, c := range cases {
		if out := ByteUpper(c.in); out != c.out {
			t.Errorf("TestByteLower failed with in=%c, want_out=%c, get_out=%c", c.in, c.out, out)
		}
	}
}
