package shell_test

import (
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/sorting/shell"
	"github.com/pavelnikolov/algorithms/algorithms/sorting/sorttest"
)

func TestShellSort(t *testing.T) {
	sorttest.Test(t, shell.Sort)
}

func BenchmarkShellSort100(b *testing.B)    { sorttest.Benchmark(b, 100, shell.Sort) }
func BenchmarkShellSort1000(b *testing.B)   { sorttest.Benchmark(b, 1000, shell.Sort) }
func BenchmarkShellSort10000(b *testing.B)  { sorttest.Benchmark(b, 10000, shell.Sort) }
func BenchmarkShellSort100000(b *testing.B) { sorttest.Benchmark(b, 100000, shell.Sort) }

func TestShellSort2(t *testing.T) {
	sorttest.Test(t, shell.Sort2)
}

func BenchmarkShellSort2_100(b *testing.B)    { sorttest.Benchmark(b, 100, shell.Sort2) }
func BenchmarkShellSort2_1000(b *testing.B)   { sorttest.Benchmark(b, 1000, shell.Sort2) }
func BenchmarkShellSort2_10000(b *testing.B)  { sorttest.Benchmark(b, 10000, shell.Sort2) }
func BenchmarkShellSort2_100000(b *testing.B) { sorttest.Benchmark(b, 100000, shell.Sort2) }
