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

func (g *adjList) GetNeighbours(id int) []interface{} {
	return nil
}

func (g *adjList) IsAdjacent(a, b int) bool {
	return false
}

func (g *adjList) IterEdges(f func(int, int)) {
}

func (g *adjList) IterVertices(f func(int)) {
}
