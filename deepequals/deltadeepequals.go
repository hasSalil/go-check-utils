package deepequals

import (
	"math"
	"reflect"
	"time"
	"unsafe"

	"github.com/hasSalil/customdeepequal"
	ch "gopkg.in/check.v1"
)

// DeltaDeepEquals is the standard deltaDeepEqualsChecker instance
var DeltaDeepEquals = deltaDeepEqualsChecker(0.001, time.Second)

// marginOfErrorDeepEqualsChecker does deep equals between a pair of structs and returns
// true if recursive floats and time values are within the defined margin of error
type marginOfErrorDeepEqualsChecker struct {
	*ch.CheckerInfo
	deepEquals customdeepequal.CustomDeepEquals
}

// deltaDeepEqualsChecker creates a pointer to marginOfErrorDeepEqualsChecker
func deltaDeepEqualsChecker(floatDelta float64, timeGran time.Duration) *marginOfErrorDeepEqualsChecker {
	checker := &marginOfErrorDeepEqualsChecker{
		CheckerInfo: &ch.CheckerInfo{Name: "DeepEquals", Params: []string{"obtained", "expected"}},
	}
	de := customdeepequal.NewCustomDeepEquals()
	f32 := float32(0)
	f64 := float64(0)
	t := time.Now()
	de.RegisterEquivalenceForType(reflect.TypeOf(f32), func(a, b unsafe.Pointer) bool {
		af := *(*float32)(a)
		bf := *(*float32)(b)
		return math.Abs(float64(af-bf)) <= floatDelta
	})
	de.RegisterEquivalenceForType(reflect.TypeOf(f64), func(a, b unsafe.Pointer) bool {
		af := *(*float64)(a)
		bf := *(*float64)(b)
		return math.Abs(af-bf) <= floatDelta
	})
	de.RegisterEquivalenceForType(reflect.TypeOf(t), func(a, b unsafe.Pointer) bool {
		aT := (*time.Time)(a)
		bT := (*time.Time)(b)
		aNano := aT.UnixNano()
		bNano := bT.UnixNano()
		nanoGrain := timeGran.Nanoseconds()
		return (aNano / nanoGrain) == (bNano / nanoGrain)
	})
	checker.deepEquals = de
	return checker
}

// WithDeepEqualForType registers the equals function for the given type
func (checker *marginOfErrorDeepEqualsChecker) WithDeepEqualForType(ty reflect.Type, equals func(a, b unsafe.Pointer) bool) *marginOfErrorDeepEqualsChecker {
	checker.deepEquals.RegisterEquivalenceForType(ty, equals)
	return checker
}

// UseFloatDelta sets the float delta
func (checker *marginOfErrorDeepEqualsChecker) UseFloatDelta(delta float64) *marginOfErrorDeepEqualsChecker {
	f32 := float32(0)
	f64 := float64(0)
	checker.deepEquals.RegisterEquivalenceForType(reflect.TypeOf(f32), func(a, b unsafe.Pointer) bool {
		af := *(*float32)(a)
		bf := *(*float32)(b)
		return math.Abs(float64(af-bf)) <= delta
	})
	checker.deepEquals.RegisterEquivalenceForType(reflect.TypeOf(f64), func(a, b unsafe.Pointer) bool {
		af := *(*float64)(a)
		bf := *(*float64)(b)
		return math.Abs(af-bf) <= delta
	})
	return checker
}

// UseTimeGranularity sets the float delta
func (checker *marginOfErrorDeepEqualsChecker) UseTimeGranularity(timeGran time.Duration) *marginOfErrorDeepEqualsChecker {
	t := time.Now()
	checker.deepEquals.RegisterEquivalenceForType(reflect.TypeOf(t), func(a, b unsafe.Pointer) bool {
		aT := (*time.Time)(a)
		bT := (*time.Time)(b)
		aNano := aT.UnixNano()
		bNano := bT.UnixNano()
		nanoGrain := timeGran.Nanoseconds()
		return (aNano / nanoGrain) == (bNano / nanoGrain)
	})
	return checker
}

// Check implements check.Checker
func (checker *marginOfErrorDeepEqualsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	return checker.deepEquals.DeepEqual(params[0], params[1]), ""
}
