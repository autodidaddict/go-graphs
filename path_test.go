package graphs_test

import (
  . "github.com/autodidaddict/go-graphs"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "fmt"
)

var _ = Describe("Go Graphs", func() {
    Describe("Path Management", func() {

        Context("Build a simple path", func() {
          It("Should pretty-print a chain of vertices and edges.", func() {

            g := NewGraph()
            start, _ := g.AddVertex("start")
            finish, _ := g.AddVertex("finish")
            e, _ := g.AddEdge("start", "finish", 1)
            e.Label = "run"

            path := NewPath(start)
            path.AppendVertex(e, finish)

            pretty := path.String()
            fmt.Println(pretty)
            Ω(pretty).Should(Equal("start--(run:1)-->finish"))
          })
        })
    })
})
