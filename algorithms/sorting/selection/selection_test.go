package selection_test

import (
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/selection"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestSelectionSort(t *testing.T) {
	sorttest.Test(t, selection.Sort)
}

func BenchmarkSelectionSort100(b *testing.B)    { sorttest.Benchmark(b, 100, selection.Sort) }
func BenchmarkSelectionSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, selection.Sort) }
func BenchmarkSelectionSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, selection.Sort) }
func BenchmarkSelectionSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, selection.Sort) }
