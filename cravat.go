package main

import "reflect"

// Cravat sets a code block to be run before a call, and one for afterwords
type Cravat struct {
	Show func(reflect.Value, []reflect.Value) []reflect.Value // Code to execute before
}

// PutCravat implements the decorator
func PutCravat[T interface{}](c Cravat, fn T) (T, bool) {
	fnValue := reflect.ValueOf(fn)
	fnType := reflect.TypeOf(fn)

	if fnType.Kind() != reflect.Func {
		return fn, false
	}

	return (reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		return c.Show(fnValue, args)
	}).Interface().(T)), true
}
