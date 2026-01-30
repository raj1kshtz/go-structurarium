package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testEdgeGraph struct {
	weight int
}

type CoreGraphTestSuite struct {
	suite.Suite
	g *genericAdjacencyListGraph[int, testEdgeGraph]
}

func (s *CoreGraphTestSuite) SetupTest() {
	s.g = newGenericAdjacencyListGraph[int, testEdgeGraph]()
}

func (s *CoreGraphTestSuite) TestCoreGraph() {
	s.Run("AddVertexAndHasVertex", func() {
		s.g.addVertex(1)
		s.g.addVertex(2)
		s.True(s.g.hasVertex(1))
		s.True(s.g.hasVertex(2))
	})
	s.Run("AddEdgeAndHasEdge", func() {
		s.g.addVertex(1)
		s.g.addVertex(2)
		s.g.addEdge(1, 2, testEdgeGraph{weight: 10})
		s.True(s.g.hasEdge(1, 2))
		s.True(s.g.hasEdge(2, 1))
	})
	s.Run("Neighbors", func() {
		s.g.addVertex(1)
		s.g.addVertex(2)
		s.g.addEdge(1, 2, testEdgeGraph{weight: 10})
		neighbors := s.g.neighbors(1)
		s.Equal([]int{2}, neighbors)
	})
	s.Run("RemoveEdge", func() {
		s.g.addVertex(1)
		s.g.addVertex(2)
		s.g.addEdge(1, 2, testEdgeGraph{weight: 10})
		s.g.removeEdge(1, 2)
		s.False(s.g.hasEdge(1, 2))
		s.False(s.g.hasEdge(2, 1))
	})
	s.Run("RemoveVertex", func() {
		s.g.addVertex(1)
		s.g.removeVertex(1)
		s.False(s.g.hasVertex(1))
	})
	s.Run("VerticesAndEdges", func() {
		s.g.addVertex(2)
		s.g.addVertex(3)
		s.g.addEdge(2, 3, testEdgeGraph{weight: 20})
		vs := s.g.vertices()
		s.ElementsMatch([]int{2, 3}, vs)
		es := s.g.edges()
		s.Len(es, 1)
	})
}

func TestCoreGraphTestSuite(t *testing.T) {
	suite.Run(t, new(CoreGraphTestSuite))
}

func TestGenericAdjacencyListGraphBasic(t *testing.T) {
	g := newGenericAdjacencyListGraph[int, testEdgeGraph]()

	g.addVertex(1)
	g.addVertex(2)
	g.addEdge(1, 2, testEdgeGraph{weight: 10})

	if !g.hasVertex(1) || !g.hasVertex(2) {
		t.Error("Vertices not added correctly")
	}
	if !g.hasEdge(1, 2) {
		t.Error("Edge not added correctly")
	}
	neighbors := g.neighbors(1)
	if len(neighbors) != 1 || neighbors[0] != 2 {
		t.Error("Neighbors not returned correctly")
	}
	g.removeEdge(1, 2)
	if g.hasEdge(1, 2) {
		t.Error("Edge not removed correctly")
	}
	g.removeVertex(1)
	if g.hasVertex(1) {
		t.Error("Vertex not removed correctly")
	}
}
