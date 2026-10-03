Basic pointer manipulation code: 
```go
package main 

import "fmt"

func main() {
	// normal int
	a := 2
	fmt.Printf("%T %v\n", a, &a)
	// pointers to that variable
	var p *int = &a
	fmt.Printf("%T %v\n", p, &p)  // here it is a double pointer
	// change value using pointer
	*p = 21
	fmt.Println(p, *p) // while printing as well, use the pointer reference only
}
```