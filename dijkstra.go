package gorph

import (
	"fmt"
	"math"
	"container/heap"
)

var dist map[interface{}]float64
var index map[interface{}]int

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return dist[pq[i]] < dist[pq[j]]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	index[pq[i]] = i
	index[pq[j]] = j
}

// TODO REPLACE WITH DEFAULT CODE
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	index[item] = n
	*pq = append(*pq, item)
}

// TODO REPLACE WITH DEFAULT CODE
func (pq *PriorityQueue) Pop() interface{} {
	if len(*pq) > 0 {
		old := *pq
		n := len(old)
		item := old[n-1]
		index[item] = -1 // for safety
		*pq = old[0 : n-1]
		return item
	} else {
		return nil
	}
}

// Updates the priority of a Node in the queue
func (pq *PriorityQueue) update(item *Node, priority float64) {
	dist[item] = priority
	heap.Fix(pq, index[item])
}

func (g *Graph) Dijkstra(source interface{}, target interface{}) Graph{

	dist = make(map[interface{}]float64)
	index = make(map[interface{}]int)
	prev := make(map[interface{}]interface{})

	Q := &PriorityQueue{}
	heap.Init(Q)

	for i, _ := range g.nodes {
		heap.Push(Q, &g.nodes[i])
		Q.update(&g.nodes[i], math.Inf(1))
		prev[&g.nodes[i]] = nil
	}

	Q.update(g.Search(source), 0)
	
	/*
	fmt.Println()
	fmt.Println(dist)
	fmt.Println(index)
	fmt.Println(prev)
	fmt.Println(Q)
	fmt.Println()
	*/

	for Q.Len() > 0{

		q := heap.Pop(Q)
		if q == nil {
			break
		}
		u := q.(*Node)

		//fmt.Println(u)

		if dist[u] == math.Inf(1) {
			break
		}

		for _, e := range u.edges {
			alt := dist[u] + e.weight
			if alt < dist[e.target]{
				Q.update(e.target, alt)
				prev[e.target] = u
			}
		}

	}

	S := Graph{}
	for u := g.Search(target); prev[u] != nil; u = prev[u].(*Node) {
		fmt.Println(u)
		// add node
		// link to node ahead (last iteration)
	}
	// add source node and link
	return S
}
