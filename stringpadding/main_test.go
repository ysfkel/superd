package stringpadding

import (
	"testing"
)

func TestPadNumbers(t *testing.T) {
	tests := []struct {
		input string
		x     int
		want  string
	}{
		{"James Bond 7", 3, "James Bond 007"},
		{"PI=3.14", 2, "PI=03.14"},
		{"It's 3:13pm", 2, "It's 03:13pm"},
		{"It's 12:13pm", 2, "It's 12:13pm"},
		{"99UR1337", 6, "000099UR001337"},
	}

	for _, tt := range tests {
		got := PadNumbers(tt.input, tt.x)
		if got != tt.want {
			t.Errorf("PadNumbers(%q, %d) = %q; want %q", tt.input, tt.x, got, tt.want)
		}
	}
}
