package graph

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type testEdgeUndirected struct {
	label string
}


type UndirectedGraphTestSuite struct {
	suite.Suite
	g *UndirectedGraph[int, testEdgeUndirected]
}

func (s *UndirectedGraphTestSuite) SetupTest() {
	s.g = NewUndirectedGraph[int, testEdgeUndirected]()
}

func (s *UndirectedGraphTestSuite) TestUndirectedGraph() {
	s.Run("AddVertexAndHasVertex", func() {
		s.g.AddVertex(1)
		s.g.AddVertex(2)
		s.True(s.g.HasVertex(1))
		s.True(s.g.HasVertex(2))
	})
	s.Run("AddEdgeAndHasEdge", func() {
		s.g.AddVertex(1)
		s.g.AddVertex(2)
		s.g.AddEdge(1, 2, testEdgeUndirected{label: "a-b"})
		s.True(s.g.HasEdge(1, 2))
		s.True(s.g.HasEdge(2, 1))
	})
	s.Run("Neighbors", func() {
		s.g.AddVertex(1)
		s.g.AddVertex(2)
		s.g.AddEdge(1, 2, testEdgeUndirected{label: "a-b"})
		neighbors1 := s.g.Neighbors(1)
		neighbors2 := s.g.Neighbors(2)
		s.Equal([]int{2}, neighbors1)
		s.Equal([]int{1}, neighbors2)
	})
	s.Run("RemoveEdge", func() {
		s.g.AddVertex(1)
		s.g.AddVertex(2)
		s.g.AddEdge(1, 2, testEdgeUndirected{label: "a-b"})
		s.g.RemoveEdge(1, 2)
		s.False(s.g.HasEdge(1, 2))
		s.False(s.g.HasEdge(2, 1))
	})
	s.Run("RemoveVertex", func() {
		s.g.AddVertex(1)
		s.g.RemoveVertex(1)
		s.False(s.g.HasVertex(1))
	})
	s.Run("VerticesAndEdges", func() {
		s.g.AddVertex(2)
		s.g.AddVertex(3)
		s.g.AddEdge(2, 3, testEdgeUndirected{label: "b-c"})
		vs := s.g.Vertices()
		s.ElementsMatch([]int{2, 3}, vs)
		es := s.g.Edges()
		s.Len(es, 1)
	})
}

func TestUndirectedGraphTestSuite(t *testing.T) {
	suite.Run(t, new(UndirectedGraphTestSuite))
}
