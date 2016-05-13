package bitmap

import "reflect"

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
	tmp := a[0]

	for i := 0; i < len(b); i++ {
		if i == len(a) {
			a = append(a, tmp)
		}
		a[i] = append(a[i], b[i]...)
	}
	return a
}

func concatArrays(a, b [][]interface{}) [][]interface{} {
	var tmp []interface{}

	for i := 0; i < len(a); i++ {
		tmp = append(tmp, a[i]...)
	}
	for j := 0; j < len(b); j++ {
		tmp = append(tmp, b[j]...)
	}
	return append([][]interface{}{}, tmp)
}

// TODO
func Flatten(iface interface{}) [][]interface{} {
	var flatArray [][]interface{}
	concreteVal := reflect.ValueOf(iface)

	switch concreteVal.Kind() {
	case reflect.Ptr:
		return Flatten(concreteVal.Elem().Interface())
	case reflect.Struct:
		for i := 0; i < concreteVal.NumField(); i++ {
			f := concreteVal.Field(i)
			if f.Kind() == reflect.Map {
				flatArray = appendMatrix(flatArray, Flatten(f.Interface()))
			} else {
				flatArray = concatArrays(flatArray, Flatten(f.Interface()))
			}
		}
	case reflect.Map:
		for _, k := range concreteVal.MapKeys() {
			v := concreteVal.MapIndex(k)
			flatArray = append(flatArray, Flatten(v.Interface())...)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < concreteVal.Len(); i++ {
			flatArray = append(flatArray, Flatten(concreteVal.Index(i).Interface())...)
		}
	default:
		return [][]interface{}{{concreteVal}}
	}
	return flatArray
}
