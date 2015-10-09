package graphs

// Path represents the output of a traversal or can be used to record
// an arbitrary path through a graph.
type Path struct {
    Total             int
    Destination       *Vertex
    Previous          *Path
}

// NewPath creates a new path with its values properly initialized.
func NewPath() *Path {
  p := new(Path)
  p.Destination = &Vertex{}
  p.Total = -1 // indicate an uninitialized path versus an empty path.
  return p
}
