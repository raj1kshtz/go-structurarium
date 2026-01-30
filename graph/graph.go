package graph

type genericAdjacencyListGraph[N comparable, E any] struct {
	adj map[N]map[N]E
}

func newGenericAdjacencyListGraph[N comparable, E any]() *genericAdjacencyListGraph[N, E] {
	return &genericAdjacencyListGraph[N, E]{adj: make(map[N]map[N]E)}
}

func (g *genericAdjacencyListGraph[N, E]) addVertex(node N) {
	if _, exists := g.adj[node]; !exists {
		g.adj[node] = make(map[N]E)
	}
}

func (g *genericAdjacencyListGraph[N, E]) removeVertex(node N) {
	delete(g.adj, node)
	for _, neighbors := range g.adj {
		delete(neighbors, node)
	}
}

func (g *genericAdjacencyListGraph[N, E]) addEdge(from, to N, edge E) {
	g.addVertex(from)
	g.addVertex(to)
	g.adj[from][to] = edge
	g.adj[to][from] = edge
}

func (g *genericAdjacencyListGraph[N, E]) removeEdge(from, to N) {
	if _, ok := g.adj[from]; ok {
		delete(g.adj[from], to)
	}
	if _, ok := g.adj[to]; ok {
		delete(g.adj[to], from)
	}
}

func (g *genericAdjacencyListGraph[N, E]) neighbors(node N) []N {
	neighbors := []N{}
	if adj, ok := g.adj[node]; ok {
		for n := range adj {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (g *genericAdjacencyListGraph[N, E]) hasVertex(node N) bool {
	_, ok := g.adj[node]
	return ok
}

func (g *genericAdjacencyListGraph[N, E]) hasEdge(from, to N) bool {
	if adj, ok := g.adj[from]; ok {
		_, exists := adj[to]
		return exists
	}
	return false
}

func (g *genericAdjacencyListGraph[N, E]) vertices() []N {
	vs := make([]N, 0, len(g.adj))
	for v := range g.adj {
		vs = append(vs, v)
	}
	return vs
}

func (g *genericAdjacencyListGraph[N, E]) edges() [][3]interface{} {
	es := [][3]interface{}{}
	seen := make(map[[2]interface{}]bool)
	for from, neighbors := range g.adj {
		for to, edge := range neighbors {
			key := [2]interface{}{from, to}
			keyRev := [2]interface{}{to, from}
			if !seen[key] && !seen[keyRev] {
				es = append(es, [3]interface{}{from, to, edge})
				seen[key] = true
				seen[keyRev] = true
			}
		}
	}
	return es
}
