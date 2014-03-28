package graph

// Graph interfece is a super set of standard graph struct interface
type Graph interface {
	SetVertex(id int, v interface{})
	DelVertex(id int)
	GetVertex(id int) interface{}

	SetEdge(a, b int, e interface{})
	DelEdge(a, b int)
	GetEdge(a, b int) interface{}

	GetNeighbours(id int) []interface{}
	IsAdjacent(a, b int) bool

	IterVertices(func(int))
	IterEdges(func(int, int))
}
