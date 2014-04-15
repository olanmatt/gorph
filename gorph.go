package gorph

type Graph struct {
	nodes []Node
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

