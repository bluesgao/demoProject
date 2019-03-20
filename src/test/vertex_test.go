package test

import (
	"petmate/src/vertex"
	"testing"
)

func TestVertexAbs(t *testing.T) {
	v := vertex.Vertex{3, 4}
	ret := v.Abs()

	if ret != 5 {
		t.FailNow()
	}
}

func BenchmarkAbs(b *testing.B) {
	v := vertex.Vertex{3, 4}
	for i := 0; i < b.N; i++ {
		_ = v.Abs()
	}
}
