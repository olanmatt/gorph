package gorph

import (
	"testing"
	"fmt"
)

func TestDijkstra(t *testing.T) {
	g := Graph{}
	g.AddNode(Node{1, "A", []Edge{}})
	g.AddNode(Node{2, "B", []Edge{}})
	g.AddNode(Node{3, "C", []Edge{}})
	g.AddNode(Node{4, "D", []Edge{}})
	g.AddNode(Node{5, "E", []Edge{}})
	g.AddNode(Node{6, "F", []Edge{}})
	g.AddNode(Node{7, "G", []Edge{}})
	g.AddNode(Node{8, "H", []Edge{}})
	g.AddEdge(1, 2, 8.0)
	g.AddEdge(1, 5, 12.0)
	g.AddEdge(5, 1, 12.0)
	g.AddEdge(3, 2, 4.0)
	g.AddEdge(2, 4, 3.0)
	g.AddEdge(4, 3, 5.0)
	g.AddEdge(2, 6, 9.0)
	g.AddEdge(4, 7, 1.0)
	g.AddEdge(7, 6, 2.0)
	g.AddEdge(6, 8, 1.0)
	g.AddEdge(8, 5, 1.0)
	fmt.Println(g)
}
