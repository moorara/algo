package symboltable

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

const (
	minLen = 10
	maxLen = 100
	chars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randIntSlice(size int) []int {
	vals := make([]int, size)
	for i := 0; i < len(vals); i++ {
		vals[i] = r.Int()
	}

	return vals
}

func randStringKey(minLen, maxLen int) string {
	n := len(chars)
	l := minLen + r.Intn(maxLen-minLen+1)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[r.Intn(n)]
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

func runPutBenchmark(b *testing.B, st SymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		st.Put(keys[n], vals[n])
	}
}

func runGetBenchmark(b *testing.B, st SymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		st.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		st.Get(keys[n])
	}
}

func runDeleteBenchmark(b *testing.B, st SymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		st.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		st.Delete(keys[n])
	}
}

func BenchmarkSymbolTable_Put(b *testing.B) {
	hashKey := hash.HashFuncForString[string](nil)
	eqKey := generic.NewEqualFunc[string]()
	eqVal := generic.NewEqualFunc[int]()
	opts := HashOpts{}

	b.Run("ChainHashTable.Put", func(b *testing.B) {
		ht := NewChainHashTable[string, int](hashKey, eqKey, eqVal, opts)
		runPutBenchmark(b, ht)
	})

	b.Run("LinearHashTable.Put", func(b *testing.B) {
		ht := NewLinearHashTable[string, int](hashKey, eqKey, eqVal, opts)
		runPutBenchmark(b, ht)
	})

	b.Run("QuadHashTable.Put", func(b *testing.B) {
		ht := NewQuadraticHashTable[string, int](hashKey, eqKey, eqVal, opts)
		runPutBenchmark(b, ht)
	})

	b.Run("DoubleHashTable.Put", func(b *testing.B) {
		ht := NewDoubleHashTable[string, int](hashKey, eqKey, eqVal, opts)
		runPutBenchmark(b, ht)
	})
}

func BenchmarkSymbolTable_Get(b *testing.B) {
	hashKey := hash.HashFuncForString[string](nil)
	eqKey := generic.NewEqualFunc[string]()
	eqVal := generic.NewEqualFunc[int]()
	opts := HashOpts{}

	b.Run("ChainHashTable.Get", func(b *testing.B) {
		ht := NewChainHashTable(hashKey, eqKey, eqVal, opts)
		runGetBenchmark(b, ht)
	})

	b.Run("LinearHashTable.Get", func(b *testing.B) {
		ht := NewLinearHashTable(hashKey, eqKey, eqVal, opts)
		runGetBenchmark(b, ht)
	})

	b.Run("QuadHashTable.Get", func(b *testing.B) {
		ht := NewQuadraticHashTable(hashKey, eqKey, eqVal, opts)
		runGetBenchmark(b, ht)
	})

	b.Run("DoubleHashTable.Get", func(b *testing.B) {
		ht := NewDoubleHashTable(hashKey, eqKey, eqVal, opts)
		runGetBenchmark(b, ht)
	})
}

func BenchmarkSymbolTable_Delete(b *testing.B) {
	hashKey := hash.HashFuncForString[string](nil)
	eqKey := generic.NewEqualFunc[string]()
	eqVal := generic.NewEqualFunc[int]()
	opts := HashOpts{}

	b.Run("ChainHashTable.Delete", func(b *testing.B) {
		ht := NewChainHashTable(hashKey, eqKey, eqVal, opts)
		runDeleteBenchmark(b, ht)
	})

	b.Run("LinearHashTable.Delete", func(b *testing.B) {
		ht := NewLinearHashTable(hashKey, eqKey, eqVal, opts)
		runDeleteBenchmark(b, ht)
	})

	b.Run("QuadHashTable.Delete", func(b *testing.B) {
		ht := NewQuadraticHashTable(hashKey, eqKey, eqVal, opts)
		runDeleteBenchmark(b, ht)
	})

	b.Run("DoubleHashTable.Delete", func(b *testing.B) {
		ht := NewDoubleHashTable(hashKey, eqKey, eqVal, opts)
		runDeleteBenchmark(b, ht)
	})
}

func BenchmarkOrderedSymbolTable_Put(b *testing.B) {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	b.Run("BST.Put", func(b *testing.B) {
		bst := NewBST(cmpKey, eqVal)
		runPutBenchmark(b, bst)
	})

	b.Run("AVL.Put", func(b *testing.B) {
		avl := NewAVL(cmpKey, eqVal)
		runPutBenchmark(b, avl)
	})

	b.Run("RedBlack.Put", func(b *testing.B) {
		rb := NewRedBlack(cmpKey, eqVal)
		runPutBenchmark(b, rb)
	})
}

func BenchmarkOrderedSymbolTable_Get(b *testing.B) {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	b.Run("BST.Get", func(b *testing.B) {
		bst := NewBST(cmpKey, eqVal)
		runGetBenchmark(b, bst)
	})

	b.Run("AVL.Get", func(b *testing.B) {
		avl := NewAVL(cmpKey, eqVal)
		runGetBenchmark(b, avl)
	})

	b.Run("RedBlack.Get", func(b *testing.B) {
		rb := NewRedBlack(cmpKey, eqVal)
		runGetBenchmark(b, rb)
	})
}

func BenchmarkOrderedSymbolTable_Delete(b *testing.B) {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	b.Run("BST.Delete", func(b *testing.B) {
		bst := NewBST(cmpKey, eqVal)
		runDeleteBenchmark(b, bst)
	})

	b.Run("AVL.Delete", func(b *testing.B) {
		avl := NewAVL(cmpKey, eqVal)
		runDeleteBenchmark(b, avl)
	})

	b.Run("RedBlack.Delete", func(b *testing.B) {
		rb := NewRedBlack(cmpKey, eqVal)
		runDeleteBenchmark(b, rb)
	})
}
