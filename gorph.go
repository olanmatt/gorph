package main

import "fmt"

type Graph struct {
	nodes []Node
	edges []Edge
}

func (g *Graph) AddNode(n Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(a int, b int, weight float32){
	bn := g.Search(b)
	if bn != nil {
		for i, n := range g.nodes {
			if n.id  == a {
				g.nodes[i].edges = append(n.edges, Edge{1, bn, weight})
			}
		}
	}
}

func (g *Graph) Search(id int) *Node{
	for i, n := range g.nodes {
		if n.id  == id {
			return &(g.nodes[i])
		}
	}
	return nil
}

type Node struct {
	id int
	val string
	edges []Edge
}

/*
   Edges are only unidirectional. Bidirectionality comes from edges mutually
   pointing to each other.
 */
type Edge struct {
	id uint32
	p *Node
	weight float32
}

func main() {
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
