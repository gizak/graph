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

type adjList []adjNode

// NewAdjList returns a adjancency list structure implementing Graph interface
func NewAdjList() *adjList {
	foo := adjList([]adjNode{})
	return &foo
}

func (g *adjList) getNode(id int) *adjNode {
	for i := range *g {
		if (*g)[i].id == id {
			return &(*g)[i]
		}
	}
	return nil
}

func (g *adjList) SetVertex(id int, val interface{}) {
	for i, v := range *g {
		if v.id == id {
			(*g)[i].val = val
			return
		}
	}
	*g = append(*g, adjNode{node{id, val}, nil})
}

func (g adjList) GetVertex(id int) interface{} {
	for _, v := range g {
		if v.id == id {
			return v.val
		}
	}
	return nil
}

func (g *adjList) DelVertex(id int) {
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

func (g *adjList) SetEdge(from, to int, v interface{}) {
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

func (g *adjList) GetEdge(from, to int) interface{} {
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

func (g *adjList) DelEdge(from, to int) {
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

func (g *adjList) GetNeighbours(id int) []int {
	nb := []int{}
	n := g.getNode(id)
	if n != nil {
		for p := n.next; p != nil; p = p.next {
			nb = append(nb, p.id)
		}
	}
	return nb
}

func (g *adjList) GetInverseNbs(id int) []int {
	nb := []int{}
	g.IterVertices(func(gg Graph, from int) {
		if gg.GetEdge(from, id) != nil {
			nb = append(nb, from)
		}
	})
	return nb
}

func (g *adjList) hasEdge(from, to int) bool {
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

func (g *adjList) IsAdjacent(a, b int) bool {
	return g.hasEdge(a, b) || g.hasEdge(b, a)
}

func (g *adjList) IterEdges(f func(Graph, int, int)) {
	g.IterVertices(func(g Graph, from int) {
		nb := g.GetNeighbours(from)
		for _, to := range nb {
			f(g, from, to)
		}
	})
}

func (g *adjList) IterVertices(f func(Graph, int)) {
	for i := range *g {
		f(g, (*g)[i].id)
	}
}
