package maxuniqesplit

func MaxUniqeSplit(s string) int {
	ss := map[string]struct{}{}
	max := 0
	var backtrace func(index int, cur int)
	backtrace = func(index int, cur int) {
		if index == len(s) {
			if cur > max {
				max = cur
			}
			return
		}
		for i := index; i < len(s); i++ {
			sub := s[index : i+1]
			if _, exists := ss[sub]; !exists {
				ss[sub] = struct{}{}
				backtrace(i+1, cur+1)
				delete(ss, sub)
			}
		}
	}
	backtrace(0, 0)
	return max
}
