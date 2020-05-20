package graph

type visitor struct {
	vertices []int
	edges    [][2]int
	weights  []float64
}

func newVisitor() *visitor {
	return &visitor{
		vertices: make([]int, 0),
		edges:    make([][2]int, 0),
		weights:  make([]float64, 0),
	}
}

func (vis *visitor) VisitVertex(v int) bool {
	vis.vertices = append(vis.vertices, v)
	return true
}

func (vis *visitor) VisitEdge(v, w int) bool {
	vis.edges = append(vis.edges, [2]int{v, w})
	return true
}

func (vis *visitor) VisitWeightedEdge(v, w int, weight float64) bool {
	vis.edges = append(vis.edges, [2]int{v, w})
	vis.weights = append(vis.weights, weight)
	return true
}
