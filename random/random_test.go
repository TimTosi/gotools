package random

import "testing"

// -----------------------------------------------------------------------------

func BenchmarkGenerateBytes_10(b *testing.B) {
	rg := NewRandomGenerator(888)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rg.GenerateBytes(10)
	}
}
