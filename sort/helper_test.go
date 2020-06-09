package sort

import (
	"math/rand"

	"github.com/moorara/algo/compare"
)

func sorted(items []interface{}, compare compare.Func) bool {
	for i := 0; i < len(items)-1; i++ {
		if compare(items[i], items[i+1]) > 0 {
			return false
		}
	}

	return true
}

func randIntSlice(size, min, max int) []interface{} {
	items := make([]interface{}, size)
	for i := 0; i < len(items); i++ {
		items[i] = min + rand.Intn(max-min+1)
	}

	return items
}
