package sorting

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestBubbleSort(t *testing.T) {
	real := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := []int{1, 5, 4, 8, 6, 7, 9, 0, 3, 2}
	result := BubbleSort(test)

	assert.Equal(t, real, result)
}
