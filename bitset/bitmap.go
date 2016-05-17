package bitset

import "reflect"

// -----------------------------------------------------------------------------

// Bitset is a structure representing a bitmap set.
type Bitset map[string]map[string][]int

// -----------------------------------------------------------------------------

// appendMatrix is a helper function used by `Flatten`.
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

// concatArrays is a helper function used by `Flatten`.
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

// -----------------------------------------------------------------------------

// ValToBitSet generates a bitmap set from a two-dimensional array.
// Each column name is stored in `index`.
//
// NOTE: `index` lenght MUST be equal or longer than slice's length contained
// in `valArrays`.
func ValToBitSet(valArrays [][]interface{}, index []string) (bs Bitset) {
	bs = make(Bitset)

	for i, valArray := range valArrays {
		for j, val := range valArray {
			colName := index[j]
			v := val.(reflect.Value).Interface().(string)
			if _, ok := bs[colName]; !ok {
				bs[colName] = make(map[string][]int)
			}
			if _, ok := bs[colName][v]; !ok {
				bs[colName][v] = make([]int, len(valArrays))
			}
			bs[colName][v][i] = 1
		}
	}
	return
}

// Flatten generates a two-dimensional array from several Go type, including
// composed types and nested types. The purpose of this array is to be
// converted into a bitmap set.
//
// NOTE: `struct` and `map` types MUST always be the last field of `iface`.
// NOTE: This function panics if a struct's field is not exported.
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
