package bfs

import (
	tree "github.com/pawutj/go-algorithm/binary_tree"
	queue "github.com/pawutj/go-algorithm/quese"
)

func Bfs() []int {

	Head := tree.GetTreeComplex()
	Queue := queue.GetQueue()
	Queue.Enqueue(Head)
	v := []int{}
	for Queue.HasValue() {
		n := Queue.Dequeue()
		if n != nil {
			if n.Left != nil {
				Queue.Enqueue(*n.Left)
			}
			if n.Right != nil {
				Queue.Enqueue(*n.Right)
			}
			v = append(v, n.Value)
		}
	}

	return v
}
