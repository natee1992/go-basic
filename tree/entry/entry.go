package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.Print()
	root.Right.Left.SetValue(4)
	fmt.Print("\n")
	root.Right.Left.Print()
	root.Print()
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(300)
	root.Print()
	root.Traverse()
	fmt.Print("\n")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
