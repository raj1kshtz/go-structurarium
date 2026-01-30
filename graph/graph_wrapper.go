package graph

type GraphWrapper[N comparable, E any] struct {
	graph *genericAdjacencyListGraph[N, E]
}

func NewGraphWrapper[N comparable, E any]() *GraphWrapper[N, E] {
	return &GraphWrapper[N, E]{graph: newGenericAdjacencyListGraph[N, E]()}
}

func (gw *GraphWrapper[N, E]) AddVertex(node N) {
	gw.graph.addVertex(node)
}

func (gw *GraphWrapper[N, E]) RemoveVertex(node N) {
	gw.graph.removeVertex(node)
}

func (gw *GraphWrapper[N, E]) AddEdge(from, to N, edge E) {
	gw.graph.addEdge(from, to, edge)
}

func (gw *GraphWrapper[N, E]) RemoveEdge(from, to N) {
	gw.graph.removeEdge(from, to)
}

func (gw *GraphWrapper[N, E]) Neighbors(node N) []N {
	return gw.graph.neighbors(node)
}

func (gw *GraphWrapper[N, E]) HasVertex(node N) bool {
	return gw.graph.hasVertex(node)
}

func (gw *GraphWrapper[N, E]) HasEdge(from, to N) bool {
	return gw.graph.hasEdge(from, to)
}

func (gw *GraphWrapper[N, E]) Vertices() []N {
	return gw.graph.vertices()
}

func (gw *GraphWrapper[N, E]) Edges() [][3]interface{} {
	return gw.graph.edges()
}
