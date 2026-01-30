package graph

type DirectedGraph[N comparable, E any] struct {
	g *genericAdjacencyListGraph[N, E]
}

func NewDirectedGraph[N comparable, E any]() *DirectedGraph[N, E] {
	return &DirectedGraph[N, E]{g: newGenericAdjacencyListGraph[N, E]()}
}

func (dg *DirectedGraph[N, E]) AddVertex(node N) {
	dg.g.addVertex(node)
}

func (dg *DirectedGraph[N, E]) RemoveVertex(node N) {
	dg.g.removeVertex(node)
}

func (dg *DirectedGraph[N, E]) AddEdge(from, to N, edge E) {
	dg.g.addVertex(from)
	dg.g.addVertex(to)
	dg.g.adj[from][to] = edge // Only one direction
}

func (dg *DirectedGraph[N, E]) RemoveEdge(from, to N) {
	if _, ok := dg.g.adj[from]; ok {
		delete(dg.g.adj[from], to)
	}
}

func (dg *DirectedGraph[N, E]) Neighbors(node N) []N {
	return dg.g.neighbors(node)
}

func (dg *DirectedGraph[N, E]) HasVertex(node N) bool {
	return dg.g.hasVertex(node)
}

func (dg *DirectedGraph[N, E]) HasEdge(from, to N) bool {
	return dg.g.hasEdge(from, to)
}

func (dg *DirectedGraph[N, E]) Vertices() []N {
	return dg.g.vertices()
}

func (dg *DirectedGraph[N, E]) Edges() [][3]interface{} {
	es := [][3]interface{}{}
	for from, neighbors := range dg.g.adj {
		for to, edge := range neighbors {
			es = append(es, [3]interface{}{from, to, edge})
		}
	}
	return es
}
