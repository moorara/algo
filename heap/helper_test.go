package heap

import "math/rand"

const (
	minLen = 10
	maxLen = 100
	chars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randIntSlice(size, min, max int) []int {
	vals := make([]int, size)
	for i := 0; i < len(vals); i++ {
		vals[i] = min + rand.Intn(max-min+1)
	}

	return vals
}

func randStringKey(minLen, maxLen int) string {
	n := len(chars)
	l := minLen + rand.Intn(maxLen-minLen+1)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[rand.Intn(n)]
	}

	return string(b)
}

func randStringSlice(size int) []string {
	s := make([]string, size)
	for i := range s {
		s[i] = randStringKey(minLen, maxLen)
	}

	return s
}
