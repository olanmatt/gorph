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
	g.AddEdge(1, 2, 8.0)
	fmt.Println(g)
}
