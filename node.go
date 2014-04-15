package gorph

type Node struct {
	key interface{}
	value interface{}
	edges []Edge
}
