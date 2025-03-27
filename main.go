package main

import (
	"fmt"
	"log/slog"
	"time"
)

// A timing decorator
var timingDecorator = Cravat{
	Before: func() interface{} {
		v := time.Now()
		return v
	},
	After: func(start interface{}) {
		rStart := start.(time.Time)
		duration := time.Since(rStart)
		fmt.Printf("it took %v\n", duration)
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
	fmt.Println(len(res))

	decoratedGenAndSortBigArrayOfInt, ok = PutCravat(timingDecorator, SliceWithSleep)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	res = decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println(len(res))
}
