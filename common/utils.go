package common

import "regexp"

// reverse string slice
func ReverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// replace and clear up js functions string
func ReplaceJsFuncs(fn string) string {
	pat, _ := regexp.Compile(`\n|\t`)
	fn = pat.ReplaceAllString(fn, "")
	return "__x__" + fn + "__x__"
}
