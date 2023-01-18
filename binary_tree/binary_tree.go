package binary_tree

type TreeNode struct {
	value int
	Left  *TreeNode
	Right *TreeNode
}

func GetTree() TreeNode {

	treeNode1 := TreeNode{1, nil, nil}
	treeNode2 := TreeNode{2, nil, nil}
	treeNode3 := TreeNode{3, nil, nil}

	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3

	return treeNode1
}

func GetTreeComplex() TreeNode {
	treeNode1 := TreeNode{1, nil, nil}
	treeNode2 := TreeNode{2, nil, nil}
	treeNode3 := TreeNode{3, nil, nil}
	treeNode4 := TreeNode{4, nil, nil}
	treeNode5 := TreeNode{5, nil, nil}
	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode1.Left.Left = &treeNode4
	treeNode1.Left.Right = &treeNode5

	return treeNode1
}
