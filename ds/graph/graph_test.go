package graph

type testVisitors struct {
	*Visitors
	preOrderVertices  []int
	postOrderVertices []int
	preOrderEdges     [][2]int
	preOrderWeights   []float64
}

func newTestVisitors() *testVisitors {
	tv := &testVisitors{
		preOrderVertices:  make([]int, 0),
		postOrderVertices: make([]int, 0),
		preOrderEdges:     make([][2]int, 0),
		preOrderWeights:   make([]float64, 0),
	}

	tv.Visitors = &Visitors{
		VertexPreOrder: func(v int) bool {
			tv.preOrderVertices = append(tv.preOrderVertices, v)
			return true
		},
		VertexPostOrder: func(v int) bool {
			tv.postOrderVertices = append(tv.postOrderVertices, v)
			return true
		},
		EdgePreOrder: func(v, w int, weight float64) bool {
			tv.preOrderEdges = append(tv.preOrderEdges, [2]int{v, w})
			tv.preOrderWeights = append(tv.preOrderWeights, weight)
			return true
		},
	}

	return tv
}
