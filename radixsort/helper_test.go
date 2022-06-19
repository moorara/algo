package radixsort

import "strings"

func isSortedString(vals []string) bool {
	for i := 0; i < len(vals)-1; i++ {
		if strings.Compare(vals[i], vals[i+1]) > 0 {
			return false
		}
	}
	return true
}

func isSortedInt(vals []int) bool {
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] > vals[i+1] {
			return false
		}
	}
	return true
}
