package graphs

import "fmt"

// Graph represents the container in which vertices are placed, which
// includes functionality for traversing and interrogating graphs.
type Graph struct {
  Vertices    map[string]*Vertex
  Edges       []*Edge
}

// NewGraph creates an instance of a new graph
func NewGraph() *Graph {
	var graph Graph
	graph.Vertices = make(map[string]*Vertex, 0)
	graph.Edges = make([]*Edge, 0)

	return &graph
}

// AddVertex adds a vertex to an existing graph
func (g *Graph) AddVertex(key string) (*Vertex, error) {
	if _, ok := g.Vertices[key]; ok {
		return nil, fmt.Errorf("Cannot add vertex \"%v\", already in Graph.", key)
	}
  v :=  &Vertex{
		  Key: key,
		  Neighbors: make([]*Edge, 0),
	}
  g.Vertices[key] = v
	return v, nil
}

// AddEdge adds a directed edge between a source and a target vertex of
// a given weight. Weight is used in assisting pathfinding algorithms.
func (g *Graph) AddEdge(from string, to string, weight int) error {
	var ok bool
	fromVertex, ok := g.Vertices[from]
	if !ok {
		return fmt.Errorf("Vertex \"%v\" does not exist", from)
	}
	toVertex, ok := g.Vertices[to]
	if !ok {
		return fmt.Errorf("Vertex \"%v\" does not exist", from)
	}

	edge := Edge{
		From:   fromVertex,
		To:     toVertex,
		Weight: weight,
	}

	g.Edges = append(g.Edges, &edge)
	fromVertex.Neighbors= append(fromVertex.Neighbors, &edge)

	return nil
}
