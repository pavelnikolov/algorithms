package bubble_test

import (
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/bubble"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestBubble(t *testing.T) {
	sorttest.Test(t, bubble.Sort)
}

func BenchmarkBubbleSort100(b *testing.B)    { sorttest.Benchmark(b, 100, bubble.Sort) }
func BenchmarkBubbleSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, bubble.Sort) }
func BenchmarkBubbleSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, bubble.Sort) }
func BenchmarkBubbleSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, bubble.Sort) }
