package main

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

var lbl = map[int64]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e"}
var id = map[string]int64{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}

func main() {
	d := simple.NewDirectedGraph()
	for n := int64(0); n < 5; n++ {
		d.AddNode(simple.Node(n))
	}
	for _, e := range [][2]string{
		{"a", "b"}, {"b", "c"}, {"a", "c"}, {"c", "d"}, {"d", "e"},
	} {
		d.SetEdge(d.NewEdge(simple.Node(id[e[0]]), simple.Node(id[e[1]])))
	}

	order, _ := topo.Sort(d)
	var os []string
	for _, n := range order {
		os = append(os, lbl[n.ID()])
	}
	fmt.Println(strings.Join(os, ","))
}
