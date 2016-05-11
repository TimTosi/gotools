package bitmap

import (
	"fmt"
	"reflect"
)

// -----------------------------------------------------------------------------

// TODO
func valToBitMap(valArrays ...[]interface{}) map[interface{}][]int {
	var bitMap map[interface{}][]int

	for i, valArray := range valArrays {
		for _, val := range valArray {
			if _, ok := bitMap[val]; !ok {
				bitMap[val] = make([]int, len(valArray))
			}
			bitMap[val][i] = 1
		}
	}
	return bitMap
}

// TODO
func appendMatrix(a, b [][]interface{}) [][]interface{} {
	var res [][]interface{}

	fmt.Println("MATRIX TEST------------------------------------")
	fmt.Printf("A: %v \n", a)
	fmt.Printf("B: %v \n", b)

	if len(a) == 0 {
		fmt.Println("MATRIX ZERO LEN A")
		fmt.Printf("\nRES: %v \n", b)
		fmt.Println("MATRIX OVER--------------------------------------")
		return b
	} else if len(b) == 0 {
		fmt.Println("MATRIX ZERO LEN B")
		fmt.Printf("\nRES: %v \n", a)
		fmt.Println("MATRIX OVER--------------------------------------")
		return a
	}

	res = make([][]interface{}, len(a)*len(b))

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(b); k++ {
				res[i] = append(res[i], a[j]...)
				res[i] = append(res[i], b[k]...)
			}
		}
	}
	fmt.Printf("\nRES: %v \n", res)
	fmt.Println("MATRIX OVER--------------------------------------")
	return res
}

func appendArrays(a, b [][]interface{}) [][]interface{} {
	var res [][]interface{}

	res = make([][]interface{}, 0)

	for i := 0; i < len(a); i++ {
		res = append(res, a[i])
	}
	for j := 0; j < len(b); j++ {
		res = append(res, b[j])
	}

	// for i := 0; i < len(res); i++ {
	// 	for j := 0; j < len(a); j++ {
	// 		for k := 0; k < len(b); k++ {
	// 			res[i] = append(res[i], a[j]...)
	// 			res[i] = append(res[i], b[k]...)
	// 		}
	// 	}
	// }
	fmt.Printf("\nRES: %v \n", res)
	fmt.Println("MATRIX OVER--------------------------------------")
	return res
}

// TODO
func Flatten(iface interface{}) [][]interface{} {
	var flatArray [][]interface{}
	concreteVal := reflect.ValueOf(iface)

	// fmt.Printf("TYPEOF is:%s\n", typeOf)
	// fmt.Printf("VALUEOF is:%s\n", valOf)
	// fmt.Printf("KINDOF is:%s\n", valOf.Kind())

	switch concreteVal.Kind() {
	case reflect.Ptr:
		return Flatten(concreteVal.Elem().Interface())
	case reflect.Struct:
		for i := 0; i < concreteVal.NumField(); i++ {
			f := concreteVal.Field(i)
			flatArray = appendMatrix(flatArray, Flatten(f.Interface()))
		}
	case reflect.Map:
		for _, k := range concreteVal.MapKeys() {
			v := concreteVal.MapIndex(k)
			flatArray = appendArrays(flatArray, Flatten(v.Interface()))
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < concreteVal.Len(); i++ {
			flatArray = appendMatrix(flatArray, Flatten(concreteVal.Index(i).Interface()))
			fmt.Println(concreteVal.Index(i))
		}
		fmt.Println("J'suis array / slice")
	case reflect.Invalid:
		fmt.Println("J'suis NIL")
	case reflect.UnsafePointer,
		reflect.Uintptr,
		reflect.Interface,
		reflect.Func,
		reflect.Chan:
		panic("Type not handled.")
	default:
		// fmt.Println("DEFAULT TEST")
		var tmpArray []interface{}
		tmpArray = append(tmpArray, concreteVal)
		flatArray = append(flatArray, tmpArray)
		// fmt.Println("DEFAULT OVER")
	}
	fmt.Printf("COUCOU mon gars !:%s\n", concreteVal.Kind())

	return flatArray
}

// func DeepFields(iface interface{}) []reflect.Value {
// fields := make([]reflect.Value, 0)
// ifv := reflect.ValueOf(iface)
// ift := reflect.TypeOf(iface)
//
// for i := 0; i < ift.NumField(); i++ {
// v := ifv.Field(i)
//
// switch v.Kind() {
// case reflect.Struct:
// fields = append(fields, DeepFields(v.Interface())...)
// default:
// fields = append(fields, v)
// }
// }
//
// return fields
// }
