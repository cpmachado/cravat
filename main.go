package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"slices"
	"time"
)

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

func main() {
	decoratedGenAndSortBigArrayOfInt, ok := PutCravat(timingDecorator, GenAndSortBigArrayOfInt)

	if !ok {
		slog.Error("main/decorator", slog.String("error", "Failed to set decorator"))
	}

	// Call the decorated function
	result := decoratedGenAndSortBigArrayOfInt(2025)
	fmt.Println("Result:", result)
}
