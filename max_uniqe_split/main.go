package maxuniqesplit

func MaxUniqeSplit(s string) int {
	ss := make(map[string]struct{}, len(s))
	return backtrace(s, ss)
}

func backtrace(s string, ss map[string]struct{}) int {
	if len(s) == 0 {
		return len(ss)
	}
	res := len(ss)
	for i := 0; i < len(s); i++ {
		if _, exists := ss[s[:i+1]]; !exists {
			ss[s[:i+1]] = struct{}{}
			res = max(res, backtrace(s[i+1:], ss))
			delete(ss, s[:i+1])
		}
	}
	return res
}
