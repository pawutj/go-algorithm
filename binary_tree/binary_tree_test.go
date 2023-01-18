package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTree(t *testing.T) {

	treeNode := GetTree()

	assert.Equal(t, treeNode.Value, 1)
	assert.Equal(t, treeNode.Left.Value, 2)
	assert.Equal(t, treeNode.Right.Value, 3)
}

func TestCreateTreeComplex(t *testing.T) {

	treeNode := GetTreeComplex()

	assert.Equal(t, treeNode.Value, 1)
	assert.Equal(t, treeNode.Left.Value, 2)
	assert.Equal(t, treeNode.Right.Value, 3)
	assert.Equal(t, treeNode.Left.Left.Value, 4)
	assert.Equal(t, treeNode.Left.Right.Value, 5)
}
