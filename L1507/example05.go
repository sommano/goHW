package example05

import "testing"

func benchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkStringFromAssignment(100)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}

func benchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}
