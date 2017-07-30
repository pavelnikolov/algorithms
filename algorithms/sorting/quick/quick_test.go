package quick_test

import (
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/quick"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestQuickSort(t *testing.T) {
	sorttest.Test(t, quick.Sort)
}

func BenchmarkQuickSort100(b *testing.B)    { sorttest.Benchmark(b, 100, quick.Sort) }
func BenchmarkQuickSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, quick.Sort) }
func BenchmarkQuickSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, quick.Sort) }
func BenchmarkQuickSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, quick.Sort) }
