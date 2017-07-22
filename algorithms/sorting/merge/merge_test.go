package merge_test

import (
	"fmt"
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/merge"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestMergeSort(t *testing.T) {
	sorttest.Test(t, merge.Sort)
}

func BenchmarkMergeSort100(b *testing.B)    { sorttest.Benchmark(b, 100, merge.Sort) }
func BenchmarkMergeSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, merge.Sort) }
func BenchmarkMergeSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, merge.Sort) }
func BenchmarkMergeSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, merge.Sort) }

func TestMerge(t *testing.T) {
	data := []struct {
		name string
		a1   []int
		a2   []int
	}{
		{
			name: "one_to_ten",
			a1:   []int{1, 3, 5, 7, 9},
			a2:   []int{2, 4, 6, 8, 10},
		},
		{
			name: "a1_empty",
			a1:   []int{},
			a2:   []int{2, 4, 6, 8, 10},
		},
		{
			name: "a2_empty",
			a1:   []int{1, 3, 5, 7, 9},
			a2:   []int{},
		},
	}

	for _, d := range data {
		a := merge.Merge(d.a1, d.a2)
		for i := 0; i < len(a)-2; i++ {
			if a[i] > a[i+1] {
				fmt.Println(a)
				t.Errorf("a[%d](%d) > a[%d](%d), expected to be less than or equal to", i, a[i], i+1, a[i+1])
				t.FailNow()
			}
		}
	}
}
