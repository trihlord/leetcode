package first_palindrome_test

import (
	"leetcode/first_palindrome"
	"testing"
)

func TestFirstPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output string
	}{
		{name: "Example 1", input: []string{"abc", "car", "ada", "racecar", "cool"}, output: "ada"},
		{name: "Example 2", input: []string{"notapalindrome", "racecar"}, output: "racecar"},
		{name: "Example 3", input: []string{"def", "ghi"}, output: ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got, want := test.output, first_palindrome.FirstPalindrome(test.input); got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
