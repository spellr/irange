package irange

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		params   []int
		expected []int
	}{
		{
			"No params",
			[]int{},
			[]int{},
		},
		{
			"One param",
			[]int{4},
			[]int{0, 1, 2, 3},
		},
		{
			"One negative param",
			[]int{-4},
			[]int{},
		},
		{
			"Two params",
			[]int{3, 6},
			[]int{3, 4, 5},
		},
		{
			"Three params",
			[]int{3, 9, 2},
			[]int{3, 5, 7},
		},
		{
			"Three params II",
			[]int{3, 8, 2},
			[]int{3, 5, 7},
		},
		{
			"Three params negative",
			[]int{8, -3, -2},
			[]int{8, 6, 4, 2, 0, -2},
		},
		{
			"Three params negative",
			[]int{8, -4, -2},
			[]int{8, 6, 4, 2, 0, -2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Range(tt.params...)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf(`%v: Expecting %v`, got, tt.expected)
			}
		})
	}
}
