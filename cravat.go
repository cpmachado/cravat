package main

import "reflect"

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
