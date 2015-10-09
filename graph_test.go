package graphs_test

import (
  . "github.com/autodidaddict/go-graphs"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Go Graphs", func() {
    Describe("Graph Building", func() {
      // initialize fixtures here?

      Context("Build a simple graph", func() {
        It("Should add vertices with the AddVertex method", func() {
          g := NewGraph()
          v, err := g.AddVertex("bob")
          Ω(v).ShouldNot(BeNil())
          Ω(err).Should(BeNil())
          Ω(v.Key).Should(Equal("bob"))
          Ω(len(g.Vertices)).Should(Equal(1))
        })

        It("Should properly add an edge between two vertices", func() {
          g := NewGraph()
          start, err := g.AddVertex("start")
          Ω(err).Should(BeNil())
          finish, err := g.AddVertex("finish")
          g.AddEdge("start", "finish", 1)
          Ω(len(g.Vertices)).Should(Equal(2))
          Ω(len(g.Edges)).Should(Equal(1))
          Ω(len(start.Neighbors)).Should(Equal(1))
          Ω(len(finish.Neighbors)).Should(Equal(0))
          // start-(edge[0])-finish
          Ω(start.Neighbors[0].To.Key).Should(Equal("finish"))
        })

      })
    })
})
