package bitmap

import (
	"reflect"
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

// TODO
func generateSubset(iface interface{}) [][]interface{} {
	subset := [][]interface{}{
		{reflect.ValueOf(iface)},
	}
	return subset
}

// -----------------------------------------------------------------------------

func TestFlatten_singleValue_int(t *testing.T) {
	ensure.Subset(t, Flatten(int(50)), generateSubset(int(50)))
}

func TestFlatten_singleValue_int8(t *testing.T) {
	ensure.Subset(t, Flatten(int8(8)), generateSubset(int8(8)))
}

func TestFlatten_singleValue_int16(t *testing.T) {
	ensure.Subset(t, Flatten(int16(16)), generateSubset(int16(16)))
}

func TestFlatten_singleValue_int32(t *testing.T) {
	ensure.Subset(t, Flatten(int32(32)), generateSubset(int32(32)))
}

func TestFlatten_singleValue_int64(t *testing.T) {
	ensure.Subset(t, Flatten(int64(64)), generateSubset(int64(64)))
}

func TestFlatten_singleValue_uint(t *testing.T) {
	ensure.Subset(t, Flatten(uint(50)), generateSubset(uint(50)))
}

func TestFlatten_singleValue_uint8(t *testing.T) {
	ensure.Subset(t, Flatten(uint8(8)), generateSubset(uint8(8)))
}

func TestFlatten_singleValue_uint16(t *testing.T) {
	ensure.Subset(t, Flatten(uint16(16)), generateSubset(uint16(16)))
}

func TestFlatten_singleValue_uint32(t *testing.T) {
	ensure.Subset(t, Flatten(uint32(32)), generateSubset(uint32(32)))
}

func TestFlatten_singleValue_uint64(t *testing.T) {
	ensure.Subset(t, Flatten(uint64(64)), generateSubset(uint64(64)))
}

func TestFlatten_singleValue_float32(t *testing.T) {
	ensure.Subset(t, Flatten(float32(50.32)), generateSubset(float32(50.32)))
}

func TestFlatten_singleValue_float64(t *testing.T) {
	ensure.Subset(t, Flatten(float64(50.64)), generateSubset(float64(50.64)))
}

func TestFlatten_singleValue_complex64(t *testing.T) {
	ensure.Subset(t, Flatten(complex64(64)), generateSubset(complex64(64)))
}

func TestFlatten_singleValue_complex128(t *testing.T) {
	ensure.Subset(t, Flatten(complex128(128)), generateSubset(complex128(128)))
}

func TestFlatten_singleValue_ptr(t *testing.T) {
	ensure.Subset(t, Flatten(int(50)), generateSubset(int(50)))
}

// -----------------------------------------------------------------------------
