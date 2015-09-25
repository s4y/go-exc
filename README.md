# exc: exception handling for Golang

[Go](https://golang.org/) has exceptions… almost. [**panic**, **recover**, and **defer**](http://blog.golang.org/defer-panic-and-recover) provide the basic control flow semantics, but the language lacks a nice way to catch specific panicked error values and switch between panic-style error handling and returning errors as values.

Though Go convention is to return errors as values whevever possible, this leads to tremendously awkward code:

```go
func safelyDoThings() (ret string, err error) {
  intermediateValue, err := getStarted()
  if err != nil {
    return
  }
  
  err = intermediateValue.Finagle(true)
  if err != nil {
    return
  }
  
  ret = intermediateValue.Finalize()
  return
}
```

go-exc aims to make it easy to use panics for control flow, simplifying your code.

## Examples

To catch any error, just call `exc.Catch()` with a function:

```go
err := exc.Catch(func() { doThings(true) })
```

`err` will be an `exc.Panic`, which has a `Value` property (containing the original panicked value) and `Stack` (a string representation of the stack at the time the panic was thrown).

To catch only certain errors, pass an instance of the error type to `exc.CatchOnly()`:

```go
type PotatoError struct{}

err := exc.CatchOnly(doMoreThings, PotatoError{})
```

## Caveats

- Exc will **never** catch runtime errors, like out-of-bounds array accesses. This is by design.
