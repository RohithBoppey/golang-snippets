package scanning_test

import (
	"fmt"
	"strings"

	"golang-snippets/scanning"
)

func ExampleReadLine() {
	line, _ := scanning.ReadLine(strings.NewReader("hello world\nignored"))
	fmt.Println(line)
	// Output: hello world
}

func ExampleScanWords() {
	words := scanning.ScanWords("go is   fun")
	fmt.Println(words)
	fmt.Println(len(words))
	// Output:
	// [go is fun]
	// 3
}

func ExampleScanInt() {
	a, _ := scanning.ScanInt(strings.NewReader("42"))
	fmt.Println(a)
	// Output: 42
}
