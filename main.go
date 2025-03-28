package main

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"
)

// A timing decorator
var timingDecorator = Cravat{
	Show: func(fn reflect.Value, args []reflect.Value) []reflect.Value {
		start := time.Now()
		result := fn.Call(args)
		duration := time.Since(start)
		fmt.Printf("it took %v\n", duration)
		return result
	},
}

// SliceWithSleep creates slice of n+1 elements from [0, n]
func SliceWithSleep(n int) []int {
	result := []int{}
	for len(result) < n {
		result = append(result, len(result))
	}
	return append(result, n)
}

func main() {
	decoratedGenAndSortBigArrayOfInt, ok := SimplerTimer(SliceWithSleep)
	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}
	// Call the decorated function
	res := decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println("SimplerTimer: ", len(res), res[len(res)-1])

	decoratedGenAndSortBigArrayOfInt, ok = PutCravat(timingDecorator, SliceWithSleep)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	res = decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println("Cravat", len(res), res[len(res)-1])

	decoratedGenAndSortBigArrayOfInt, ok = SimplerTimerAddExtra(SliceWithSleep, 500)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	res = decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println("SimplerTimerAddExtra", len(res), res[len(res)-1])
}
