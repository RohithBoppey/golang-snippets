package scanning

import (
	"strings"
	"testing"
)

func TestScanInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "positive", input: "42", want: 42},
		{name: "negative", input: "-7\n", want: -7},
		{name: "leading spaces", input: "   5 6", want: 5},
		{name: "not a number", input: "abc", wantErr: true},
		{name: "empty", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScanInt(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("ScanInt(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ScanInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
