package add_two_numbers_test

import (
	"leetcode/add_two_numbers"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			l1 []int
			l2 []int
		}
		want []int
	}{
		{
			name: "Example 1",
			input: struct {
				l1 []int
				l2 []int
			}{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		}, {
			name: "Example 2",
			input: struct {
				l1 []int
				l2 []int
			}{
				l1: []int{0},
				l2: []int{0},
			},
			want: []int{0},
		}, {
			name: "Example 3",
			input: struct {
				l1 []int
				l2 []int
			}{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l1 := add_two_numbers.NewListNode(tt.input.l1...)
			l2 := add_two_numbers.NewListNode(tt.input.l2...)
			got := add_two_numbers.AddTwoNumbers(l1, l2).String()
			want := add_two_numbers.NewListNode(tt.want...).String()
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
