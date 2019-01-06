package goecharts

import "regexp"

// reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func replaceJsFuncs(fn string) string {
	pat, _ := regexp.Compile(`\n|\t`)
	fn = pat.ReplaceAllString(fn, "")
	return "__x__" + fn + "__x__"
}
