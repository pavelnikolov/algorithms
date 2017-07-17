package utils_test

import (
	"fmt"
	"testing"

	"github.com/pavelnikolov/algorithms/algorithms/utils"
)

func TestGenerateArray(t *testing.T) {
	for i := 0; i < 100000; i++ {
		if i%300 == 0 {
			array := utils.GenerateArray(i)

			if len(array) != i {
				fmt.Println(array)
				t.Error()
			}
		}
	}
}
