Basic syntax: `func (input params) (output params)`

### Variadic parameters
- Similar to function in python, we need to have default valued parameters at the end only 
- so generally the normal named arguments with types come first, only 1 would be there at the end, which is of array type mostly

so if this: 
```go
func CreateTables(opts ...bool) error {}
```
 
only 1 optional array of params can be given, and they should be of the same type, in this case: 
- You can pass zero or more booleans
- All arguments passed to opts must be of type bool
- There cannot be another parameter after opts