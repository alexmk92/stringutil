package stringutil

import "strings"

func CaseInsenstiveContains(src, quantifier string) bool {
	return strings.Contains(strings.ToUpper(src), strings.ToUpper(quantifier))
}

func CaseInsensitiveIndexOf(src,  quantifier string) int {
	index := -1
	b := CaseInsenstiveContains(src, quantifier)
	if(b) {
		index = strings.Index(strings.ToUpper(src), strings.ToUpper(quantifier))
	}
	return index
}