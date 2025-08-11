package first_palindrome

func FirstPalindrome(words []string) string {
	for _, word := range words {
		p := true
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			if word[i] != word[j] {
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
