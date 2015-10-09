package graphs

// Vertex represents a vertex within some graph canvas
type Vertex struct {
      Key           string
      Neighbors     []*Edge
}
