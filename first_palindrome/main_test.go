package firstpalindrome

import "testing"

func TestFirstPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		out  string
	}{
		{name: "Example 1", in: []string{"abc", "car", "ada", "racecar", "cool"}, out: "ada"},
		{name: "Example 2", in: []string{"notapalindrome", "racecar"}, out: "racecar"},
		{name: "Example 3", in: []string{"def", "ghi"}, out: ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got, want := test.out, FirstPalindrome(test.in); got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
