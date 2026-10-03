Basic [looping](https://www.freecodecamp.org/news/iteration-in-golang/) can be done like this: 

- basic for loop
```go
package main
import "fmt"

func main() {
  for i := 1; i< 10; i++ {.  // note that there is no brackets around the for loop interaction
      fmt.Println(i)
  }
}
```

- basic looping of integers
```go
package main
import "fmt"

func main() {
  numbers := []int{1,2,3,4,5} // creating an array on the fly 
  for i := 0; i< len(numbers); i++ {
      fmt.Println(numbers[i])
  }
}
```

- using the range to traverse an array
```go
package main
import "fmt"

func main() {
  numbers := []int{1,2,3,4,5} 
  
//   using range to extract index, value
  for index, value := range(numbers) {
      fmt.Printf("%v %v\n", index, value)
  }
}
```

- using "range" to traverse a string ; the same can be done for the 
```go
package main
import "fmt"

func main() {
    name := "Rohith"
  
//   using range to extract index, value
  for index, value := range(name) {
    //   fmt.Printf("%v %v\n", index, value) // value is just bytes, need to convert it into string again
      fmt.Printf("%v %v\n", index, string(value))
  }
}
```

- iterating maps 
```go
package main
import "fmt"

func main() {
    dummy_map := map[string]int{
        "ramesh": 1,
        "suneetha": 2,
        "jaswanth": 3, 
        "rohith": 4,
    }
    
    for key, value := range dummy_map{
        fmt.Println(key, value)
    }
    
}
```


