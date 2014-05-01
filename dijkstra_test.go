package gorph

import (
	"testing"
	//"fmt"
)

func TestDijkstra(t *testing.T) {
	g := Graph{}
	g.AddNode("A", "hello")
	g.AddNode("B", "world")
	g.AddNode("C", "this")
	g.AddNode("D", "is")
	g.AddNode("E", "matt")
	g.AddNode("F", "here")
	g.AddNode("G", "talking")
	g.AddNode("H", "loudly")
	g.AddEdge("A", "B", 8.0)
	g.AddEdge("A", "E", 12.0)
	g.AddEdge("E", "A", 12.0)
	g.AddEdge("C", "B", 4.0)
	g.AddEdge("B", "D", 3.0)
	g.AddEdge("D", "C", 5.0)
	g.AddEdge("B", "F", 9.0)
	g.AddEdge("D", "G", 1.0)
	g.AddEdge("G", "F", 2.0)
	g.AddEdge("F", "H", 1.0)
	g.AddEdge("H", "E", 1.0)
	// fmt.Println(g)
	g.Dijkstra("A", "H")
}
