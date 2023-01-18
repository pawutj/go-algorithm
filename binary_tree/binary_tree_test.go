package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTree(t *testing.T) {

	treeNode := GetTree()

	assert.Equal(t, treeNode.value, 1)
	assert.Equal(t, treeNode.Left.value, 2)
	assert.Equal(t, treeNode.Right.value, 3)
}

func TestCreateTreeComplex(t *testing.T) {

	treeNode := GetTreeComplex()

	assert.Equal(t, treeNode.value, 1)
	assert.Equal(t, treeNode.Left.value, 2)
	assert.Equal(t, treeNode.Right.value, 3)
	assert.Equal(t, treeNode.Left.Left.value, 4)
	assert.Equal(t, treeNode.Left.Right.value, 5)
}
