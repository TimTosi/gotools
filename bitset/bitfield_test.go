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
	testCases := []struct {
		name     string
		mockData string
		bitValue int
		limit    int64
		expected []int64
	}{
		{"bitSetRegular", "Q", 1, 8, []int64{1, 3, 7}},
		{"bitSetEmptyBitField", "", 1, 8, ([]int64)(nil)},
		{"bitNotSetRegular", "Q", 0, 8, []int64{0, 2, 4, 5, 6}},
		{"bitNotSetEmptyBitField", "", 0, 8, ([]int64)(nil)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ensure.Subset(
				t,
				NewBitField([]byte(tc.mockData)).WhichSet(tc.bitValue, tc.limit),
				tc.expected,
			)
		})
	}
}

func TestBitField_WhichSetInclusive(t *testing.T) {
	testCases := []struct {
		name     string
		mockData string
		mockIdxs []string
		mockMod  bool
		expected []int64
	}{
		{"inclusiveRegular", "Q", []string{"1", "5", "7"}, true, []int64{1, 7}},
		{"inclusiveEmptyBitField", "", []string{"1", "5", "7"}, true, ([]int64)(nil)},
		{"inclusiveEmptyIdxs", "Q", []string{}, true, ([]int64)(nil)},
		{"inclusiveNoIdxsMatch", "Q", []string{"0", "2", "5"}, true, ([]int64)(nil)},
		{"exclusiveRegular", "Q", []string{"1", "7"}, false, []int64{3}},
		{"exclusiveEmptyBitField", "", []string{"1", "7"}, false, ([]int64)(nil)},
		{"exclusiveEmptyIdxs", "Q", []string{}, false, []int64{1, 3, 7}},
		{"exclusiveNoIdxsMatch", "Q", []string{"0", "2", "5"}, false, []int64{1, 3, 7}},
		{"exclusiveAllIdxsMatch", "Q", []string{"1", "3", "7"}, false, ([]int64)(nil)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ensure.Subset(
				t,
				NewBitField([]byte(tc.mockData)).WhichSetInclusive(tc.mockIdxs, tc.mockMod),
				tc.expected,
			)
		})
	}
}

// -----------------------------------------------------------------------------

func BenchmarkWhichSet(b *testing.B) {
	benchmarks := []struct {
		name     string
		limit    int
		expected int64
	}{
		{"100Bytes", 100, 300},
		{"1000Bytes", 1000, 3000},
		{"10000Bytes", 10000, 30000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			bf := NewBitField([]byte(mockData(bm.limit)))
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				_ = bf.WhichSet(1, bm.expected)
			}
		})
	}
}
