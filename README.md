# Powerfunc

Give your functions superpowers!

Did you know that you can add methods to functions in Go? This package provides a set of methods that can be used to decorate
functions with additional functionality.

```go
var adder = func(a, b int) int { return a + b }

var f powerfunc.Func2Value = adder
add2 := f.Curry(2)
add2(3) // 5

var asyncCall = func(ctx context.Context) (int, error) {
    // do something
    return 1, nil
}

var asyncF powerfunc.CtxFuncResult = asyncCall
asyncF.WithTimeout(1 * time.Second)
 .OnErr("async call") // error will now be "async call: %w"
 .Map(add2) // will apply add2 to the result of asyncCall
 .Exec(ctx) // Convenience method if you find ()() weird.
```

## Why?

This package is mostly an experiment to see how far we can push the language. **It is not intended to be used in production**, or at the very least, not in performance critical code (this creates a lot of function calls and allocations, since it's just higher order funcs under the hood).

### What's cool about it!

#### Namespacing

The utility functions are scoped by the type of the function they decorate. This reduces namespace pollution and makes it easier to find the functions you need:

- Need to curry a function? Just look for the `Curry` method, it's available on _all_ functions where it makes sense.
- Need to add a timeout to a function? Look no further than the `WithTimeout` method, available on _all_ functions that accept a context.
- Want to discover all the utilities for your function? Your IDE will happily provide them by just typing `.` after your function name.

#### Composition

Still a function, just with superpowers. The functions returned by the utilities are still functions, so you can compose them with other functions, pass them around, etc. This means that powerfuncs are a drop-in replacement for functions in your code base. No adapters, no wrappers, no boilerplate. Just functions.

Want to use powerfuncs? Just change the type of your function and you're done.

```go
var pf powerfunc.FuncResult = myFunc
// pf is now a powerfunc, you can use it anywhere you would
// use myFunc but now with batteries included

// alternatively, you can use the casting syntax, but this can
// make you believe that we're wrapping myFunc with something
// else, which is not the case. `pf` _is_ `myFunc`.
pf := powerfunc.FuncResult(myFunc)
```

### Drawbacks

#### Performance

Using a powerfunc here and there is probably going to be fine. But avoid chaining many utilities together, as each one creates a new function and a new closure, so 6 chained utilities means 6 new functions and 6 new closures. Just inline the damned behaviour yourself at that point.

#### Go type system

The Go LSP still struggles to infer types from functions, so you are going to have to write that type annotation yourself. Honestly, it's quite a pain when you're dealing with a 6 arity function and you're trying to figure out what the hell the type is. I'm hoping that the Go team will improve this in the future.

```go
func GetMetadata(ctx context.Context, client *bigquery.Client, projectID string, datasetID string, tableID string) (*bigquery.TableMetadata, error) {
    // ...
}

// rage:
// - need to count the number of parameters to figure out the type
// - need to correctly place each type in the type signature
var getMeta powerfunc.CtxFunc4Result[*bigquery.TableMetadata, *bigquery.Client, string, string, string] = GetMetadata

return getMeta.
    WithTimeout(1 * time.Second).
    OnErr("get metadata").
    Exec(ctx, client, projectID, datasetID, tableID)
```

## Installation

```bash
go get github.com/jonathanmontane/powerfunc
```
