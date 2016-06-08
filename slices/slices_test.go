package slices

import (
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

func TestStringInArray_doesContain(t *testing.T) {
	ensure.True(t, StringInArray([]string{"a", "b", "c", "d"}, "c"))
}

func TestStringInArray_doesNotContain(t *testing.T) {
	ensure.False(t, StringInArray([]string{"a", "b", "c", "d"}, "z"))
}

func TestStringInArray_emptySlice(t *testing.T) {
	ensure.False(t, StringInArray([]string{}, "a"))
}

func TestStringInArray_emptyString(t *testing.T) {
	ensure.False(t, StringInArray([]string{"a", "b", "c", "d"}, ""))
}

func TestStringInArray_nilSlice(t *testing.T) {
	ensure.False(t, StringInArray(nil, "a"))
}
