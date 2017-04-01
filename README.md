# go-check-utils
Golang utilities to be used with the gopkg.in/check.v1 unit testing package

## Deep Equals
This is an implementation of gopkg.in/check.v1.Checker that allows the developer to provide custom deep equals for specified types when using the gopkg.in/check.v1 library to do an assert equal. The custom deep equals works recursively through the field hierarchy when comparing structs

### Sample Usage
```
import "github.com/hasSalil/go-check-utils/deepequals"
...
...

c.Assert(&obtained, deepequals.DeltaDeepEquals, &expected)
// IMPORTANT!!! Make sure that when passing structs or primitive types to Assert
// with de.DeltaDeepEquals, always pass in the *address* of the struct/primitve
// or a pointer. If you do not, then the Assert call will panic


// The default deepequals.DeltaDeepEquals object uses a float delta of 0.001
// and a time granularity of 1 second. You can easily override this, or add
// your custom equals checker for specific types using:
deepequals.DeltaDeepEquals.UseFloatDelta(0.0001)
deepequals.DeltaDeepEquals.UseTimeGranularity(time.Minute)
deepequals.DeltaDeepEquals.WithDeepEqualForType(
  ty reflect.Type,
  equals func(a unsafe.Pointer, b unsafe.Pointer) bool,
)
```
