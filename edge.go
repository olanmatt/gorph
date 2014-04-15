package gorph

/*
Edges are only unidirectional. Bidirectionality comes from edges mutually
pointing to each other.
*/
type Edge struct {
	p *Node
	weight float32
}

