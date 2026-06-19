package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

var lbl = map[int64]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e"}
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

	// Connectivity: a path exists between a and e iff e is reachable from a.
	pt := path.DijkstraFrom(simple.Node(id["a"]), g)
	route, _ := pt.To(id["e"])
	fmt.Println(len(route) > 0)
}
