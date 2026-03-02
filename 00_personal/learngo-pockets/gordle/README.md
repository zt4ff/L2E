## Sentinel error
aims to treat errors as a const value available through the package so comparision can work

```go
msg := "error message"
a := errors.New(msg)
b := errors.New(msg)

a == b // false
```

```go
type customError string

func (c customError) Error() string {
    return string(e)
}

const CustomError = customError("this is a custom error")

// you can now compare like this
a := CustomError
b := CustomError

a == b // true
```