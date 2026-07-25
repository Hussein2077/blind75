package arrayshashing

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [25]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++ // single quotes!
		count[t[i]-'a']--

	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
