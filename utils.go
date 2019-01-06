package goecharts

import "regexp"

// reverse string slice
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	pat, _ := regexp.Compile(`\n|\t`)
	fn = pat.ReplaceAllString(fn, "")
	return "__x__" + fn + "__x__"
}

// switch chartType options and return itself struct
func switchChartOpts(options ...interface{}) (bool, interface{}) {
	for i := 0; i < len(options); i++ {
		switch options[i].(type) {
		case BarChartOpts:
			return true, options[i].(BarChartOpts)
		case LineChartOpts:
			return true, options[i].(LineChartOpts)
		}
	}
	return false, nil
}
