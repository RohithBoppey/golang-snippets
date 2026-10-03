
# **1️⃣ Basic Go types & functions**

* Go has **basic types**: `int`, `float64`, `string`, `bool`, `struct`, `slice`, `map`, etc.
* Functions are **statically typed**: you must declare argument and return types.

**Python analogy**:

```python
def add(a, b):
    return a + b  # dynamically works for int, float, str
```

In Go, you must explicitly declare:

```go
func Add(a int, b int) int {
    return a + b
}
```

---

# **2️⃣ Interfaces in Go**

* An **interface** defines a set of methods that a type must implement.
* `interface{}` or `any` can hold **any type** (like Python’s dynamic typing).

**Python analogy**:

```python
def print_len(obj):
    print(len(obj))  # works for str, list, etc.
```

Go equivalent using empty interface:

```go
func PrintLen(v any) {
    switch val := v.(type) {
    case string:
        fmt.Println(len(val))
    case []int:
        fmt.Println(len(val))
    }
}
```

---

# **3️⃣ Slices & pointers**

* **Slice**: dynamic array (`[]int`, `[]string`) — like Python lists.
* **Pointer**: reference to a variable; needed when you want to modify original data.

**Python analogy**:

```python
lst = [1,2,3]
def modify(l):
    l.append(4)  # modifies original list
modify(lst)
print(lst)  # [1,2,3,4]
```

Go equivalent with slice (already reference-like):

```go
func Modify(s []int) {
    s = append(s, 4)
}
lst := []int{1, 2, 3}
Modify(lst)
fmt.Println(lst) // [1,2,3,4]
```

For single variables (int, struct), you need pointers to modify:

```go
func Increment(x *int) {
    *x++
}
```