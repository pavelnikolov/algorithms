package sorttest

import (
	"fmt"
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/utils"
)

type SortFunc func([]int)

func Benchmark(b *testing.B, n int, fn SortFunc) {
	list := utils.GenerateArray(n)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(list)
	}
}

func Test(t *testing.T, fn SortFunc) {
	list := utils.GenerateArray(100)

	fn(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Errorf("list[%d](%d) > list[%d](%d), expected to be less than or equal to", i, list[i], i+1, list[i+1])
			t.FailNow()
		}
	}
}
