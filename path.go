package graphs

import (
  "bytes"
  "fmt"
)

// Path represents the output of a traversal or can be used to record
// an arbitrary path through a graph.
type Path struct {
    Step      *Vertex
    Outbound  *Edge
    Next      *Path
}

// NewPath creates a new path with no outbound edges starting at the
// indicated vertex.
func NewPath(start *Vertex) *Path {
  p := &Path{}
  p.Step = start
  return p
}

// AppendVertex adds a vertex to the path, traversing the indicated
// edge to reach said vertex.
func (p *Path) AppendVertex(edge *Edge, vert *Vertex) {
  newPath := NewPath(vert)
  p.Outbound = edge
  p.Next = newPath
}


// String provides a nice readable view of a path in the format
// V(id)-->(edge label)-->V(id)--  etc.
func (p *Path) String() string {
    var buffer bytes.Buffer

    buffer.WriteString(p.Step.Key)

    for p.Next != nil {
      buffer.WriteString(fmt.Sprintf("--(%s:%d)-->", p.Outbound.Label, p.Outbound.Weight))
      p = p.Next
      buffer.WriteString(p.Step.Key)
    }

    return buffer.String()
}
