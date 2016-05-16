package reflection

import (
	"fmt"
	"reflect"
)

// -----------------------------------------------------------------------------

// PrintStruct prints the `testStruct` description including its type and its
// fields. It returns `false` if `testStruct` is not a `struct`.
func PrintStruct(testStruct interface{}) bool {
	concreteVal := reflect.ValueOf(testStruct)

	if concreteVal.Kind() != reflect.Struct {
		return false
	}

	fmt.Printf("Structure Name: %s\n", concreteVal.Type().Name())
	for i := 0; i < concreteVal.NumField(); i++ {
		fmt.Printf(
			"Field %d: %s - %s\n",
			i,
			concreteVal.Type().Field(i).Name,
			concreteVal.Type().Field(i).Type,
		)
	}
	return true
}
