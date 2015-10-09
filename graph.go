package graphs

// Graph represents the container in which vertices are placed, which
// includes functionality for traversing and interrogating graphs.
type Graph struct {
  Canvas        []*Vertex
}

// AddVertex adds a vertex to an existing graph
func (g *Graph) AddVertex(key string) *Vertex {
  v := &Vertex{}
  v.Key = key
  g.Canvas = append(g.Canvas, v)
  return v
}

// AddEdge adds a directed edge between a source and a target vertex of
// a given weight. Weight is used in assisting pathfinding algorithms.
func (g *Graph) AddEdge(source *Vertex, target *Vertex, weight int) {
  e := &Edge{}
  e.Neighbor = target
  e.Weight = weight
  source.Neighbors = append(source.Neighbors, e)
}
