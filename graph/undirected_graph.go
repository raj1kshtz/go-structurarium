package graph

type UndirectedGraph[N comparable, E any] struct {
	g *genericAdjacencyListGraph[N, E]
}

func NewUndirectedGraph[N comparable, E any]() *UndirectedGraph[N, E] {
	return &UndirectedGraph[N, E]{g: newGenericAdjacencyListGraph[N, E]()}
}

func (ug *UndirectedGraph[N, E]) AddVertex(node N) {
	ug.g.addVertex(node)
}

func (ug *UndirectedGraph[N, E]) RemoveVertex(node N) {
	ug.g.removeVertex(node)
}

func (ug *UndirectedGraph[N, E]) AddEdge(from, to N, edge E) {
	ug.g.addEdge(from, to, edge)
}

func (ug *UndirectedGraph[N, E]) RemoveEdge(from, to N) {
	ug.g.removeEdge(from, to)
}

func (ug *UndirectedGraph[N, E]) Neighbors(node N) []N {
	return ug.g.neighbors(node)
}

func (ug *UndirectedGraph[N, E]) HasVertex(node N) bool {
	return ug.g.hasVertex(node)
}

func (ug *UndirectedGraph[N, E]) HasEdge(from, to N) bool {
	return ug.g.hasEdge(from, to)
}

func (ug *UndirectedGraph[N, E]) Vertices() []N {
	return ug.g.vertices()
}

func (ug *UndirectedGraph[N, E]) Edges() [][3]interface{} {
	return ug.g.edges()
}
