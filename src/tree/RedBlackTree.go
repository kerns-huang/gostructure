package tree

import "sync"

type Item string

type BRTreeNode struct {
	isRed bool
	item Item
	parent *BRTreeNode
	left *BRTreeNode
	right *BRTreeNode
}

type BRTree struct {
	root *BRTreeNode
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
		(*tree).root = &BRTreeNode{true,item,nil,nil,nil}
		tree.size=1
		return
	}
}

// 左旋操作,当前节点成为左节点
func (tree *BRTree) rotateLeft(node *BRTreeNode){
	 right := *node.right

	 if (*node).parent == nil{ //如果节点的父亲是空的，说明node是根节点
         tree.root = right
	 }

}
// 右旋操作，当前节点成为右节点
func (tree *BRTree) rotateRight(node BRTreeNode){

}