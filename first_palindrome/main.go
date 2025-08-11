package first_palindrome

func FirstPalindrome(words []string) string {
	for _, word := range words {
		if isPolindrome(word) {
			return word
		}
	}
	return ""
}

func isPolindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
