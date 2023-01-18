package queue

import (
	"testing"

	"github.com/pawutj/go-algorithm/binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	Queue := GetQueue()
	Node0 := binary_tree.GetNode(0)
	Node1 := binary_tree.GetNode(1)
	Node2 := binary_tree.GetNode(2)

	Queue.Enqueue(Node0)
	Queue.Enqueue(Node1)
	Queue.Enqueue(Node2)

	result0 := Queue.Dequeue()
	assert.Equal(t, 0, result0.Value)
	result1 := Queue.Dequeue()
	assert.Equal(t, 1, result1.Value)
	result2 := Queue.Dequeue()
	assert.Equal(t, 2, result2.Value)

}
