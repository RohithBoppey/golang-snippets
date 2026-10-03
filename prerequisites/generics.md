- simple function to understand the generic function syntax: 

```go
package main
import "fmt"

func Identity [T any] (v T) (out T){
    return v; 
}

func main() {
    fmt.Println(Identity(5));
    fmt.Println(Identity("str"));
}
```

- take any template and swap the 2 variables inside it
```go
package main
import "fmt"

func Swap[T any] (a T, b T) (T, T) {
    return b, a
}

func main() {
	x, y := Swap(1, 2)
	fmt.Println(x, y) // 2 1

	s1, s2 := Swap("a", "b")
	fmt.Println(s1, s2) // b a
}
```

if comparision is there, then use comparable instead of any
- a contains function using 
```go
package main
import "fmt"

// contains function

func CustomContains [T comparable] (arr []T, val T) (bool) {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    
    return false
}

func main() {
    // int arr contains
    arr := []int{1,2,3,4,5} 
    fmt.Println(CustomContains(arr, 5));
    fmt.Println(CustomContains(arr, 6));
    
    // string arr contains
    arrs := []string{"abc", "def"}
    fmt.Println(CustomContains(arrs, "abc"));
    fmt.Println(CustomContains(arrs, "def"));
}
```

- take a function as a parameter as an interface on it, do it 
```go
package main
import "fmt"

// idea to write a template-ish function which takes in an array of values, takes a function to run on it, and returns a new array of a new kind of functions  

func Apply [T any, R any] (arr []T, fn func(T) R) []R {
    // make a new array of the same length
    res := make([]R, len(arr))
    for i, v := range(arr) {
        res[i] = fn(v)
    }
    return res
}

func sum_concat [X int | float64 | string] (v X) X { return v + v }
 
func main () { 
    arr := []int{1,2,3,4,5} 
    fmt.Println(arr)
    // now transform
    arr2 := Apply(arr, sum_concat)
    fmt.Println(arr2)    
    
    // strings
    sarr := []string{"nice", "dad"} 
    fmt.Println(sarr)
    // now transform
    sarr2 := Apply(sarr, sum_concat)
    fmt.Println(sarr2)    
}
```

- generic max function (but the type needs to be explicitly mentioned)
```go
package main
import "fmt"

// custom allowed interface
type Ordered interface {
    ~int | ~int64 | ~float64 | ~string
}

// contains function
func Maxx[T Ordered] (a, b T) T {
    if(a > b) {
        return a; 
    }else {
        return b; 
    }
}

func main() {
	fmt.Println(Maxx(10, 20))    // 20
	fmt.Println(Maxx(3.5, 2.1))  // 3.5
	fmt.Println(Maxx("string", "string1"))  
}
```


- generic filter function to take it differnt kind of functions, and filter out the same type of elements on what to take

```go
package main
import "fmt"

// generic filter function 
func GenericFilter [T any] (arr []T, fn func(T) bool) []T {
    // we do not know the length
    var res []T
    // filtering logic 
    for _, val := range arr {
        if fn(val) { 
            // it passed
            res = append(res, val)
        }
    }
    return res; 
}

// 1. 
func evenCheck(x int) bool {
    return x % 2 == 0; 
}

// 2. 
func vowelCheck(s string) bool{
    for _, ch := range(s){
        if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
            return false; 
        }
    }
    return true; 
}

func main () { 
    // 1. even filter on int
    arr := []int{1,2,3,4,5}
    res := GenericFilter(arr, evenCheck)
    fmt.Println(res)
    
    // 2. vowel check for string
    arr2:= []string{"ab", "cd", "our", "us", "rsh"}
    res2 := GenericFilter(arr2, vowelCheck)
    fmt.Println(res2)
}
```