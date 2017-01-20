package stringutil

import "strings"

func CaseInsenstiveContains(src string, quantifiers ...string) bool {
	found := false

	for _, quantifier := range quantifiers {
		if found == false {
			found = strings.Contains(strings.ToUpper(src), strings.ToUpper(quantifier))
		}
	}

	return found
}

func CaseInsensitiveIndexOf(src,  quantifier string) int {
	index := -1
	b := CaseInsenstiveContains(src, quantifier)
	if(b) {
		index = strings.Index(strings.ToUpper(src), strings.ToUpper(quantifier))
	}
	return index
}

func CaseInsensitiveSliceContainsString(src []string, quantifier string) bool {
	for _, v := range src {
		if v == quantifier {
			return true
		}
	}

	return false
}