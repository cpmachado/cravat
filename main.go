package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"reflect"
	"slices"
	"time"
)

func main() {
	decoratedGenAndSortBigArrayOfInt, ok := PutCravat(timingDecorator, GenAndSortBigArrayOfInt)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	result := decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println("Result:", result)
}

// Cravat sets a code block to be run before a call, and one for afterwords
type Cravat struct {
	Before func() interface{} // Code to execute before
	After  func(interface{})  // Code to execute after, which has return as a parameter
}

// PutCravat implements the decorator
func PutCravat[T interface{}](c Cravat, fn T) (T, bool) {
	fnValue := reflect.ValueOf(fn)
	fnType := reflect.TypeOf(fn)

	if fnType.Kind() != reflect.Func {
		return fn, false
	}

	return (reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		arg := c.Before()
		result := fnValue.Call(args)
		c.After(arg)
		return result
	}).Interface().(T)), true
}

// A timing decorator
var timingDecorator = Cravat{
	Before: func() any {
		return time.Now()
	},
	After: func(start interface{}) {
		rStart := start.(time.Time)
		duration := time.Since(rStart)
		fmt.Printf("it took %v\n", duration)
	},
}

// GenAndSortBigArrayOfInt generates and retrieves a sorted array of random
// numbers, of size n between [0, n)
func GenAndSortBigArrayOfInt(n int) []int {
	result := []int{}
	for len(result) < n {
		result = append(result, rand.Intn(n))
	}

	slices.Sort(result)
	return result
}
