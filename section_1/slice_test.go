package section1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEvenNumberInSlice(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "Empty slice", nums: []int{}, want: 0},
		{name: "Slice with no even numbers", nums: []int{1, 3, 5}, want: 0},
		{name: "Slice with all even numbers", nums: []int{2, 4, 6}, want: 12},
		{name: "Slice with mixed even and odd numbers", nums: []int{1, 2, 3, 4}, want: 6},
		{name: "Slice with negative even numbers", nums: []int{-2, 1, 4, -8}, want: -6},
		{name: "Large slice", nums: make([]int, 1000), want: 0},
		{name: "Slice with duplicates", nums: []int{2, 4, 2, 6}, want: 14},
		{name: "Slice with very large even numbers", nums: []int{2147483644, 1, 2, 4}, want: 2147483650},
		{name: "Nil slice", nums: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumEvenNumberInSlice(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
