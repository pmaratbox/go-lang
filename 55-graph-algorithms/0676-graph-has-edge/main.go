package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
)

var id = map[string]int64{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}

func main() {
	g := simple.NewWeightedUndirectedGraph(0, 0)
	for n := int64(0); n < 5; n++ {
		g.AddNode(simple.Node(n))
	}
	type E struct {
		u, v string
		w    float64
	}
	for _, e := range []E{{"a", "b", 1}, {"a", "c", 4}, {"b", "c", 1}, {"b", "d", 5}, {"c", "d", 1}, {"d", "e", 1}} {
		g.SetWeightedEdge(g.NewWeightedEdge(simple.Node(id[e.u]), simple.Node(id[e.v]), e.w))
	}

	bc := g.HasEdgeBetween(id["b"], id["c"])
	ae := g.HasEdgeBetween(id["a"], id["e"])
	fmt.Println(bc, ae)
}
