package tree

import (
	"errors"
	"fmt"
)

/**
 * b+ 树
 *一颗m阶的B树定义如下：
 * 1）每个结点最多有m-1个关键字。
 * 2）根结点最少可以只有1个关键字。
 * 3）非根结点至少有Math.ceil(m/2)-1个关键字。
 * 4）每个结点中的关键字都按照从小到大的顺序排列，每个关键字的左子树中的所有关键字都小于它，而右子树中的所有关键字都大于它。
 *  5）所有叶子结点都位于同一层，或者说根结点到每个叶子结点的长度都相同。
 */
var (
	err error
	defaultOrder = 4
	minOrder     = 3
	maxOrder     = 20
	order          = defaultOrder
	queue          *BPlugTreeNode
	verbose_output = false
	version        = 0.1
)
type BPlugTree struct {
	Root *BPlugTreeNode
}
 type BPlugTreeNode struct {
 	Pointers []interface{}  //指向数据的指针
 	Keys []int //key
 	NumKeys int //拥有多少数据
 	Parent *BPlugTreeNode //父亲节点
 	Children []*BPlugTreeNode  // 子节点，但如果是从底部
 	IsLeaf bool  //是否是叶子节点
 	Next *BPlugTreeNode //兄弟节点
 }

type Record struct {
	Value []byte
}

func NewBPlugTree() *BPlugTree{
	return &BPlugTree{}
}
//插入数据
func (t *BPlugTree) Insert(key int, value []byte) error {
	var pointer *Record
	var leaf *BPlugTreeNode

	if _, err := t.Find(key, false); err == nil {
		return errors.New("key already exists")
	}

	pointer, err := makeRecord(value)
	if err != nil {
		return err
	}

	if t.Root == nil { //没有数据的情况下，构建树接口
		return t.startNewTree(key, pointer)
	}

	leaf = t.findLeaf(key, false)

	if leaf.NumKeys < order-1 {
		insertIntoLeaf(leaf, key, pointer)
		return nil
	}

	return t.insertIntoLeafAfterSplitting(leaf, key, pointer)
}
func cut(length int) int {
	if length%2 == 0 {
		return length / 2
	}

	return length/2 + 1
}
// 插入数据到叶子节点
func insertIntoLeaf(leaf *BPlugTreeNode, key int, pointer *Record) {
	var i, insertion_point int

	for insertion_point < leaf.NumKeys && leaf.Keys[insertion_point] < key {
		insertion_point += 1
	}

	for i = leaf.NumKeys; i > insertion_point; i-- {
		leaf.Keys[i] = leaf.Keys[i-1]
		leaf.Pointers[i] = leaf.Pointers[i-1]
	}
	leaf.Keys[insertion_point] = key
	leaf.Pointers[insertion_point] = pointer
	leaf.NumKeys += 1
	return
}

func (t *BPlugTree) insertIntoLeafAfterSplitting(leaf *BPlugTreeNode, key int, pointer *Record) error {
	var new_leaf *BPlugTreeNode
	var insertion_index, split, new_key, i, j int
	var err error

	new_leaf, err = makeLeaf()
	if err != nil {
		return nil
	}

	temp_keys := make([]int, order)
	if temp_keys == nil {
		return errors.New("Error: Temporary keys array.")
	}

	temp_pointers := make([]interface{}, order)
	if temp_pointers == nil {
		return errors.New("Error: Temporary pointers array.")
	}

	for insertion_index < order-1 && leaf.Keys[insertion_index] < key {
		insertion_index += 1
	}

	for i = 0; i < leaf.NumKeys; i++ {
		if j == insertion_index {
			j += 1
		}
		temp_keys[j] = leaf.Keys[i]
		temp_pointers[j] = leaf.Pointers[i]
		j += 1
	}

	temp_keys[insertion_index] = key
	temp_pointers[insertion_index] = pointer

	leaf.NumKeys = 0

	split = cut(order - 1)

	for i = 0; i < split; i++ {
		leaf.Pointers[i] = temp_pointers[i]
		leaf.Keys[i] = temp_keys[i]
		leaf.NumKeys += 1
	}

	j = 0
	for i = split; i < order; i++ {
		new_leaf.Pointers[j] = temp_pointers[i]
		new_leaf.Keys[j] = temp_keys[i]
		new_leaf.NumKeys += 1
		j += 1
	}

	new_leaf.Pointers[order-1] = leaf.Pointers[order-1]
	leaf.Pointers[order-1] = new_leaf

	for i = leaf.NumKeys; i < order-1; i++ {
		leaf.Pointers[i] = nil
	}
	for i = new_leaf.NumKeys; i < order-1; i++ {
		new_leaf.Pointers[i] = nil
	}

	new_leaf.Parent = leaf.Parent
	new_key = new_leaf.Keys[0]

	return t.insertIntoParent(leaf, new_key, new_leaf)
}

func (t *BPlugTree) insertIntoNewRoot(left *BPlugTreeNode, key int, right *BPlugTreeNode) error {
	t.Root, err = makeNode()
	if err != nil {
		return err
	}
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = left
	t.Root.Pointers[1] = right
	t.Root.NumKeys += 1
	t.Root.Parent = nil
	left.Parent = t.Root
	right.Parent = t.Root
	return nil
}

func getLeftIndex(parent, left *BPlugTreeNode) int {
	left_index := 0
	for left_index <= parent.NumKeys && parent.Pointers[left_index] != left {
		left_index += 1
	}
	return left_index
}


func insertIntoNode(n *BPlugTreeNode, left_index, key int, right *BPlugTreeNode) {
	var i int
	for i = n.NumKeys; i > left_index; i-- {
		n.Pointers[i+1] = n.Pointers[i]
		n.Keys[i] = n.Keys[i-1]
	}
	n.Pointers[left_index+1] = right
	n.Keys[left_index] = key
	n.NumKeys += 1
}

//在分割之后插入
func (t *BPlugTree) insertIntoNodeAfterSplitting(old_node *BPlugTreeNode, left_index, key int, right *BPlugTreeNode) error {
	var i, j, split, k_prime int
	var new_node, child *BPlugTreeNode
	var temp_keys []int
	var temp_pointers []interface{}
	var err error

	temp_pointers = make([]interface{}, order+1)
	if temp_pointers == nil {
		return errors.New("Error: Temporary pointers array for splitting nodes.")
	}

	temp_keys = make([]int, order)
	if temp_keys == nil {
		return errors.New("Error: Temporary keys array for splitting nodes.")
	}

	for i = 0; i < old_node.NumKeys+1; i++ {
		if j == left_index+1 {
			j += 1
		}
		temp_pointers[j] = old_node.Pointers[i]
		j += 1
	}

	j = 0
	for i = 0; i < old_node.NumKeys; i++ {
		if j == left_index {
			j += 1
		}
		temp_keys[j] = old_node.Keys[i]
		j += 1
	}

	temp_pointers[left_index+1] = right
	temp_keys[left_index] = key

	split = cut(order)
	new_node, err = makeNode()
	if err != nil {
		return err
	}
	old_node.NumKeys = 0
	for i = 0; i < split-1; i++ {
		old_node.Pointers[i] = temp_pointers[i]
		old_node.Keys[i] = temp_keys[i]
		old_node.NumKeys += 1
	}
	old_node.Pointers[i] = temp_pointers[i]
	k_prime = temp_keys[split-1]
	j = 0
	for i += 1; i < order; i++ {
		new_node.Pointers[j] = temp_pointers[i]
		new_node.Keys[j] = temp_keys[i]
		new_node.NumKeys += 1
		j += 1
	}
	new_node.Pointers[j] = temp_pointers[i]
	new_node.Parent = old_node.Parent
	for i = 0; i <= new_node.NumKeys; i++ {
		child, _ = new_node.Pointers[i].(*BPlugTreeNode)
		child.Parent = new_node
	}

	return t.insertIntoParent(old_node, k_prime, new_node)
}

func (t *BPlugTree) insertIntoParent(left *BPlugTreeNode, key int, right *BPlugTreeNode) error {
	var left_index int
	parent := left.Parent

	if parent == nil {
		return t.insertIntoNewRoot(left, key, right)
	}
	left_index = getLeftIndex(parent, left)
	if parent.NumKeys < order-1 {
		insertIntoNode(parent, left_index, key, right)
		return nil
	}
	return t.insertIntoNodeAfterSplitting(parent, left_index, key, right)
}
//构建新的树
func (t *BPlugTree) startNewTree(key int, pointer *Record) error {
	t.Root, err = makeLeaf()
	if err != nil {
		return err
	}
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = pointer
	t.Root.Pointers[order-1] = nil
	t.Root.Parent = nil
	t.Root.NumKeys += 1
	return nil
}

func makeLeaf() (node *BPlugTreeNode, e error) {
	leaf, err := makeNode()
	if err != nil {
		return nil, err
	}
	leaf.IsLeaf = true
	return leaf, nil
}

func makeNode() (node *BPlugTreeNode,e error) {
	new_node :=new(BPlugTreeNode)
	if new_node == nil {
		return nil, errors.New("Error: Node creation.")
	}
	new_node.Keys = make([]int, order-1)
	if new_node.Keys == nil {
		return nil, errors.New("Error: New node keys array.")
	}
	new_node.Pointers = make([]interface{}, order)
	if new_node.Keys == nil {
		return nil, errors.New("Error: New node pointers array.")
	}
	new_node.IsLeaf = false
	new_node.NumKeys = 0
	new_node.Parent = nil
	new_node.Next = nil
	return new_node, nil
}

//创建值对象
func makeRecord(value []byte) (*Record, error) {
	new_record := new(Record)
	if new_record == nil {
		return nil, errors.New("Error: Record creation.")
	} else {
		new_record.Value = value
	}
	return new_record, nil
}

func (t *BPlugTree) Find(key int, verbose bool) (*Record, error) {
	i := 0
	c := t.findLeaf(key, verbose)
	if c == nil {
		return nil, errors.New("key not found")
	}
	for i = 0; i < c.NumKeys; i++ {
		if c.Keys[i] == key {
			break
		}
	}
	if i == c.NumKeys {
		return nil, errors.New("key not found")
	}

	r, _ := c.Pointers[i].(*Record)

	return r, nil
}

//查找叶子节点
func (t *BPlugTree) findLeaf(key int, verbose bool) *BPlugTreeNode {
	i := 0
	c := t.Root
	if c == nil {  //如果根节点为空，没有叶子节点
		if verbose {
			fmt.Printf("Empty tree.\n")
		}
		return c
	}
	for !c.IsLeaf {//如果c不是叶子节点，查找key索引
		i = 0
		for i < c.NumKeys {
			if key >= c.Keys[i] {
				i += 1
			} else {
				break
			}
		}
		c, _ = c.Pointers[i].(*BPlugTreeNode)  //按照理解，这是一个内部迭代的逻辑，但point值执行子节点or data ？
	}
	return c
}