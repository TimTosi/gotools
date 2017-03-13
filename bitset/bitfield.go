package bitset

import (
	"fmt"
	"strconv"

	"github.com/timtosi/gotools/slices"
)

// -----------------------------------------------------------------------------

// PrintByte displays each bit value of the specific `byte` of `bf.data`
// at index `idx`.
func (bf *BitField) PrintByte(idx int64) {
	if idx < 0 || idx >= bf.bitCount/8 {
		fmt.Printf("Index out of range\n")
	}

	fmt.Printf(
		"%d%d%d%d%d%d%d%d",
		bf.data[idx]>>7&1,
		bf.data[idx]>>6&1,
		bf.data[idx]>>5&1,
		bf.data[idx]>>4&1,
		bf.data[idx]>>3&1,
		bf.data[idx]>>2&1,
		bf.data[idx]>>1&1,
		bf.data[idx]>>0&1,
	)
}

// Pos returns the value of the bit at index `idx` in `bf.data`.
func (bf *BitField) Pos(idx int64) int {
	if idx < 0 || idx >= bf.bitCount {
		return 0
	}
	return int(bf.data[idx/8]>>(7-uint(idx%8))) & 1
}

// WhichSet returns a slice of `[]int64` of up to `limit` values representing
// indexes of `bf.data` where a bit is set to `v`.
func (bf *BitField) WhichSet(v int, limit int64) (res []int64) {
	for i, j := int64(0), int64(0); i < bf.bitCount && j < limit; i++ {
		if int(bf.data[i/8]>>(7-uint(i%8))&1) == v {
			res = append(res, int64(i))
			j++
		}
	}
	return res
}

// WhichSetFrom returns a slice of `[]int64` of up to `limit` values representing
// indexes of `bf.data` where a bit is set to `v` and the index of the last bit
// processed.
func (bf *BitField) WhichSetFrom(v int, limit, from int64) (int64, []int64) {
	var res []int64

	for i, j := from, int64(0); i < bf.bitCount && j < limit; i++ {
		if int(bf.data[i/8]>>(7-uint(i%8))&1) == v {
			res = append(res, int64(i))
			j++
		}
	}
	return from, res
}

// WhichSetInclusive returns a slice of `[]int64` values representing indexes
// of `bf.data` where a bit is set to `1`. If `inclusive` is `true`, values
// found in `res` are composed of a subset of `idxs`. If `inclusive` is `false`,
// values found in `res` cannot be found in `idxs`.
func (bf *BitField) WhichSetInclusive(idxs []string, inclusive bool) (res []int64) {
	if inclusive == true {
		for _, idx := range idxs {
			i, err := strconv.ParseInt(idx, 10, 64)
			if err == nil && i < bf.bitCount && int(bf.data[i/8]>>(7-uint(i%8))&1) == 1 {
				res = append(res, i)
			}
		}
	} else {
		for i, j := int64(0), int64(0); i < bf.bitCount && j < 10; i++ {
			if int(bf.data[i/8]>>(7-uint(i%8))&1) == 1 &&
				slices.StringInArray(idxs, strconv.FormatInt(i, 10)) == false {
				res = append(res, int64(i))
				j++
			}
		}
	}
	return res
}

// BitCount returns the number of bit contained in bf.data.
func (bf *BitField) BitCount() int64 { return bf.bitCount }

// -----------------------------------------------------------------------------

// BitMatrix is a `slice` of `bitset.BitField`.
// type BitMatrix []*BitField

// BitField is a `struct` exposing methods related to bit manipulation.
type BitField struct {
	data     []byte
	bitCount int64
}

// NewBitField returns a new `BitField`.
func NewBitField(data []byte) *BitField { return &BitField{data: data, bitCount: int64(len(data) * 8)} }
