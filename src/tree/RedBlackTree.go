package tree

import "sync"

type Item string

type RBTreeNode struct {
	isRed               bool
	item                Item
	parent, left, right *RBTreeNode
}

type RBTree struct {
	root *RBTreeNode
	size int
	lock sync.RWMutex
}

func New() *RBTree {
	tree := new(RBTree)
	tree.size = 0
	return tree
}

func (tree *RBTree) Add(item Item) {
	//如果没有节点，就把新增的节点设置成跟节点
	if (*tree).root == nil {
		(*tree).root = &RBTreeNode{true, item, nil, nil, nil}
		tree.size = 1
		return
	}
}

/*
	* 对红黑树的节点(x)进行左旋转，意味着x做为左节点存在
	*
	* 左旋示意图(对节点x进行左旋)：
	*      px                              px
	*     /                               /
	*    x                               y
	*   /  \      --(左旋)-.           / \                #
	*  lx   y                          x  ry
	*     /   \                       /  \
	*    ly   ry                     lx  ly
	*
	*
	*/
func (tree *RBTree) rotateLeft(x *RBTreeNode)  {
	y := x.right //新对象引用
	if y.left != nil {
		y.left.parent = x
	}
	if x.parent != nil {
		tree.root = y
	} else {
		y.parent = x.parent
		x.parent.left = y
	}
	y.left = x
	x.parent = y
	x.right = y.left
}

/*
	* 对红黑树的节点(y)进行右旋转，y 节点成为右节点
	*
	* 右旋示意图(对节点y进行左旋)：
	*            py                               py
	*           /                                /
	*          y                                x
	*         /  \      --(右旋)-.            /  \                     #
	*        x   ry                           lx   y
	*       / \                                   / \                   #
	*      lx  rx                                rx  ry
	*
	*/
// 右旋操作，当前节点成为右节点
func (tree *RBTree) rotateRight(y *RBTreeNode) () {
	x := y.left
	if y.parent == nil {
		tree.root = x
	} else {
		x.parent = y.parent
		y.parent.left = x
	}
	x.right = y
	y.parent = x
	if x.right != nil {
		y.left = x.right
		x.right.parent = y
	}
}

/*
 * 红黑树插入修正函数
 *
 * 在向红黑树中插入节点之后(失去平衡)，再调用该函数；
 * 目的是将它重新塑造成一颗红黑树。
 *
 * 参数说明：
 *     node 插入的结点
 */
func insertFixUp(node RBTreeNode) {

}
