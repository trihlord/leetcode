package lettercasepermutation

import (
	"slices"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  []string
	}{
		{
			name: "example 1",
			in:   "a1b2",
			out:  []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			name: "example 2",
			in:   "3z4",
			out:  []string{"3z4", "3Z4"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, want := LetterCasePermutation(test.in), test.out
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
