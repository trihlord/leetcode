package first_palindrome

func FirstPalindrome(words []string) string {
	for _, word := range words {
		rr := []rune(word)
		p := true
		for i, j := 0, len(rr)-1; i < j; i, j = i+1, j-1 {
			if rr[i] != rr[j] {
				p = false
				break
			}
		}
		if p {
			return word
		}
	}
	return ""
}
