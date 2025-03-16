package letter_case_permutation

import (
	"slices"
	"unicode"
)

func LetterCasePermutation(s string) []string {
	ss := []string{s}
	for i := range len(s) {
		for _, ssS := range ss[:] {
			rr := []rune(ssS)
			rr[i] = unicode.ToUpper(rr[i])
			if s := string(rr); !slices.Contains(ss, s) {
				ss = append(ss, s)
			}
			rr[i] = unicode.ToLower(rr[i])
			if s := string(rr); !slices.Contains(ss, s) {
				ss = append(ss, s)
			}
		}
	}
	return ss
}
