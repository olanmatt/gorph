package gorph

type Node struct {
	key interface{}
	value interface{}
	edges []Edge
}

func (n *Node) GetKey() interface{} {
	return n.key
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func (n *Node) SetValue(value interface{}) {
	n.value = value
}

func (n *Node) GetEdges() []Edge {
	return n.edges
}
