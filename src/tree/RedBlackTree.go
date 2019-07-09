package tree

import "sync"

type Item interface {
	Compare(o Item) int
}

type RBTreeNode struct {
	isRed               bool
	item                Item
	parent, left, right *RBTreeNode
}

const (
	RED bool= true

	BLACK bool= false
)

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

func Compare(o1, o2 Item) int {
	return o1.Compare(o2)
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
func (tree *RBTree) rotateLeft(x *RBTreeNode) {
	if x.right == nil {
		return
	}
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
	if y.left == nil {
		return
	}
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
 * 从root开始比较，如果比左节点大
 */
func (tree *RBTree) insert(node *RBTreeNode) *RBTreeNode {
	if tree.root == nil{
		tree.root=node
		tree.size++
		return node
	}
	no1 := tree.root
	no2 := tree.root
	for no1 != nil {
		no2 = no1
		if node.item.Compare(no1.item) == 1 {
			no1 = no1.right
		} else {
			no1 = no1.left
		}
	}

	if node.item.Compare(no2.item) ==1 {
		no2.right=node

	}else{
		no2.left=node
	}
	node.parent=no2
	tree.insertFixUp(node) //节点修复
	tree.size++
	return nil
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
func (tree *RBTree)  insertFixUp(node *RBTreeNode) {
	for node.parent.isRed { //如果是红节点，需要旋转，调色，重新设置成平衡。
		//
		// Howerver, we do not need the assertion of non-nil grandparent
		// because
		//
		//  2) The root is black
		//
		// Since the color of the parent is RED, so the parent is not root
		// and the grandparent must be exist.
		//
		if node.parent == node.parent.parent.left {
			// Take y as the uncle, although it can be NIL, in that case
			// its color is BLACK
			y := node.parent.parent.right
			if y.isRed {
				//
				// Case 1:
				// Parent and uncle are both RED, the grandparent must be BLACK
				// due to
				//
				//  4) Both children of every red node are black
				//
				// Since the current node and its parent are all RED, we still
				// in violation of 4), So repaint both the parent and the uncle
				// to BLACK and grandparent to RED(to maintain 5)
				//
				//  5) Every simple path from root to leaves contains the same
				//     number of black nodes.
				//
				node.parent.isRed = false
				y.isRed = false
				node.parent.parent.isRed = true
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					//
					// Case 2:
					// Parent is RED and uncle is BLACK and the current node
					// is right child
					//
					// A left rotation on the parent of the current node will
					// switch the roles of each other. This still leaves us in
					// violation of 4).
					// The continuation into Case 3 will fix that.
					//
					node = node.parent
					tree.rotateLeft(node)
				}
				//
				// Case 3:
				// Parent is RED and uncle is BLACK and the current node is
				// left child
				//
				// At the very beginning of Case 3, current node and parent are
				// both RED, thus we violate 4).
				// Repaint parent to BLACK will fix it, but 5) does not allow
				// this because all paths that go through the parent will get
				// 1 more black node. Then repaint grandparent to RED (as we
				// discussed before, the grandparent is BLACK) and do a right
				// rotation will fix that.
				//
				node.parent.isRed = false
				node.parent.parent.isRed = true
				tree.rotateRight(node.parent.parent)
			}
		} else { // same as then clause with "right" and "left" exchanged
			y := node.parent.parent.left
			if y.isRed {
				node.parent.isRed = false
				y.isRed = false
				node.parent.parent.isRed = true
				node = node.parent.parent

			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rotateRight(node)
				}
				node.parent.isRed = false
				node.parent.parent.isRed = true
				tree.rotateLeft(node.parent.parent)
			}
		}
	}
	tree.root.isRed = false
}
