package bitset

import "fmt"

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

// BitCount returns the number of bit contained in bf.data.
func (bf *BitField) BitCount() int64 { return bf.bitCount }

// -----------------------------------------------------------------------------

// BitField is a `struct` exposing methods related to bit manipulation.
type BitField struct {
	data     []byte
	bitCount int64
}

// NewBitField returns a new `BitField`.
func NewBitField(data []byte) *BitField { return &BitField{data: data, bitCount: int64(len(data) * 8)} }
