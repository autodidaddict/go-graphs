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
          g := &Graph{}
          v := g.AddVertex("bob")
          Ω(v).ShouldNot(BeNil())
          Ω(v.Key).Should(Equal("bob"))
          Ω(len(g.Canvas)).Should(Equal(1))
        })

        It("Should properly add an edge between two vertices", func() {
          g := &Graph{}
          start := g.AddVertex("start")
          finish := g.AddVertex("finish")
          g.AddEdge(start, finish, 1)
          Ω(len(g.Canvas)).Should(Equal(2))
          Ω(len(start.Neighbors)).Should(Equal(1))
          // start-(edge[0])-finish
          Ω(start.Neighbors[0].Neighbor.Key).Should(Equal("finish"))
        })
      })
    })
})
