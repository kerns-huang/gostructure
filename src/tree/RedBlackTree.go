package tree

import "sync"

type Item string

type RBTreeNode struct {
	isRed bool
	item Item
	parent,left,right *RBTreeNode
}

type BRTree struct {
	root *RBTreeNode
	size int
	lock sync.RWMutex

}

func NewBRTree() *BRTree{
	tree := new(BRTree)
	tree.size=0
	return tree
}

func (tree *BRTree) Add(item Item){
	//如果没有节点，就把新增的节点设置成跟节点
	if  (*tree).root == nil {
		(*tree).root = &RBTreeNode{true,item,nil,nil,nil}
		tree.size=1
		return
	}
}

// 左旋操作,当前节点成为左节点
func (node *RBTreeNode) rotateLeft() (*RBTreeNode,error){
	var root *RBTreeNode
	 grandson :=node.right.left
	 node.right.left=node
	 node.parent=node.right
	 node.right=grandson
	 return root,nil

}
// 右旋操作，当前节点成为右节点
func (tree *BRTree) rotateRight(node RBTreeNode){

}