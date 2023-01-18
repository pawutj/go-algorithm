package stack

import (
	"testing"

	"github.com/pawutj/go-algorithm/binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	Stack := GetStack()
	Node0 := binary_tree.GetNode(0)
	Node1 := binary_tree.GetNode(1)
	Node2 := binary_tree.GetNode(2)

	Stack.Push(Node0)
	Stack.Push(Node1)
	Stack.Push(Node2)

	result0 := Stack.Pop()
	assert.Equal(t, 2, result0.Value)
	result1 := Stack.Pop()
	assert.Equal(t, 1, result1.Value)
	result2 := Stack.Pop()
	assert.Equal(t, 0, result2.Value)

}
