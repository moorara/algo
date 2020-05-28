package radixsort

import "strings"

func isSortedString(items []string) bool {
	for i := 0; i < len(items)-1; i++ {
		if strings.Compare(items[i], items[i+1]) > 0 {
			return false
		}
	}
	return true
}
