package captcha

import "testing"

func BenchmarkNBitNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nBitNum(i % 6)
	}
}

func TestNBitBun(t *testing.T) {
	for i := 1; i < 8; i++ {
		t.Log(nBitNum(i))
	}
}
