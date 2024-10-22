package workerpool

import (
	"fmt"
	"testing"
)

func TestJobCreator(t *testing.T) {
	inputs := []struct {
		num      int
		expected int
	}{
		{
			num:      8433,
			expected: 18,
		},
		{
			num:      5670,
			expected: 18,
		},
	}

	for _, input := range inputs {
		fmt.Printf("res is=%v\n", sumdigits(input.num))
	}
}
