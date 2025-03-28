package main

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"github.com/cpmachado/cravat"
)

type TimingDecorator struct{}

func (t TimingDecorator) Show(fn reflect.Value, args []reflect.Value) []reflect.Value {
	start := time.Now()
	result := fn.Call(args)
	duration := time.Since(start)
	fmt.Printf("it took %v\n", duration)
	return result
}

// A timing decorator
var timingDecorator = TimingDecorator{}

// MakeSliceN creates slice of n+1 elements from [0, n]
func MakeSliceN(n int) []int {
	result := []int{}
	for len(result) < n {
		result = append(result, len(result))
	}
	return append(result, n)
}

func DecoratedTest[T func(func(int) []int) (func(int) []int, bool)](name string, f T) {
	decoratedFunc, ok := f(MakeSliceN)
	if !ok {
		slog.Error("main/decorator", slog.String("func", name), slog.String("error", "Failed to set decorator"))
	}
	fmt.Printf("---------------------\n")
	// Call the decorated function
	res := decoratedFunc(2025)
	fmt.Printf("%s: %d, %d\n", name, len(res), res[len(res)-1])
}

func main() {
	DecoratedTest("SimplerTimer", SimplerTimer)
	DecoratedTest("Cravat", func(f func(int) []int) (func(int) []int, bool) {
		return cravat.PutCravat(timingDecorator, f)
	})
	DecoratedTest("SimplerTimerAddExtra", SimplerTimerAddExtra)
}
