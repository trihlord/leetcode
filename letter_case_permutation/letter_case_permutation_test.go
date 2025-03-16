package letter_case_permutation_test

import (
	"leetcode/letter_case_permutation"
	"slices"
	"testing"
)

var tests = []struct {
	in  string
	out []string
}{
	{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
	{"3z4", []string{"3z4", "3Z4"}},
}

func TestLetterCasePermutation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			in := letter_case_permutation.LetterCasePermutation(tt.in)
			out := tt.out
			slices.Sort(in)
			slices.Sort(out)
			if !slices.Equal(in, out) {
				t.Errorf("got %q, want %q", in, out)
			}
		})
	}
}
