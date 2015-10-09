package graphs

// Edge represents a directed edge from an origin vertex to
// a destination vertex, with an accompanying label and weight.
type Edge struct {
      From        *Vertex
      To          *Vertex
      Weight      int
      Label       string
}
