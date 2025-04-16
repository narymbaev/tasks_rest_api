package utils

func LongestSubstring(s string) string {
	var l, r, maxLen int
	stringLen := len(s)
	longestSubstring := ""
	if len(s) == 0 {
		return ""
	}
	mp := make(map[string]bool)
	runes := []rune(s)

	for r < stringLen {
		if _, ok := mp[string(runes[r])]; ok {
			delete(mp, string(runes[l]))
			l++
		} else {
			mp[string(runes[r])] = true
			currentLen := r - l + 1
			if currentLen > maxLen {
				longestSubstring = string(runes[l : r+1])
				maxLen = currentLen
			}
			r++
		}
	}

	return longestSubstring
}
