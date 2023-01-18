package stack

import (
	"github.com/pawutj/go-algorithm/binary_tree"
)

type Stack struct {
	node []binary_tree.TreeNode
}

func GetStack() Stack {
	return Stack{[]binary_tree.TreeNode{}}
}

func (s *Stack) Push(n binary_tree.TreeNode) {
	s.node = append(s.node, n)
}

func (s *Stack) Pop() binary_tree.TreeNode {
	temp := s.node[len(s.node)-1]
	s.node = s.node[:len(s.node)-1]
	return temp
}
