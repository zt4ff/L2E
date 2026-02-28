# Testing Output in Go: Capturing `fmt` Writes in Tests

When writing tests in Go, one common challenge is verifying what your program prints to standard output. Since `fmt.Printf` writes directly to `os.Stdout`, you can't simply "read" the output after the fact, but you can intercept it. This article walks you through the technique step by step.

## The Core Problem

Consider a function like this:

```go
func PrintSomething(str string) {
    fmt.Printf("Print something: %s", str)
}
```

The function writes to `os.Stdout`. There's no return value to assert against, and `os.Stdout` is a `*os.File` meaning you can't read from it directly after writing. So how do you test it?


## The Solution: Pipe Swapping

The trick is to temporarily replace `os.Stdout` with the **write end of a pipe**, then read what was written from the **read end**. Here's how it works conceptually:

1. Save the original `os.Stdout`
2. Create an `os.Pipe()`
3. Point `os.Stdout` at the pipe's writer
4. Call your function (it now writes into the pipe)
5. Close the writer to signal EOF
6. Restore the original `os.Stdout`
7. Read from the pipe's reader into a buffer
8. Assert against the buffer's contents



## Step-by-Step Implementation

### Step 1: The Function Under Test

```go
package main

import "fmt"

func PrintSomething(str string) {
    fmt.Printf("Print something: %s", str)
}
```

Note that this function uses `fmt.Printf`, which writes to `os.Stdout` by default. That's the target we'll intercept.

### Step 2: Set Up the Test Structure

```go
package main

import (
    "bytes"
    "os"
    "testing"
)

func Test_PrintSomething(t *testing.T) {
    type testcase struct {
        input string
        want  string
    }

    tt := map[string]testcase{
        "happy path": {
            input: "testing",
            want:  "Print something: testing",
        },
    }

    for name, tc := range tt {
        t.Run(name, func(t *testing.T) {
            // ... 
        })
    }
}
```

### Step 3: Save and Replace `os.Stdout`

Inside the test loop, save the original stdout and create a pipe:

```go
old := os.Stdout

r, w, err := os.Pipe()
if err != nil {
    t.Fatal(err)
}

os.Stdout = w
```

From this point on, anything written to `os.Stdout` (including by `fmt.Printf`) goes into the pipe.

### Step 4: Call the Function

```go
PrintSomething(tc.input)
```

### Step 5: Close the Writer and Restore Stdout

```go
w.Close()
os.Stdout = old
```

Closing `w` is important. Without it, `buf.ReadFrom(r)` will block forever waiting for more data.

### Step 6: Read the Captured Output

```go
var buf bytes.Buffer
buf.ReadFrom(r)

got := buf.String()
```

### Step 7: Assert

```go
if got != tc.want {
    t.Errorf("got: %q, want: %q", got, tc.want)
}
```


## Full Working Example

```go
package main

import (
    "bytes"
    "fmt"
    "os"
    "testing"
)

func PrintSomething(str string) {
    fmt.Printf("Print something: %s", str)
}

func Test_PrintSomething(t *testing.T) {
    type testcase struct {
        input string
        want  string
    }

    tt := map[string]testcase{
        "happy path": {
            input: "testing",
            want:  "Print something: testing",
        },
        "empty string": {
            input: "",
            want:  "Print something: ",
        },
    }

    for name, tc := range tt {
        t.Run(name, func(t *testing.T) {
            old := os.Stdout
            r, w, err := os.Pipe()
            if err != nil {
                t.Fatal(err)
            }
            os.Stdout = w

            PrintSomething(tc.input)

            w.Close()
            os.Stdout = old

            var buf bytes.Buffer
            buf.ReadFrom(r)

            got := buf.String()
            if got != tc.want {
                t.Errorf("got: %q, want: %q", got, tc.want)
            }
        })
    }
}
```

---

## A Cleaner Approach: Accept an `io.Writer`

If you control the function's design, a much cleaner pattern is to accept an `io.Writer` as a parameter instead of writing directly to `os.Stdout`. This eliminates the need for pipe swapping entirely.

```go
func PrintSomething(w io.Writer, str string) {
    fmt.Fprintf(w, "Print something: %s", str)
}
```

Now in your test, just pass a `bytes.Buffer` directly:

```go
func Test_PrintSomething(t *testing.T) {
    var buf bytes.Buffer
    PrintSomething(&buf, "testing")

    got := buf.String()
    want := "Print something: testing"

    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}
```

And in production, call it with `os.Stdout`:

```go
PrintSomething(os.Stdout, "hello")
```

This is idiomatic Go. The standard library itself uses this pattern extensively (e.g., `fmt.Fprintf`, `http.ResponseWriter`). **Prefer this approach when writing new code.**

---

## Can't I Read Directly from `os.Stdout`?

No.

`os.Stdout` is a `*os.File` that wraps the OS file descriptor `1` (stdout). It is **write-only** from your program's perspective. The OS doesn't provide a way for a process to read back what it has already written to its own stdout.

What you *can* do is replace where stdout points (as shown above with pipe swapping), or use a higher-level abstraction like `io.Writer` injection. There is no syscall or Go API that lets you "peek" at already-written stdout data.

The pipe technique works precisely because you're not reading from the original stdout, you're redirecting the write to a buffer you control.

---

## Summary

| Approach | When to Use |
|---|---|
| **Pipe swapping** | Testing existing functions that write directly to `os.Stdout` |
| **`io.Writer` injection** | The cleanest and most testable design |
| **Read from `os.Stdout` directly** | Not possible since stdout is write-only from the process perspective |

When working with legacy code or third-party functions that hardcode `os.Stdout`, pipe swapping is your best tool. For new code, designing functions to accept an `io.Writer` makes them inherently testable without any redirection gymnastics.