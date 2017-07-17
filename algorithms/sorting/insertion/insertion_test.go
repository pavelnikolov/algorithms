package insertion_test

import (
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/insertion"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestInsertionSort(t *testing.T) {
	sorttest.Test(t, insertion.Sort)
}

func BenchmarkInsertionSort100(b *testing.B)    { sorttest.Benchmark(b, 100, insertion.Sort) }
func BenchmarkInsertionSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, insertion.Sort) }
func BenchmarkInsertionSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, insertion.Sort) }
func BenchmarkInsertionSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, insertion.Sort) }
