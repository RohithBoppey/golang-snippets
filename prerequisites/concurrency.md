- goroutine is a lightweight thread managed by "Go runtime"
  - lightweight way to run a function simultaneously with the main execution flow
- Think of it like asking Go: “Hey, run this function on the side, without blocking me. I’ll continue doing my work.”
- Go says: “Sure, I’ll handle it. I’ll manage the scheduling, memory, everything. You just say ‘go f()’.”
- f(x,y,z) ---> evaluation happens in the current goroutine, but the execution of "f" happens in the new goroutine
- The goroutines run in the same address space, so the access to shared memory should be syncronized. (can use "sync" for this, otherwise "race conditions" happen)
- Generally when run `f(x,y,z)` -- syncronous and blocking - wait until it finishes; but if ran `go f(x,y,z)` -- start execution of this immediately and do not wait for it, meanwhile continue with the remaining flow in the main thread

Goroutines are:

- extremely cheap (few KBs of memory)
- can grow/shrink dynamically
- can be created in thousands or millions
- managed by Go runtime, not by the OS?

Practical example: New Thread is like opening google chrome window (tiresome and heavy), but goroutine is like opening a new tab (in an existing thread).

- Interesting code example about Golang:

```go
package main

import (
	"fmt"
// 	"time"
)

func say() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go say()
	fmt.Println("Hello from main")
    // time.Sleep(time.Second)
}
```

- with this snippet, we are unable to see "Hello from Goroutine" message
- the flow is: go just says: run this function say() whenever you have time i.e. like scheduling a task for future
- but when the main() function execution happens i.e. when the main goroutine completes execution, all goroutines are abruptly killed even if they have not completed their run
- so time.Sleep() gives the program enough time to complete the goroutine execution.

##### Channels

- channels are a medium in which sending and receiving values is possible
- default way is: blocking channel only -- allows goroutines to send data without locks or conditional variables
- Ex: take an array of integers -> leverage concurrent goroutines to make the sum of half and half, then read from channels to make the sum

```go
package main

import (
    "fmt"
)

// read the numbers from the array and put the sum in the channel
func summer(arr []int, ch chan int) {
    s := 0
    for _, val := range(arr) {
        s += val
    }
    ch <- s;
}

func main() {
    // creating a channel of int
    ch := make(chan int)
    // creating an array
    arr := []int{1,2,3,4,5}

    // now split the array into left and right to make the sums and put them in channel
    // and read them from the main execution code
    go summer(arr[: len(arr) / 2], ch)
    go summer(arr[len(arr) / 2 : ], ch)

    // now read the answers from the channel for further processing
    left, right := <-ch, <-ch
    fmt.Println(left, right)
}
```

- note that channels are logically FIFO systems (queue) - first in is first out
- but in the case of unbuffered channels (send and receive should happen at the same time) i.e. there is no storage in the channel i.e. 2 parties are involved in this transaction at the same time (sender: wants to send something, but there is a receiver: who wants to receive something)
- so whenever we are using unbuffered channels, that specific goroutine is blocked until the condition is met - someone sent to channel and someone is there to recieve it

###### Step 1: main starts

```go
ch := make(chan int)
```

###### Step 2: goroutines start

```go
go summer(leftPart, ch)
go summer(rightPart, ch)
```

They start running **whenever the Go scheduler decides**.

---

###### Step 3: main tries to receive

```go
left, right := <-ch, <-ch
```

Now main is doing:

1️⃣ `<-ch` → “I’m ready to receive”
2️⃣ waits until **some goroutine sends**

This does not mean: “First goroutine’s result goes to left”.  
It means: “Whichever goroutine reaches ch <- s first will unblock the first <-ch”

---

###### Step 4: one goroutine reaches `ch <- s`

Whichever goroutine hits this line **first**:

```go
ch <- s
```

- sees that main is already waiting
- immediately hands over the value
- unblocks main
- **this value becomes `left`**

---

###### Step 5: second receive

Now main executes the second `<-ch`:

- waits again
- the remaining goroutine sends
- value becomes `right`

---

###### ❗ Important clarification

> ❌ “First goroutine created sends first”
> ❌ “Smaller array finishes first”
> ❌ “Left slice sends first”

All of these are **false assumptions**.

✔️ **Only this matters**:

> Which goroutine reaches `ch <- s` first

And that is **not predictable**.

---

In the case of unbuffered channel, what happens in the following cases:

Case 1: Sender arrives first

```go
ch <- s
```

- Sender reaches the channel
- Looks around
- ❌ No receiver yet
- **Sender WAITS (blocks) -- this means this goroutine execution is stopped**

* It just stands there holding the value `s`.

Case 2: Receiver arrives first

```go
<-ch
```

- Receiver reaches the channel
- ❌ No sender yet
- **Receiver WAITS (blocks)**

It just stands there waiting to receive _something_.

---

Better example to understanding how unbuffered channels help in coordination and not syncronization

```go
package main

import (
	"fmt"
	"time"
)

func worker(name string, ch chan string) {
	fmt.Println(name, "started work")
	time.Sleep(2 * time.Second) // simulate work

	fmt.Println(name, "reached checkpoint")
	ch <- name // MUST meet main here

	fmt.Println(name, "continues after checkpoint")
}

func main() {
	ch := make(chan string)

	go worker("A", ch)
	go worker("B", ch)

	fmt.Println("main waiting at checkpoint")

	first := <-ch
	fmt.Println("main met worker:", first)

	second := <-ch
	fmt.Println("main met worker:", second)

	fmt.Println("main done")
}
```

Synchronous would mean:

- worker A runs fully
- then worker B runs fully
- then main runs

❌ That’s **not** what happened.

Instead:

- A and B ran **at the same time**
- main paused **only at the checkpoint**
- workers paused **only at the checkpoint**

This is **coordination**, not serialization.

##### Buffered channels (Queue)

- capacity would be there for the channel
- should be closed from the sender side only
- can be checked using: `v, ok := <-ch`
- example code to best explain: use an buffered channel to populate fibonacci sequence and put it in the channel, and downstream consume it

```go
package main

import (
    "fmt"
)

func fibo(n int, ch chan int){
    a, b := 0, 1    // init values
    for _ = range(n) {
        ch <- a     // inserting the first value in the channel
        a, b = b, a + b    // now swap for the next iteration
    }
    close(ch)
}

func main() {
    // buffered channel
    ch := make(chan int, 5)
    // populate the buffered channel with the fibonacci sequence
    go fibo(cap(ch), ch)

    // using while loop in go
    for {
        // if the channel is empty you can break it
        v, ok := <- ch
        if !ok {
            fmt.Println("The channel is closed")
            break
        }
        fmt.Println(v)
    }

    fmt.Println("main done")

}
```

- note that even with the buffered channels, it does not act as a proper queue
- for instance if commenting the: `go fibo()` code and running it, the channel is never populated with any value, and now when we are reading from it, it says all goroutines are asleep - deadlock

```go
package main

import (
    "fmt"
)

func fibo(n int, ch chan int){
    a, b := 0, 1    // init values
    for _ = range(n) {
        ch <- a     // inserting the first value in the channel
        a, b = b, a + b    // now swap for the next iteration
    }
    close(ch)
}

func main() {
    // buffered channel
    ch := make(chan int, 5)
    // populate the buffered channel with the fibonacci sequence
    go fibo(cap(ch), ch)

    // using while loop in go
    for {
        // if the channel is empty you can break it
        v, ok := <- ch
        if !ok {
            fmt.Println("The channel is closed")
            break
        }
        fmt.Println(v)
    }

    fmt.Println("main done")

}
```

- if we are using unbuffered channels, then cap(ch) = 0, and if you hardcode the number, then handshake kind of thing happens: goroutine sends 1, main prints 1, and so on

```go
package main

import (
    "fmt"
)

func fibo(n int, ch chan int){
    a, b := 0, 1    // init values
    for _ = range(n) {
        ch <- a     // inserting the first value in the channel
        a, b = b, a + b    // now swap for the next iteration
    }
    close(ch)
}

func main() {
    // unbuffered channel
    ch := make(chan int)
    // populate the buffered channel with the fibonacci sequence
    go fibo(5, ch)

    // using while loop in go
    for {
        // if the channel is empty you can break it
        v, ok := <- ch
        if !ok {
            fmt.Println("The channel is closed")
            break
        }
        fmt.Println(v)
    }

    fmt.Println("main done")

}
```

- in this example, if you do not close the channel, then again a deadlock happens, because the main goroutine is waiting to receive, but there is nothing to send

---

- `select` can block a goroutine from running until one of the statements run
  - the statements can have multiple if conditions and until one of it is run, the goroutine is blocked

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
    // 	running the loop infinitely
	for {
	   // block the goroutine until one of the following operation is possible
		select {
		  //  try inserting into channel -- since blocking, there must be a receiver ready too -- ready in fmt.Println() statement
		case c <- x:
          //  if successful, update
			x, y = y, x+y
		case <-quit:
          // if quit signal comes first, then quit the loop
			fmt.Println("quit")
			return
		}
	}
}

func main() {
    // unbuffered channels, so blocking the couroutine
	ch := make(chan int)
	quit := make(chan int)
    // 	running a goroutine which is blocked waiting for input
	go func() {
		for i := 0; i < 10; i++ {
		    //  take one value from the channel and print it
			fmt.Println(<-c)
		}
        // once you're done, then send the quit signal
		quit <- 0
	}()

	fibonacci(c, quit)
}
```

- if I add a close statement right after inserting the value in the channel the first time, something interesting happens:

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 2
    // 	running the loop infinitely
	for {
	   // block the goroutine until one of the following operation is possible
		select {
		  //  try inserting into channel -- since blocking, there must be a receiver ready too -- ready in fmt.Println() statement
		case c <- x:
          //  if successful, update
			x, y = y, x+y
			close(c)
		case <-quit:
          // if quit signal comes first, then quit the loop
			fmt.Println("quit")
			return
		}
	}
}

func main() {
    // unbuffered channels, so blocking the couroutine
	c := make(chan int)
	quit := make(chan int)
    // 	running a goroutine which is blocked waiting for input
	go func() {
		for i := 0; i < 10; i++ {
		    //  take one value from the channel and print it
		    v, ok := <-c
			fmt.Println(v, ok)
		}
        // once you're done, then send the quit signal
		quit <- 0
	}()

	fibonacci(c, quit)
}
```

```text
1 true
0 false
0 false
0 false
0 false
0 false
0 false
panic: send on closed channel
```

- because reading from the closed channel gives 0 and does not cause panic (but you can check using ok signal)
- but sending the signal into a closed channel throws a panic error
- before the loop continues in the goroutine, we are rapidly reading from the main execution causing this 0s
- once the loop in the goroutines comes again to insert into the channel, then panic comes
- Reading from a closed channel is not synchronized with the select in the other goroutine. Therefore the number of 0s printed is nondeterministic.

---

### Examples

---

### Questions

1. What is the flow from creating a goroutine -> Go runtime managing it (How?) -> How go speaks with OS level threads and cores for this?

2. Does using unbuffered channels lead to syncronous execution of code?

- no
  > **Unbuffered channels do NOT make your program synchronous.**
  > They make **specific goroutines wait at specific points**.
- That goroutine's execution is stopped i.e. that specific goroutine is only blocked and will be waiting (for someone to read from that channel) for unblocked to continue execution

```text
time →
main:        go summer()  go summer()   <-ch (BLOCKS)   <-ch
summer1:                  compute        ch<-3 (MEET)
summer2:                  compute        ch<-12
```

What happens:

- `main` blocks at `<-ch`
- **BUT** both `summer` goroutines are still running
- CPU switches between goroutines freely
- Once a `summer` sends → main unblocks

👉 This is **not synchronous execution**
👉 It is **coordination**

---

Sync Mutex - used to prevent race conditions among goroutines, allows only a single goroutine to access the variables at a single time.

- `defer`generally means no matter how the function exits, you must always execute the "defer" code
- so generally `defer` is used to schedule closing of the mutexes opened
- defer makes code error free in the future; if we do not use defer and not close the lock, like in a return statement, we are risking a deadlock from happening

###### Points about differ

- `defer` can be logically assumed to be `finally` keyword (only difference being differ is function scoped, but finally is block scoped)
- `defer` runs on panic as well, so both in try blocks and catch blocks, `defer` runs
- if multiple `defer` statements are created, then it is executed in the stack order (LIFO order)

###### Mental model to keep forever 🧠

- Mutex = “only one goroutine allowed inside”
- Maps are not thread-safe (inherently)
- Reading and writing both count as “inside”, so sync mutex is needed
- defer Unlock() = “don’t forget to clean up”

```go
package main
import (
    "fmt"
    "sync"
    "time"
)

// custom struct which is having mutexes
type CustomMap struct {
    mutex sync.Mutex
    mp map[string]int
}

// constructor for the function
func newCustomMap() *CustomMap {
    // create a new map
    mpp := make(map[string]int)
    cmp := &CustomMap{}
    // update the map
    cmp.mp = mpp
    return cmp
}

// METHODS on the custom struct
// increment func
func (cmp *CustomMap) Inc(key string) {
    // lock the map
    cmp.mutex.Lock()
    // register an lock close the function exit
    defer cmp.mutex.Unlock()
    // now udpate the key of the string
    cmp.mp[key]++
}

// get the value function
func (cmp *CustomMap) Val(key string) int {
    // lock the mutex and register the unlock on the function exit
    cmp.mutex.Lock()
    defer cmp.mutex.Unlock()
    // return the the key now
    val, _ := cmp.mp[key]
    return val
}

func main() {
    // mutexes
    cmp := newCustomMap()
    fmt.Println(cmp, "created")

    // use the goroutines to increment call the increment function
    const NUM_GOROUTINES = 100
    const MAP_KEY = "something"
    for _ = range(NUM_GOROUTINES){
        go cmp.Inc(MAP_KEY)
    }

    // enough time for the goroutines to run
    time.Sleep(time.Second)

    // check the value now
    fmt.Println(cmp.Val(MAP_KEY))
}
```

- time.Sleep() is for the goroutine lifecycle, not related to mutexes
- without time.Sleep() command, this is how the program flow is gonna be: start execution -> spawns goroutines -> print the value -> exit the main function, it does not care whether all the goroutines have completed execution or not
- so including this time.Sleep() helps enough buffer time for the goroutines to complete their execution
- **goroutines do not the keep the function execution alive, hence time.Sleep() is needed**

###### Core example explaining mutexes neatly

```go
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CustomMap struct {
	mtx sync.Mutex
	mp  map[string]int // to store the interface
}

// creating a global variable for maintaining mutex state
var cmp = &CustomMap{mp: make(map[string]int)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// check if the url is already present or not, not defering until the end of execution
	// since in the recursive loop, a new lock is anyways required
	// so we must close the lock before recursion starts
	cmp.mtx.Lock()
	if _, ok := cmp.mp[url]; ok {
		// already present in the map
		fmt.Println(url, "already processed, exiting this loop!")
		cmp.mtx.Unlock()
		return
	}
	// mark and unlock
	cmp.mp[url] = 1
	cmp.mtx.Unlock()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```

#### sync.Mutex

- Exclusive lock
- **Only one goroutine** can hold it at a time
- Readers and writers are treated the same

* use this by default without much overthinking! ✅

```go
type Counter struct {
    mu sync.Mutex
    v  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.v++
    c.mu.Unlock()
}
```

#### sync.RWMutex

- Two kinds of locks:
  - `RLock()` → many readers allowed
  - `Lock()` → only one writer, blocks readers
- use when:
  - **Read-heavy workloads**
  - Writes are rare
  - Data is read far more often than written

```go
type Cache struct {
    mu sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(k string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[k]
}

func (c *Cache) Set(k, v string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[k] = v
}
```

---

##### 🔹 `sync.Mutex` vs `sync.RWMutex` (decision table)

| Scenario                 | Use     |
| ------------------------ | ------- |
| Simple shared variable   | Mutex   |
| Mostly reads, few writes | RWMutex |
| Complex invariants       | Mutex   |
| Unsure                   | Mutex   |

---

##### 2️⃣ WaitGroups (this replaces `time.Sleep`)

##### The problem WaitGroup solves

```go
go doWork()
fmt.Println("done") // runs immediately
```

Main exits → goroutines die.

---

##### What a WaitGroup is (mental model)

Think of it as a **counter**:

- `Add(n)` → I am waiting for `n` things to be done
- `Done()` → one thing finished
- `Wait()` → block until counter becomes zero (or wait until all are done)

```go
// single goroutine
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
    fmt.Println("hello from goroutine")
}()

wg.Wait()
fmt.Println("main exits")
```

```go
// multiple goroutines
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        fmt.Println("worker", i)
    }(i)
}

wg.Wait()
```

- **Add(1) is a very important step!!!**

##### Practical example explaining the difference between the time.sleep and wait group functionality
<img width="749" height="493" alt="image" src="https://github.com/user-attachments/assets/3f3438d6-b6e9-4ce5-bdaf-91a7142650d2" />
