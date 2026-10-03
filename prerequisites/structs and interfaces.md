- In go, we do not have proper classes and functions, here we instead have structs! 
- so all the functions we want, we need to embed into the struct that we declared

- basic creation of struct functions (or class functions)
```go
package main

import (
	"fmt"
	 "math"
	)

// struct can be declared outside of the function scope as well
type CustomStruct struct {
	x, y float64
}

// to create a class function (or equivalent struct function, we specify the type)
func (cs CustomStruct) Abs() float64 {
	return math.Sqrt(cs.x*cs.x + cs.y*cs.y)
}

func main() {
	// create a class and use the function
	cs := CustomStruct{3,4}
	fmt.Println(cs.Abs())
}
```

- if there is a need to change the value in the struct, then take it as a pointer
```go
package main

import (
	"fmt"
	 "math"
	)

// struct can be declared outside of the function scope as well
type CustomStruct struct {
	x, y float64
}

// to create a class function (or equivalent struct function, we specify the type)
// but if we want to change the value of this class, then we need to take it as pointer variables
func (cs *CustomStruct) Abs() float64 {
	cs.x *= 10
	cs.y *= 10
	return math.Sqrt(cs.x*cs.x + cs.y*cs.y)
}

func main() {
	// create a class and use the function
	cs := CustomStruct{3,4}
	fmt.Println(cs.Abs())
}
```

- but while calling the functions, we can invoke it using both the functions - value functions and pointer functions using this syntax only: 
```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```


- Misc understandings
```go
package main

import "fmt"

// struct can be declared outside of the function scope as well
type CustomStruct struct {
	x, y float64
}

// generally default values are not accepted in Go, hence we create a constructor and use that to create the go struct
func NewCustomStruct() *CustomStruct {
	// default values
	cs := CustomStruct{
		x: 2, y: 3,
	}
	return &cs
}

func main() {
	// clearing an empty struct
	st := CustomStruct{}
	// fmt.Println(st == nil) // THROWS AN ERROR
	fmt.Println(st)

	/////// is make possible? this is NOT POSSIBLE because make can also be used on maps, channels and slices
	// st2 := make(CustomStruct{})
	// fmt.Println(st == nil, st2 == nil)

	// but new() can be used --- returns a pointer to the struct
	newSt := new(CustomStruct)
	fmt.Printf("%T | %v\n", newSt, newSt)

	// difference between the new and regular pointer operation then? 
	// NO DIFFERENCE between these 2, but in new, we cannot set the values of X and Y
	// generally new is not being used because that makes us set values of all again
	neww := new(CustomStruct)
	var neww2 = &CustomStruct{}
	fmt.Println(neww, neww2)

	// constructor for this struct
	defStruct := NewCustomStruct()
	fmt.Println(defStruct)
}
```

- Interfaces example
```go
package main

import "fmt"

// dummy interface implementing a single function
type Animal interface {
	shout() string
}

// now there are some classes and constructors which returns a pointer to a class which implements this interface
type Cat struct {
	animalType string
}

func (ct *Cat) shout() string {
	return fmt.Sprintf("Hi, I am a %v", ct.animalType)
}

func NewCat() Animal {
	return &Cat{animalType: "cat"}
}

// a similar dog example
type Dog struct {
	animalType string
}

func (ct *Dog) shout() string {
	return fmt.Sprintf("Hi, I am a %v", ct.animalType)
}

func NewDog() Animal {
	return &Dog{animalType: "dog"}
}

// driver code to simulate different classes
func main() {
	// create a new class and see if it is implementing it correctly!
	ct := NewCat()
	dg := NewDog()

	// fmt.Println(ct.shout())
	// fmt.Println(dg.shout())

	// want to maintain an array of animals, but we do not want to know the type of underlying elements
	var animalsArray = make([]Animal, 0)

	// appending the animals into the array
	animalsArray = append(animalsArray, ct)
	animalsArray = append(animalsArray, dg)

	// iterate the array and call the common function of it
	for _, val := range animalsArray {
		// how to find out the type from this == using a typecasting function okay field
		if cat, ok := val.(*Cat); ok {
			fmt.Println("Cat specific field:", cat.animalType)
		}

		fmt.Println(val.shout())
	}
}

```