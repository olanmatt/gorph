package gorph

type Graph struct {
	nodes []Node
}

func (g *Graph) AddNode(key interface{}, value interface{}){
	// TODO check if node already exists
	g.nodes = append(g.nodes, Node{key, value, []Edge{}})
}

func (g *Graph) AddEdge(a interface{}, b interface{}, weight float32){
	bn := g.Search(b)
	if bn != nil {
		for i, n := range g.nodes {
			if n.key  == a {
				g.nodes[i].edges = append(n.edges, Edge{bn, weight})
			}
		}
	}
}

/*
func (g *Graph) RemoveNode(key interface{}) {

}

func (g *Graph) RemoveEdge(a interface{}, b interface{}){

}
*/

func (g *Graph) Search(key interface{}) *Node{
	for i, n := range g.nodes {
		if n.key  == key {
			return &(g.nodes[i])
		}
	}
	return nil
}

