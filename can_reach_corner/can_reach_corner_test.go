package can_reach_corner_test

import (
	"leetcode/can_reach_corner"
	"testing"
)

func TestCanReachCorner(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		xCorner int
		yCorner int
		circles [][]int
		output  bool
	}{}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if expected, actual := test.output, can_reach_corner.CanReachCorner(test.xCorner, test.yCorner, test.circles); expected != actual {
				t.Errorf("wand %t, got %t", expected, actual)
			}
		})
	}
}
