package gorph

type Graph struct {
	nodes []Node
	edges []*Edge
	// points to all edges so they can be searched through much faster
	// eg when deleting a node, need to search for all nodes that point to the
	// deleted node
}

/*
Lists all nodes in graph.
TODO check if address of returned array is same as one in struct
*/
func (g *Graph) GetNodes() []Node {
	return g.nodes
}

/*
Lists all edges in graph.
TODO check if address of returned array is same as one in struct
*/
func (g *Graph) GetEdges() []*Edge {
	return g.edges
}

/*
Tests wheather there is an edge from node x to y.
*/
func (g *Graph) Adjacent(x interface{}, y interface{}) bool{
	return true
}

/*
Lists all nodes y such that there is an edge from x to y.
*/
func (g *Graph) Neighbors(x interface{}, y interface{}) []*Node {
	return nil
}

/*
Adds the node with the key-value pair into graph, if it is not there.
*/
func (g *Graph) AddNode(key interface{}, value interface{}){
	if g.Search(key) == nil {
		g.nodes = append(g.nodes, Node{key, value, []Edge{}})
	}
}

/*
Adds the edge from x to y, if it is not there.
TODO need to check if it is not there
*/
func (g *Graph) AddEdge(a interface{}, b interface{}, weight float64){
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
Removes the node with the key from the graph, if it is there.
*/
func (g *Graph) DeleteNode(key interface{}) {
	for i, n := range g.nodes {
		if n.key == key {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			break
		}
	}
}

/*
Removes the edge from x to y, if it is there.
TODO change search to Search() function
*/
func (g *Graph) DeleteEdge(source interface{}, target interface{}){
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
Returns pointer to node with given key.
*/
func (g *Graph) Search(key interface{}) *Node{
	for i, n := range g.nodes {
		if n.key  == key {
			return &(g.nodes[i])
		}
	}
	return nil
}
