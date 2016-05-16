package bitmap

import (
	"reflect"
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

// generateSubset returns a `[][]interface{}` composed of `iface` values.
func generateSubsetSingle(iface interface{}) [][]interface{} {
	subset := [][]interface{}{
		{reflect.ValueOf(iface)},
	}
	return subset
}

// generateSubsetMultiple returns a `[][]interface{}` composed of
// `iface` values.
func generateSubsetMultiple(iface interface{}) [][]interface{} {
	var subset [][]interface{}
	concreteVal := reflect.ValueOf(iface)

	if concreteVal.Kind() == reflect.Map {
		for i, k := range concreteVal.MapKeys() {
			v := concreteVal.MapIndex(k)
			subset = append(subset, []interface{}{})
			subset[i] = append(subset[i], v)
		}
	} else {
		for i := 0; i < concreteVal.Len(); i++ {
			subset = append(subset, []interface{}{})
			subset[i] = append(subset[i], concreteVal.Index(i))
		}
	}
	return subset
}

// generateSubsetMultiple returns a `[][]interface{}` composed of
// `iface` values.
func generateSubsetStruct(iface interface{}) [][]interface{} {
	var subset [][]interface{}
	concreteVal := reflect.ValueOf(iface)
	subset = append(subset, make([]interface{}, 0))

	for i := 0; i < concreteVal.NumField(); i++ {
		if concreteVal.Field(i).Kind() == reflect.Struct {
			subset = concatArrays(
				subset,
				generateSubsetStruct(concreteVal.Field(i).Interface()),
			)
		} else {
			subset[0] = append(subset[0], concreteVal.Field(i))
		}
	}
	return subset
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_int(t *testing.T) {
	ensure.Subset(t, Flatten(int(50)), generateSubsetSingle(int(50)))
}

func TestFlatten_singleValue_int8(t *testing.T) {
	ensure.Subset(t, Flatten(int8(8)), generateSubsetSingle(int8(8)))
}

func TestFlatten_singleValue_int16(t *testing.T) {
	ensure.Subset(t, Flatten(int16(16)), generateSubsetSingle(int16(16)))
}

func TestFlatten_singleValue_int32(t *testing.T) {
	ensure.Subset(t, Flatten(int32(32)), generateSubsetSingle(int32(32)))
}

func TestFlatten_singleValue_int64(t *testing.T) {
	ensure.Subset(t, Flatten(int64(64)), generateSubsetSingle(int64(64)))
}

func TestFlatten_singleValue_uint(t *testing.T) {
	ensure.Subset(t, Flatten(uint(50)), generateSubsetSingle(uint(50)))
}

func TestFlatten_singleValue_uint8(t *testing.T) {
	ensure.Subset(t, Flatten(uint8(8)), generateSubsetSingle(uint8(8)))
}

func TestFlatten_singleValue_uint16(t *testing.T) {
	ensure.Subset(t, Flatten(uint16(16)), generateSubsetSingle(uint16(16)))
}

func TestFlatten_singleValue_uint32(t *testing.T) {
	ensure.Subset(t, Flatten(uint32(32)), generateSubsetSingle(uint32(32)))
}

func TestFlatten_singleValue_uint64(t *testing.T) {
	ensure.Subset(t, Flatten(uint64(64)), generateSubsetSingle(uint64(64)))
}

func TestFlatten_singleValue_float32(t *testing.T) {
	ensure.Subset(t, Flatten(float32(50.32)), generateSubsetSingle(float32(50.32)))
}

func TestFlatten_singleValue_float64(t *testing.T) {
	ensure.Subset(t, Flatten(float64(50.64)), generateSubsetSingle(float64(50.64)))
}

func TestFlatten_singleValue_complex64(t *testing.T) {
	ensure.Subset(t, Flatten(complex64(64)), generateSubsetSingle(complex64(64)))
}

func TestFlatten_singleValue_complex128(t *testing.T) {
	ensure.Subset(t, Flatten(complex128(128)), generateSubsetSingle(complex128(128)))
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_ptr(t *testing.T) {
	a, b := 5, 5
	ensure.Subset(t, Flatten(&a), generateSubsetSingle(b))
}

func TestFlatten_singleValue_uintptr(t *testing.T) {
	a := 5
	ensure.Subset(t, Flatten(uintptr(a)), generateSubsetSingle(uintptr(a)))
}

// func TestFlatten_singleValue_unsafePtr(t *testing.T) {
// }

func TestFlatten_singleValue_nil(t *testing.T) {
	ensure.Subset(t, Flatten(nil), generateSubsetSingle(nil))
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_bool(t *testing.T) {
	ensure.Subset(t, Flatten(true), generateSubsetSingle(true))
}

func TestFlatten_singleValue_byte(t *testing.T) {
	ensure.Subset(t, Flatten(byte(5)), generateSubsetSingle(byte(5)))
}

func TestFlatten_singleValue_rune(t *testing.T) {
	ensure.Subset(t, Flatten(rune('⌘')), generateSubsetSingle(rune('⌘')))
}

func TestFlatten_singleValue_string(t *testing.T) {
	ensure.Subset(t, Flatten("Heya"), generateSubsetSingle("Heya"))
}

// func TestFlatten_singleValue_slice(t *testing.T) {
// 	ensure.Subset(t, Flatten([]int{2, 3}), generateSubsetMultiple([]int{2, 3}))
// }

func TestFlatten_singleValue_array(t *testing.T) {
	ensure.Subset(t, Flatten([2]int{2, 3}), generateSubsetMultiple([2]int{2, 3}))
}

func TestFlatten_singleValue_map(t *testing.T) {
	ensure.Subset(t,
		Flatten(map[int]string{1: "first", 2: "second"}),
		generateSubsetMultiple(map[int]string{1: "first", 2: "second"}),
	)
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_func(t *testing.T) {
	myFunc := func() { print("Yo") }
	ensure.Subset(t, Flatten(myFunc), generateSubsetSingle(myFunc))
}

func TestFlatten_singleValue_chan(t *testing.T) {
	myChan := make(chan int)
	ensure.Subset(t, Flatten(myChan), generateSubsetSingle(myChan))
}

// -----------------------------------------------------------------------------

func TestFlatten_composedValue_singleStruct(t *testing.T) {
	mockStruct := struct {
		MockInt    int
		MockString string
	}{1, "Ok"}
	ensure.Subset(t, Flatten(mockStruct), generateSubsetStruct(mockStruct))
}

func TestFlatten_composedValue_nestedStruct(t *testing.T) {
	mockStruct := struct {
		MockInt    int
		MockString string
		MockStruct struct{ MockInt32 int32 }
	}{
		1,
		"Ok",
		struct{ MockInt32 int32 }{32},
	}
	ensure.Subset(t, Flatten(mockStruct), generateSubsetStruct(mockStruct))
}
