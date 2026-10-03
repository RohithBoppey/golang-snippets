package scanning

import (
	"bufio"
	"io"
	"strings"
)

// ReadLine reads a single line from r, without the trailing newline.
func ReadLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return scanner.Text(), nil
}

// ScanWords splits a line into its whitespace-separated words.
func ScanWords(line string) []string {
	return strings.Fields(line)
}
