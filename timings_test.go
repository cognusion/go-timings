package timings

import (
	"testing"
)

func Benchmark_Tracker(b *testing.B) {
	var t Tracker
	for i := 0; i < b.N; i++ {
		t = Tracker{}
		t.Start()
	}
}
