package bitset

import (
	"reflect"
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

var mockStructSingle = struct {
	MockInt    int
	MockString string
}{1, "Ok"}

var mockStructNested = struct {
	MockInt    int
	MockString string
	MockStruct struct{ MockInt32 int32 }
}{
	1,
	"Ok",
	struct{ MockInt32 int32 }{32},
}

// -----------------------------------------------------------------------------

// genSubset returns a `[][]interface{}` composed of `iface` values.
func genSubsetSingle(iface interface{}) [][]interface{} {
	subset := [][]interface{}{
		{reflect.ValueOf(iface)},
	}
	return subset
}

// genSubsetMultiple returns a `[][]interface{}` composed of
// `iface` values.
func genSubsetMultiple(iface interface{}) [][]interface{} {
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

// genSubsetMultiple returns a `[][]interface{}` composed of
// `iface` values.
func genSubsetStruct(iface interface{}) [][]interface{} {
	var subset [][]interface{}
	concreteVal := reflect.ValueOf(iface)
	subset = append(subset, make([]interface{}, 0))

	for i := 0; i < concreteVal.NumField(); i++ {
		if concreteVal.Field(i).Kind() == reflect.Struct {
			subset = concatArrays(
				subset,
				genSubsetStruct(concreteVal.Field(i).Interface()),
			)
		} else {
			subset[0] = append(subset[0], concreteVal.Field(i))
		}
	}
	return subset
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_int(t *testing.T) {
	ensure.Subset(t, Flatten(int(50)), genSubsetSingle(int(50)))
}

func TestFlatten_singleValue_int8(t *testing.T) {
	ensure.Subset(t, Flatten(int8(8)), genSubsetSingle(int8(8)))
}

func TestFlatten_singleValue_int16(t *testing.T) {
	ensure.Subset(t, Flatten(int16(16)), genSubsetSingle(int16(16)))
}

func TestFlatten_singleValue_int32(t *testing.T) {
	ensure.Subset(t, Flatten(int32(32)), genSubsetSingle(int32(32)))
}

func TestFlatten_singleValue_int64(t *testing.T) {
	ensure.Subset(t, Flatten(int64(64)), genSubsetSingle(int64(64)))
}

func TestFlatten_singleValue_uint(t *testing.T) {
	ensure.Subset(t, Flatten(uint(50)), genSubsetSingle(uint(50)))
}

func TestFlatten_singleValue_uint8(t *testing.T) {
	ensure.Subset(t, Flatten(uint8(8)), genSubsetSingle(uint8(8)))
}

func TestFlatten_singleValue_uint16(t *testing.T) {
	ensure.Subset(t, Flatten(uint16(16)), genSubsetSingle(uint16(16)))
}

func TestFlatten_singleValue_uint32(t *testing.T) {
	ensure.Subset(t, Flatten(uint32(32)), genSubsetSingle(uint32(32)))
}

func TestFlatten_singleValue_uint64(t *testing.T) {
	ensure.Subset(t, Flatten(uint64(64)), genSubsetSingle(uint64(64)))
}

func TestFlatten_singleValue_float32(t *testing.T) {
	ensure.Subset(t, Flatten(float32(50.32)), genSubsetSingle(float32(50.32)))
}

func TestFlatten_singleValue_float64(t *testing.T) {
	ensure.Subset(t, Flatten(float64(50.64)), genSubsetSingle(float64(50.64)))
}

func TestFlatten_singleValue_complex64(t *testing.T) {
	ensure.Subset(t, Flatten(complex64(64)), genSubsetSingle(complex64(64)))
}

func TestFlatten_singleValue_complex128(t *testing.T) {
	ensure.Subset(t, Flatten(complex128(128)), genSubsetSingle(complex128(128)))
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_ptr(t *testing.T) {
	a, b := 5, 5
	ensure.Subset(t, Flatten(&a), genSubsetSingle(b))
}

func TestFlatten_singleValue_uintptr(t *testing.T) {
	a := 5
	ensure.Subset(t, Flatten(uintptr(a)), genSubsetSingle(uintptr(a)))
}

// func TestFlatten_singleValue_unsafePtr(t *testing.T) {
// }

func TestFlatten_singleValue_nil(t *testing.T) {
	ensure.Subset(t, Flatten(nil), genSubsetSingle(nil))
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_bool(t *testing.T) {
	ensure.Subset(t, Flatten(true), genSubsetSingle(true))
}

func TestFlatten_singleValue_byte(t *testing.T) {
	ensure.Subset(t, Flatten(byte(5)), genSubsetSingle(byte(5)))
}

func TestFlatten_singleValue_rune(t *testing.T) {
	ensure.Subset(t, Flatten(rune('⌘')), genSubsetSingle(rune('⌘')))
}

func TestFlatten_singleValue_string(t *testing.T) {
	ensure.Subset(t, Flatten("Heya"), genSubsetSingle("Heya"))
}

// func TestFlatten_singleValue_slice(t *testing.T) {
// 	ensure.Subset(t, Flatten([]int{2, 3}), genSubsetMultiple([]int{2, 3}))
// }

func TestFlatten_singleValue_array(t *testing.T) {
	ensure.Subset(t, Flatten([2]int{2, 3}), genSubsetMultiple([2]int{2, 3}))
}

func TestFlatten_singleValue_map(t *testing.T) {
	ensure.Subset(t,
		Flatten(map[int]string{1: "first", 2: "second"}),
		genSubsetMultiple(map[int]string{1: "first", 2: "second"}),
	)
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_func(t *testing.T) {
	myFunc := func() { print("Yo") }
	ensure.Subset(t, Flatten(myFunc), genSubsetSingle(myFunc))
}

func TestFlatten_singleValue_chan(t *testing.T) {
	myChan := make(chan int)
	ensure.Subset(t, Flatten(myChan), genSubsetSingle(myChan))
}

// -----------------------------------------------------------------------------

func TestFlatten_composedValue_singleStruct(t *testing.T) {
	ensure.Subset(t, Flatten(mockStructSingle), genSubsetStruct(mockStructSingle))
}

func TestFlatten_composedValue_nestedStruct(t *testing.T) {
	ensure.Subset(t, Flatten(mockStructNested), genSubsetStruct(mockStructNested))
}

// -----------------------------------------------------------------------------

func ValToBitSetMap(t *testing.T) {
	mockBitSet := Bitset{
		"ID": map[string][]int{
			"1": {1, 0, 0, 0, 0, 0, 0, 0},
			"2": {0, 1, 0, 0, 0, 0, 0, 0},
			"3": {0, 0, 1, 0, 0, 0, 0, 0},
			"4": {0, 0, 0, 1, 0, 0, 0, 0},
			"5": {0, 0, 0, 0, 1, 0, 0, 0},
			"6": {0, 0, 0, 0, 0, 1, 0, 0},
			"7": {0, 0, 0, 0, 0, 0, 1, 0},
			"8": {0, 0, 0, 0, 0, 0, 0, 1},
		},
		"Device": map[string][]int{
			"Phone":     {1, 0, 1, 1, 0, 0, 0, 0},
			"Tablet":    {0, 1, 0, 0, 0, 0, 0, 0},
			"Desktop":   {0, 0, 0, 0, 1, 1, 0, 1},
			"SmartGear": {0, 0, 0, 0, 0, 0, 1, 0},
		},
		"Height": map[string][]int{
			"200": {1, 0, 0, 0, 0, 1, 0, 1},
			"300": {0, 0, 1, 0, 0, 0, 1, 0},
			"500": {0, 1, 0, 1, 1, 0, 0, 0},
		},
		"Length": map[string][]int{
			"200": {0, 0, 0, 0, 0, 0, 1, 0},
			"300": {0, 0, 1, 1, 1, 1, 0, 1},
			"500": {1, 1, 0, 0, 0, 0, 0, 0},
		},
	}
	idx := []string{"ID", "Device", "Height", "Length"}
	valArray := [][]interface{}{
		{"1", "Phone", "200", "500"},
		{"2", "Tablet", "500", "500"},
		{"3", "Phone", "300", "300"},
		{"4", "Phone", "500", "300"},
		{"5", "Desktop", "500", "300"},
		{"6", "Desktop", "200", "300"},
		{"7", "SmartGear", "300", "200"},
		{"8", "Desktop", "200", "300"},
	}
	ensure.Subset(t, ValToBitSet(valArray, idx), mockBitSet)
}
