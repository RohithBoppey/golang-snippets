[Reference](https://www.w3schools.com/go/go_variables.php)

Types in Golang: 
- bool
- number (int, int32, int64, float32, float64)
- string
- array
- slice
- struct
- map
- pointer

When declaring a variable, we can have const variables as well (inside and outside the function)        
```go
package main
import "fmt"

const PI = 3.14

/*
Rules about constants: 
- when declaring a const only, give a value
- can be inside and outside functions
- need not give a type explicitly or infer using ":="
*/

func main() {
    const LMAO = 2
    const namedConstant string = "named constant"
    fmt.Println(PI, LMAO, namedConstant)
}
```

Default Variables values:
```go
package main
import "fmt"

var a int
var b float32
var str string
var bo bool

func main() {
    fmt.Println(a, b, str, bo)
}
```

Different ways to infer the type: 
```go
package main
import "fmt"

func main() {
    var a string = "this is string" // Specifying the type
    var b = "This is nice" // this is nice type
    c := 2
    fmt.Println(a, b, c)
}
```

```go
package main
import "fmt"

func main() {
    var a string = "this is string"
    var b = "This is nice"
    c := 2
    b = "changing value of b"  // at this point, you can only have b value as string, and not float
    fmt.Println(a, b, c)
}
```

- "var" can be used inside and outside functions, but ":=" can only be used inside functions!

You can declare multiple type of variables within a block
```go
package main
import "fmt"

func main() {
    var (
    defint int = 10
    dumint int 
    sat string = "string"
    )
    
    fmt.Println(dumint, defint, sat)
}
```

---

Arrays or slices (arrays are fixed, slices are dynamic) are quite different 

```go
package main
import "fmt"

func main() {
    // give the size and array type
    var (
        arr = [5] string{"val1", "val2"}
    )
    
    // infer the array type, but have the fixed length
    arr2 := [5]string{"val1", "val2"}
    
    // infer the size as well -- but note that this is an array, not a slice (hence, cannot use append to it)
    var noSizeArr = [...]string{"hello", "hi"}

    for _, val := range(arr) {
        fmt.Printf("%v | %T\n", val, val)
    }

    fmt.Printf("%T\n\n", noSizeArr)
    
    fmt.Println(arr2)
}
```
- once you declare an array with fixed size, it is already initialized, so the only way you can add is with the indexing only 
- dynamic appending, etc. is only possible in slice (dynamic array)

Slice - dynamic array with 0 len and 0 capacity (or a specified capacity)
```go
package main
import "fmt"

func main() {
    arr := [2]int{1,2} // array
    darr := make([]string, 2, 5) // slice
    fmt.Printf("%T ; %T \n", arr, darr)
}
```

- Slice example
```go
package main
import "fmt"

func main() {
    // declaring a slice
    var (
        slice = []int{1,2,3}
    )
    // 2 functions - len and cap
    fmt.Println(slice, len(slice), cap(slice))

    // append using the append function
    // Go uses a capacity growth algorithm (*2 always if the limit is reached)
    slice = append(slice, 4, 5) // -- at this point, cap is 6
    fmt.Println(slice, len(slice), cap(slice))   

    slice = append(slice, 6, 7, 8)
    fmt.Println(slice, len(slice), cap(slice)) // ---- now this is 8,12
}
```

- Maps understanding
```go
package main
import "fmt"

func main() {
    // maps
    var ( 
        // can provide type like this
        mp map[string]float64 = map[string]float64{"bargav": 0.5}
        // or else without type as well
        mp2 = map[string]float64{"chandra": 0.51, "rohith": 0.1, "sravanth": 0.9}
    )

    // change from declared one
    mp2["rohith"] = 1
    mp["rohith"] = 1

    fmt.Println(mp, mp2)

    // ways to create an empty map
    var (
        emp1 map[string]int    // this gives true empty (but know that we cannot write into it)
        emp2 = make(map[string]int)    // this gives false - unempty
    )

    fmt.Println(emp1 == nil, emp2 == nil) // true, false

    // emp1["string"] = 1 // returns a runtime panic
    emp2["string"] = 12 // WORKS!

    fmt.Println(emp1, emp2)

    // delete from the map
    delete(emp2, "string")
    fmt.Println(emp1, emp2)

    // iterate a map
    for key, value := range(mp2){
        fmt.Println(key, value)
    }

    // check if a key exists or not, but with the same value
    mp["rohith"] = 0
    val, ok := mp["new member"]    // this false denotes that that key is not present in this
    fmt.Println(val, ok)
    val2, ok2 := mp["rohith"]
    fmt.Println(val2, ok2)
}
```

- Maps practical example
```go
package main 
import "fmt"

func main() { 
    // a person contains things he likes
    type Person struct {
        name string
        likes []string
    }
    
    // overall likes map contains string to list of persons
    var likesMap = make(map[string][]*Person)
    
    fmt.Println(likesMap)

    // create a list of persons now
    var PersonsList = []*Person{
    &Person{"rohith", []string{
        "shawarma",
        "burger",
        "fried rice",
        "biryani",
        "ice cream",
        "pasta",
        "tacos",
        "noodles",
    }},
    &Person{"sravanth", []string{
        "shawarma",
        "pizza",
        "noodles",
        "biryani",
        "sushi",
        "sandwich",
        "pani puri",
        "brownie",
    }},
    &Person{"bargav", []string{
        "burger",
        "tacos",
        "pasta",
        "idli",
        "dosa",
        "ramen",
        "bhel puri",
        "fried chicken",
    }},
    &Person{"chandra", []string{
        "pizza",
        "ice cream",
        "fried rice",
        "noodles",
        "vada pav",
        "manchuria",
        "shawarma",
        "prawn curry",
    }},
}

    fmt.Println("Persons list created!", PersonsList)

    // now from the list of persons, create the likes map
    for _, val := range(PersonsList) {
        // iterate that persons likes
        for _, like := range(val.likes) {
            // add the person to this list
            likesMap[like] = append(likesMap[like], val)
        }
    }

    fmt.Println("Likes map created!", len(likesMap))

    // print a like and their corresponding names of the people
    for like, persons := range(likesMap) {
        fmt.Println(like, ":-")
        for _, person := range(persons) {
            fmt.Println("   -", person.name)
        }
        fmt.Println()
    }
}
```

- Struct example using pointers
```go
package main 

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	// struct definition
	person := Person{"rohith", 22}
	fmt.Printf("%T\n", person)
	// pointer to the struct
	p := &person
	// can be referenced in both ways, go supports nicely
	fmt.Printf("%v | %v\n", p.name, (*p).name)
}
```

Nil Slice example:
```go
package main

import "fmt"

func main() {
	// creating a nil slice using make, and without
	var slice1 []int   // actual nil slice
	var slice2 = make([]int, 0)  // empty slice ready for filling

	fmt.Println(slice1 == nil, slice2 == nil)
}
```
#### Real-world differences

###### 1️⃣ **Memory allocation**

* Nil slice → zero allocations until first append.
* Make slice → allocates memory immediately.

###### 2️⃣ **JSON behavior**
* Nil slice → encodes to null.
* Make slice → encodes to [].
Hence important while developing APIs

###### 4️⃣ **Performance when you know the size**
* If we know the max size we are gonna reach before hand, it is better because it prevents unnecessary copying 
* A nil slice grows like: `0 → 1 → 2 → 4 → 8 → 16 → ...`


