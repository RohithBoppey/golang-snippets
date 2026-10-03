package scanning

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "single line", input: "hello world\n", want: "hello world"},
		{name: "only first line", input: "hello\nworld\n", want: "hello"},
		{name: "no trailing newline", input: "hello", want: "hello"},
		{name: "empty input", input: "", wantErr: io.EOF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLine(strings.NewReader(tt.input))
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ReadLine(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ReadLine(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestScanWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "simple", input: "hello world", want: []string{"hello", "world"}},
		{name: "extra spaces", input: "  a   b\tc  ", want: []string{"a", "b", "c"}},
		{name: "empty", input: "", want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScanWords(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScanWords(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
