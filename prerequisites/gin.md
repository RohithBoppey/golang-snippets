install gin using the command: `go get -u github.com/gin-gonic/gin`

- Sample code
```go
ginDefEngine := gin.Default()

	// Register the routes in the gin
	v1ApiGroup := ginDefEngine.Group("/api/V1")
	{
		v1ApiGroup.POST("/driver/location", handlers.IngestLocationGin(locationQueue))
	}

	// Run the gin handler directly and use it in the server while stopping
	slog.Info("📡 Server listening on :8080")
	if err := ginDefEngine.Run(":8080"); err != nil {
		slog.Error("Server failed:", "error", err.Error())
	}
	server := &http.Server{Addr: ":8080", Handler: ginDefEngine}
```


### 1. The Core Difference: The Function Signature

The fundamental difference lies in **how they accept arguments**.

* **Native (`http.HandlerFunc`):** Go's standard library gives you two separate tools: one to **Read** the request (`*http.Request`) and one to **Write** the response (`http.ResponseWriter`). You must manage both manually.
* **Gin (`gin.HandlerFunc`):** Gin wraps both of those tools into a single "Super Object" called the **Context** (`*gin.Context`). This context holds the request, the writer, and a bunch of helper methods (like `JSON()`, `Bind()`, etc.).

| Feature | Native Go (`net/http`) | Gin Framework |
| --- | --- | --- |
| **Signature** | `func(w http.ResponseWriter, r *http.Request)` | `func(c *gin.Context)` |
| **Reading Body** | `json.NewDecoder(r.Body).Decode(&obj)` | `c.ShouldBindJSON(&obj)` |
| **Writing JSON** | `w.Header().Set(...)` + `w.Write(...)` | `c.JSON(200, obj)` |
| **Passing Data** | Manual context handling | `c.Set("key", val)` / `c.Get("key")` |

---

### 2. Doubt 1: Why `Handler: nil` in Native vs. `Handler: engine` in Gin?

This is the key to understanding how Go's server works.

**The Native Way (`Handler: nil`)**

```go
http.HandleFunc("/path", myFunc)       // Step A: Register route to DefaultServeMux
server := &http.Server{Handler: nil}   // Step B: Start server with "nil"

```

* **The Secret:** When you pass `nil` as the Handler, Go's server automatically uses a global variable called **`DefaultServeMux`**.
* **What happened:** In Step A, `http.HandleFunc` registered your function into that global `DefaultServeMux`. So when the server starts with `nil`, it defaults to looking at that global list.

**The Gin Way (`Handler: ginDefEngine`)**

```go
r := gin.Default()                     // Step A: Create a NEW Router (Engine)
r.POST("/path", myFunc)                // Step B: Register route to THAT Engine
server := &http.Server{Handler: r}     // Step C: Tell Server to use THAT Engine

```

* **The Difference:** Gin does **not** use the global `DefaultServeMux`. It has its own high-performance router.
* **The Wiring:** You must explicitly tell the `http.Server`: *"Do not use the standard global router. Use this Gin Engine (`r`) instead."* If you passed `nil` here, your Gin routes would be ignored!

---

### 3. Doubt 2: Why do we need `ginDefEngine.Run()`?

Actually, **you don't strictly need it** if you are using the manual `server := &http.Server{...}` setup.

* **`ginDefEngine.Run(":8080")`** is just a "Quality of Life" helper method.
* It prints some nice logs.
* It internally creates an `http.Server`.
* It calls `ListenAndServe`.
* **Limitation:** It blocks forever and doesn't return the server instance, making Graceful Shutdown harder.


* **The Professional Way (What you wrote):**
By manually creating the `&http.Server` and passing the Gin engine as the handler, you are bypassing `Run()`. You are taking manual control of the server lifecycle so you can shut it down gracefully.

**So in your code:**
You can actually **delete** the `ginDefEngine.Run(":8080")` line inside your Gin block if you are using the `server.ListenAndServe()` goroutine below it. Currently, you might be trying to start the server twice (which would error because port 8080 is already taken)!

**Correct Pattern for Gin + Graceful Shutdown:**

```go
// 1. Create Engine
r := gin.Default()
r.POST(...)

// 2. Create Server Manually (Don't use r.Run!)
srv := &http.Server{
    Addr:    ":8080",
    Handler: r, // Pass Gin here
}

// 3. Start Server in Goroutine
go func() {
    srv.ListenAndServe()
}()

// 4. Wait for Signal...
<-stop
srv.Shutdown(ctx)

```