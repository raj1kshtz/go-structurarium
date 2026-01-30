package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testEdgeWrapper struct {
	weight int
}

type GraphWrapperTestSuite struct {
	suite.Suite
	gw *GraphWrapper[int, testEdgeWrapper]
}

func (s *GraphWrapperTestSuite) SetupTest() {
	s.gw = NewGraphWrapper[int, testEdgeWrapper]()
}

func (s *GraphWrapperTestSuite) TestGraphWrapper() {
	s.Run("AddVertexAndHasVertex", func() {
		s.gw.AddVertex(1)
		s.gw.AddVertex(2)
		s.True(s.gw.HasVertex(1))
		s.True(s.gw.HasVertex(2))
	})
	s.Run("AddEdgeAndHasEdge", func() {
		s.gw.AddVertex(1)
		s.gw.AddVertex(2)
		s.gw.AddEdge(1, 2, testEdgeWrapper{weight: 5})
		s.True(s.gw.HasEdge(1, 2))
	})
	s.Run("Neighbors", func() {
		s.gw.AddVertex(1)
		s.gw.AddVertex(2)
		s.gw.AddEdge(1, 2, testEdgeWrapper{weight: 5})
		neighbors := s.gw.Neighbors(1)
		s.Equal([]int{2}, neighbors)
	})
	s.Run("RemoveEdge", func() {
		s.gw.AddVertex(1)
		s.gw.AddVertex(2)
		s.gw.AddEdge(1, 2, testEdgeWrapper{weight: 5})
		s.gw.RemoveEdge(1, 2)
		s.False(s.gw.HasEdge(1, 2))
	})
	s.Run("RemoveVertex", func() {
		s.gw.AddVertex(1)
		s.gw.RemoveVertex(1)
		s.False(s.gw.HasVertex(1))
	})
	s.Run("VerticesAndEdges", func() {
		s.gw.AddVertex(2)
		s.gw.AddVertex(3)
		s.gw.AddEdge(2, 3, testEdgeWrapper{weight: 7})
		vs := s.gw.Vertices()
		s.ElementsMatch([]int{2, 3}, vs)
		es := s.gw.Edges()
		s.Len(es, 1)
	})
}

func TestGraphWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(GraphWrapperTestSuite))
}
