package cravat

import "reflect"

// Cravat is an interface to inject a function while maintaining fashion
// safety
type Cravat interface {
	Show(reflect.Value, []reflect.Value) []reflect.Value // Code to execute with function
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
