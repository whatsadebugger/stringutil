package stringutil

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range testCases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) gave -> %q, want %q", c.in, got, c.want)
		}
	}
}
