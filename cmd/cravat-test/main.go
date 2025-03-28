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

func main() {
	decoratedFunc, ok := SimplerTimer(MakeSliceN)
	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}
	// Call the decorated function
	res := decoratedFunc(2025)
	fmt.Println("SimplerTimer: ", len(res), res[len(res)-1])
	fmt.Printf("\n\n")

	decoratedFunc, ok = cravat.PutCravat(timingDecorator, MakeSliceN)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	res = decoratedFunc(2025)
	fmt.Println("Cravat", len(res), res[len(res)-1])
	fmt.Printf("\n\n")

	decoratedFunc, ok = SimplerTimerAddExtra(MakeSliceN, 500)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	res = decoratedFunc(2025)
	fmt.Println("SimplerTimerAddExtra", len(res), res[len(res)-1])
}
