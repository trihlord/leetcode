package canreachcorner

import "testing"

func TestCanReachCorner(t *testing.T) {
	tests := []struct {
		name string
		in   struct {
			xCorner int
			yCorner int
			circles [][]int
		}
		out bool
	}{
		{
			name: "example 1",
			in: struct {
				xCorner int
				yCorner int
				circles [][]int
			}{xCorner: 3, yCorner: 4, circles: [][]int{{2, 1, 1}}},
			out: true,
		},
		{
			name: "example 2",
			in: struct {
				xCorner int
				yCorner int
				circles [][]int
			}{xCorner: 3, yCorner: 3, circles: [][]int{{1, 1, 2}}},
			out: false,
		},
		{
			name: "example 3",
			in: struct {
				xCorner int
				yCorner int
				circles [][]int
			}{xCorner: 3, yCorner: 3, circles: [][]int{{2, 1, 1}, {1, 2, 1}}},
			out: false,
		},
		{
			name: "example 4",
			in: struct {
				xCorner int
				yCorner int
				circles [][]int
			}{xCorner: 4, yCorner: 4, circles: [][]int{{5, 5, 1}}},
			out: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got, want := CanReachCorner(test.in.xCorner, test.in.yCorner, test.in.circles), test.out; got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
