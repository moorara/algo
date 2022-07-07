package symboltable

import (
	"math/rand"
	"testing"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/sort"
)

const seed = 27

func genIntSlice(size int) []int {
	items := make([]int, size)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	sort.Shuffle(items)

	return items
}

func runPutBenchmark(b *testing.B, ost OrderedSymbolTable[int, string]) {
	items := genIntSlice(b.N)
	rand.Seed(seed)
	sort.Shuffle(items)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ost.Put(items[n], "")
	}
}

func runGetBenchmark(b *testing.B, ost OrderedSymbolTable[int, string]) {
	items := genIntSlice(b.N)
	rand.Seed(seed)
	sort.Shuffle(items)
	for n := 0; n < b.N; n++ {
		ost.Put(items[n], "")
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ost.Get(items[n])
	}
}

func BenchmarkOrderedSymbolTable(b *testing.B) {
	cmpKey := common.NewCompareFunc[int]()

	b.Run("BST.Put", func(b *testing.B) {
		ost := NewBST[int, string](cmpKey)
		runPutBenchmark(b, ost)
	})

	b.Run("BST.Get", func(b *testing.B) {
		ost := NewBST[int, string](cmpKey)
		runGetBenchmark(b, ost)
	})

	b.Run("AVL.Put", func(b *testing.B) {
		ost := NewAVL[int, string](cmpKey)
		runPutBenchmark(b, ost)
	})

	b.Run("AVL.Get", func(b *testing.B) {
		ost := NewAVL[int, string](cmpKey)
		runGetBenchmark(b, ost)
	})

	b.Run("RedBlack.Put", func(b *testing.B) {
		ost := NewRedBlack[int, string](cmpKey)
		runPutBenchmark(b, ost)
	})

	b.Run("RedBlack.Get", func(b *testing.B) {
		ost := NewRedBlack[int, string](cmpKey)
		runGetBenchmark(b, ost)
	})
}
