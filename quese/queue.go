package queue

import "github.com/pawutj/go-algorithm/binary_tree"

type Queue struct {
	node []binary_tree.TreeNode
}

func GetQueue() Queue {
	return Queue{[]binary_tree.TreeNode{}}
}

func (q *Queue) Enqueue(n binary_tree.TreeNode) {
	q.node = append(q.node, n)
}

func (q *Queue) Dequeue() binary_tree.TreeNode {
	result := q.node[0]
	q.node = q.node[1:]
	return result
}
