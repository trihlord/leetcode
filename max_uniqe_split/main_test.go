package maxuniqesplit

import "testing"

func TestMaxUnqueSplit(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  int
	}{
		{name: "Example 1", in: "ababccc", out: 5},
		{name: "Exmaple 2", in: "aba", out: 2},
		{name: "Exmaple 3", in: "aa", out: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got, want := MaxUniqeSplit(test.in), test.out; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
