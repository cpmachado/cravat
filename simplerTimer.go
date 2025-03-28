package main

import (
	"fmt"
	"reflect"
	"time"
)

func SimplerTimer[T interface{}](fn T) (T, bool) {
	fnValue, fnType := reflect.ValueOf(fn), reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return fn, false
	}
	return (reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		start := time.Now()
		result := fnValue.Call(args)
		duration := time.Since(start)
		fmt.Printf("it took %v\n", duration)
		return []reflect.Value(result)
	}).Interface().(T)), true
}

func SimplerTimerAddExtra[T interface{}](fn T, extra int) (T, bool) {
	fnValue, fnType := reflect.ValueOf(fn), reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return fn, false
	}
	return (reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		start := time.Now()
		result := fnValue.Call(args)
		v := result[0]
		updatedSlice := reflect.Append(v, reflect.ValueOf(extra))
		duration := time.Since(start)
		fmt.Printf("it took %v\n", duration)
		return []reflect.Value{updatedSlice}
	}).Interface().(T)), true
}
