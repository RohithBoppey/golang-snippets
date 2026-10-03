package scanning

import (
	"fmt"
	"io"
)

// ScanInt reads a single integer from r.
func ScanInt(r io.Reader) (int, error) {
	var a int
	if _, err := fmt.Fscan(r, &a); err != nil {
		return 0, err
	}
	return a, nil
}
