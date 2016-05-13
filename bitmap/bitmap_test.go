package bitmap

import (
	"reflect"
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_int(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int(50))},
	}
	ensure.Subset(t, Flatten(int(50)), subset)
}

func TestFlatten_singleValue_int8(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int8(50))},
	}
	ensure.Subset(t, Flatten(int8(50)), subset)
}

func TestFlatten_singleValue_int16(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int16(50))},
	}
	ensure.Subset(t, Flatten(int16(50)), subset)
}

func TestFlatten_singleValue_int32(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int32(50))},
	}
	ensure.Subset(t, Flatten(int32(50)), subset)
}

func TestFlatten_singleValue_int64(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int64(50))},
	}
	ensure.Subset(t, Flatten(int64(50)), subset)
}

func TestFlatten_singleValue_uint(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(uint(50))},
	}
	ensure.Subset(t, Flatten(uint(50)), subset)
}

func TestFlatten_singleValue_uint8(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(uint8(50))},
	}
	ensure.Subset(t, Flatten(uint8(50)), subset)
}

func TestFlatten_singleValue_uint16(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(uint16(50))},
	}
	ensure.Subset(t, Flatten(uint16(50)), subset)
}

func TestFlatten_singleValue_uint32(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(uint32(50))},
	}
	ensure.Subset(t, Flatten(uint32(50)), subset)
}

func TestFlatten_singleValue_uint64(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(uint64(50))},
	}
	ensure.Subset(t, Flatten(uint64(50)), subset)
}

func TestFlatten_singleValue_float32(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(float32(50.20))},
	}
	ensure.Subset(t, Flatten(float32(50.20)), subset)
}

func TestFlatten_singleValue_float64(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(float64(50.20))},
	}
	ensure.Subset(t, Flatten(float64(50.20)), subset)
}

func TestFlatten_singleValue_complex64(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(complex64(50))},
	}
	ensure.Subset(t, Flatten(complex64(50)), subset)
}

func TestFlatten_singleValue_complex128(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(complex128(50))},
	}
	ensure.Subset(t, Flatten(complex128(50)), subset)
}

func TestFlatten_singleValue_ptr(t *testing.T) {
	subset := [][]interface{}{
		{reflect.ValueOf(int(50))},
	}

	ensure.Subset(t, Flatten(int(50)), subset)
}

// -----------------------------------------------------------------------------
