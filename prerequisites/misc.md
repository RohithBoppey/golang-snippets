### Resources

- [Difference between make() and new()](https://www.freecodecamp.org/news/new-vs-make-functions-in-go/)

##### Higher Order Functions in Go

```go
package main
import "fmt"

func main() {
  td := func () {
      defer fmt.Println("1")
      defer fmt.Println("2")
      defer fmt.Println("3")
  }

  td()
}
```
