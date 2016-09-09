package bitset

import (
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

func mockBinary() []int {
	return []int{
		0, 1, 0, 1, 0, 1, 0, 0,
		0, 1, 1, 0, 0, 1, 0, 1,
		0, 1, 1, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0,
	}
}

func mockData(size int) string {
	var res string
	for i := 0; i < size; i++ {
		res += "a"
	}
	return res
}

// -----------------------------------------------------------------------------

func TestBitField_NewBitField(t *testing.T) {
	bf := NewBitField([]byte("ThisIsA23CharLongString"))
	ensure.NotNil(t, bf)
	ensure.DeepEqual(t, bf.bitCount, int64(23*8))
}

func TestBitField_BitCount(t *testing.T) {
	ensure.DeepEqual(
		t,
		NewBitField([]byte("ThisIsA23CharLongString")).BitCount(),
		int64(23*8),
	)
}

func TestBitField_Pos(t *testing.T) {
	bf := NewBitField([]byte("Test@"))
	for i, v := range mockBinary() {
		ensure.DeepEqual(t, bf.Pos(int64(i)), v)
	}
}

func TestBitField_WhichSet(t *testing.T) {
	bf := NewBitField([]byte("Q"))
	ensure.Subset(t, bf.WhichSet(1, 8), []int64{1, 3, 7})
	ensure.Subset(t, bf.WhichSet(0, 8), []int64{0, 2, 4, 5, 6})
}

// -----------------------------------------------------------------------------

func BenchmarkBitShift_100(b *testing.B) {
	limit := 100
	bf := NewBitField([]byte(mockData(limit)))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = bf.WhichSet(1, int64(3*limit))
	}
}

func BenchmarkBitShift_1000(b *testing.B) {
	limit := 1000
	bf := NewBitField([]byte(mockData(limit)))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = bf.WhichSet(1, int64(3*limit))
	}
}

func BenchmarkBitShift_10000(b *testing.B) {
	limit := 10000
	bf := NewBitField([]byte(mockData(limit)))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = bf.WhichSet(1, int64(3*limit))
	}
}
