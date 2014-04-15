package gorph

type Graph struct {
	nodes []Node
}

/*
Runs in O(n) time
*/
func (g *Graph) AddNode(key interface{}, value interface{}){
	if g.Search(key) == nil {
		g.nodes = append(g.nodes, Node{key, value, []Edge{}})
	}
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
Runs in O(n) time.
*/
func (g *Graph) RemoveNode(key interface{}) {
	for i, n := range g.nodes {
		if n.key == key {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			break
		}
	}
}

/*
Runs in O(n) time.
*/
func (g *Graph) RemoveEdge(source interface{}, target interface{}){
	for i, n := range g.nodes {
		if n.key == source {
			for j, e := range g.nodes[i].edges{
				if e.target.key == target {
					g.nodes[i].edges = append(g.nodes[i].edges[:j], g.nodes[i].edges[j+1:]...)
				}
			}
			break;
		}
	}
}

/*
Runs in O(n) time.
*/
func (g *Graph) Search(key interface{}) *Node{
	for i, n := range g.nodes {
		if n.key  == key {
			return &(g.nodes[i])
		}
	}
	return nil
}

