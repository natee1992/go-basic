package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

/*
	自定义工厂函数返回了局部变量的地址
	c++ <局部变量分配到栈上，函数推出，局部变量会被销毁，如果要传出需要分配到堆上，但分配到堆上，需要手动释放>
	java <new---> 分配到堆上，垃圾回收机制>
	golang 编译器自行决定根据变量情况自动分配，垃圾回收机制
 */
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

/*
 为结构定义方法
*/

func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

/*
 值传递和引用指针传递
*/
func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

/*
	tree 的遍历<中序遍历>
*/
func (node *TreeNode) Traverse()  {
	if node == nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

