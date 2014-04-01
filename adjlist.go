package graph

type node struct {
	id  int
	val interface{}
}

type edgeCh struct {
	next *edgeCh
	id   int
	edge interface{}
}

type adjNode struct {
	node
	next *edgeCh
}

// AdjList is the adjacency list implementation
type AdjList []adjNode

// NewAdjList returns a adjancency list structure implementing Graph interface
func NewAdjList() *AdjList {
	foo := AdjList([]adjNode{})
	return &foo
}

func (g *AdjList) getNode(id int) *adjNode {
	for i := range *g {
		if (*g)[i].id == id {
			return &(*g)[i]
		}
	}
	return nil
}

// SetVertex uses pointer receiver to add or update a value given id.
func (g *AdjList) SetVertex(id int, val interface{}) {
	for i, v := range *g {
		if v.id == id {
			(*g)[i].val = val
			return
		}
	}
	*g = append(*g, adjNode{node{id, val}, nil})
}

// GetVertex returns the data associating with the given id, nil will be returned if there is nothing being stored.
func (g AdjList) GetVertex(id int) interface{} {
	for _, v := range g {
		if v.id == id {
			return v.val
		}
	}
	return nil
}

// DelVertex deletes the node (if exists) given id.
func (g *AdjList) DelVertex(id int) {
	idx := -1
	for i, v := range *g {
		if v.id == id {
			idx = i
		}
	}
	if idx != -1 {
		*g = append((*g)[:idx], (*g)[idx+1:]...)
	}
	// del corresponding edges
	nbs := g.GetInverseNbs(id)
	for _, v := range nbs {
		g.DelEdge(v, id)
	}
}

// SetEdge sets an edge from an node to the other given ids with value v.
func (g *AdjList) SetEdge(from, to int, v interface{}) {
	n0 := g.getNode(from)
	n1 := g.getNode(to)
	if n0 != nil && n1 != nil {
		i := n0.next
		// add head
		if i == nil {
			n0.next = &edgeCh{nil, to, v}
			return
		}
		// update
		for ; ; i = i.next {
			if i.id == to {
				i.edge = v
				return
			}
			if i.next == nil {
				break
			}
		}

		// add tail
		i.next = &edgeCh{nil, to, v}
	}
}

// GetEdge retreive the edge value, nil returns if it is non-existance.
func (g *AdjList) GetEdge(from, to int) interface{} {
	n0 := g.getNode(from)
	n1 := g.getNode(to)

	if n0 != nil && n1 != nil {
		for i := n0.next; i != nil; i = i.next {
			if i.id == n1.id {
				return i.edge
			}
		}
	}
	return nil
}

// DelEdge removes the edge from the underlying graph.
func (g *AdjList) DelEdge(from, to int) {
	n0 := g.getNode(from)
	n1 := g.getNode(to)

	if n0 != nil && n1 != nil && g.GetEdge(from, to) != nil {
		p := n0.next
		// if first
		if p.id == to {
			n0.next = p.next
			return
		}
		// otherwise
		for i := p.next; i != nil; i = i.next {
			if i.id == to {
				p.next = i.next
				return
			}
			p = p.next
		}
	}
}

// GetNeighbours returns the nodes' ids list pointed by the given node.
func (g *AdjList) GetNeighbours(id int) []int {
	nb := []int{}
	n := g.getNode(id)
	if n != nil {
		for p := n.next; p != nil; p = p.next {
			nb = append(nb, p.id)
		}
	}
	return nb
}

// GetInverseNbs returns the nodes' id slice, which contains the nodes point to the given node.
func (g *AdjList) GetInverseNbs(id int) []int {
	nb := []int{}
	g.IterVertices(func(gg Graph, from int) {
		if gg.GetEdge(from, id) != nil {
			nb = append(nb, from)
		}
	})
	return nb
}

func (g *AdjList) hasEdge(from, to int) bool {
	n0, n1 := g.getNode(from), g.getNode(to)

	if n0 != nil && n1 != nil {
		for p := n0.next; p != nil; p = p.next {
			if p.id == n1.id {
				return true
			}
		}
	}
	return false
}

// IsAdjacent test if a and b is adjancent.
func (g *AdjList) IsAdjacent(a, b int) bool {
	return g.hasEdge(a, b) || g.hasEdge(b, a)
}

// IterEdges iterates all the edges, passing the edges' from and to nodes id to the argument function one by one.
func (g *AdjList) IterEdges(f func(Graph, int, int)) {
	g.IterVertices(func(g Graph, from int) {
		nb := g.GetNeighbours(from)
		for _, to := range nb {
			f(g, from, to)
		}
	})
}

// IterVertices iterates all the vertices in the graph, passing the nodes' ids to the given function argument.
func (g *AdjList) IterVertices(f func(Graph, int)) {
	for i := range *g {
		f(g, (*g)[i].id)
	}
}
