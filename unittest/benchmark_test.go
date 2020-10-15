package unittest

import (
	"gopractise/server"
	"testing"
)

func BenchmarkAdd(t *testing.B) {
	for i := 0; i < t.N; i++ {
		a := 1
		b := 2
		server.Add(a, b)
	}
}

func BenchmarkSub(t *testing.B) {
	for i := 0; i < t.N; i++ {
		a := 2
		b := 1
		server.Add(a, b)
	}
}
