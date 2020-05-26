package tree

import (
	"errors"
	"reflect"
)

/**
 *  b+ 树
 * 一颗m阶的B树定义如下：
 * 1）每个结点最多有m-1个关键字。
 * 2）根结点最少可以只有1个关键字。
 * 3）非根结点至少有Math.ceil(m/2)-1个关键字。
 * 4）每个结点中的关键字都按照从小到大的顺序排列，每个关键字的左子树中的所有关键字都小于它，而右子树中的所有关键字都大于它。
 *  5）所有叶子结点都位于同一层，或者说根结点到每个叶子结点的长度都相同。
 */
var (
	err            error
	defaultOrder   = 4
	minOrder       = 3
	maxOrder       = 20
	order          = defaultOrder //阶数是m，则除了根之外的每个节点都包含最少 m/2 个元素最多 m 个元素，
	queue          *BPlugTreeNode
	verbose_output = false
	version        = 0.1
)

//b+树结构
type BPlugTree struct {
	Root *BPlugTreeNode
}

// 树节点结构
type BPlugTreeNode struct {
	Pointers []interface{}    //存储的数据
	Keys     []int            //key
	NumKeys  int              //拥有多少数据
	Parent   *BPlugTreeNode   //父亲节点指针
	Children []*BPlugTreeNode // 子节点指针，但如果是从底部
	IsLeaf   bool             //是否是叶子节点
	Next     *BPlugTreeNode   //兄弟节点
}

type Record struct {
	Value interface{}
}

func NewBPlugTree() *BPlugTree {
	return &BPlugTree{}
}

//插入数据
func (t *BPlugTree) Insert(key int, value interface{}) error {
	var pointer *Record
	var leaf *BPlugTreeNode
	//检查key是否存在
	if _, err := t.Find(key); err == nil {
		return errors.New("key already exists")
	}
	pointer, err := makeRecord(value)
	if err != nil {
		return err
	}
	if t.Root == nil { //没有数据的情况下，构建树接口
		return t.startNewTree(key, pointer)
	}
	leaf = t.findLeaf(key)      //查找叶子节点，只有叶子节点才存放数据
	if leaf.NumKeys < order-1 { //叶子节点拥有的节点
		insertIntoLeaf(leaf, key, pointer)
		return nil
	}
	return t.insertIntoLeafAfterSplitting(leaf, key, pointer)
}

//删除key
func (t *BPlugTree) Delete(key int) error {
	key_record, err := t.Find(key)
	if err != nil {
		return err
	}
	key_leaf := t.findLeaf(key)
	if key_record != nil && key_leaf != nil {
		t.deleteEntry(key_leaf, key, key_record)
	}
	return nil
}

func (t *BPlugTree) deleteEntry(n *BPlugTreeNode, key int, pointer interface{}) {
	var min_keys, neighbour_index, k_prime_index, k_prime, capacity int
	var neighbour *BPlugTreeNode
	n = removeEntryFromNode(n, key, pointer)
	if n == t.Root {
		t.adjustRoot()
		return
	}
	if n.IsLeaf {
		min_keys = cut(order - 1)
	} else {
		min_keys = cut(order) - 1
	}
	if n.NumKeys >= min_keys {
		return
	}
	neighbour_index = getNeighbourIndex(n)
	if neighbour_index == -1 {
		k_prime_index = 0
	} else {
		k_prime_index = neighbour_index
	}
	k_prime = n.Parent.Keys[k_prime_index]
	if neighbour_index == -1 {
		neighbour, _ = n.Parent.Pointers[1].(*BPlugTreeNode)
	} else {
		neighbour, _ = n.Parent.Pointers[neighbour_index].(*BPlugTreeNode)
	}

	if n.IsLeaf {
		capacity = order
	} else {
		capacity = order - 1
	}

	if neighbour.NumKeys+n.NumKeys < capacity {
		t.coalesceNodes(n, neighbour, neighbour_index, k_prime)
		return
	} else {
		t.redistributeNodes(n, neighbour, neighbour_index, k_prime_index, k_prime)
		return
	}

}

func (t *BPlugTree) redistributeNodes(n, neighbour *BPlugTreeNode, neighbour_index, k_prime_index, k_prime int) {
	var i int
	var tmp *BPlugTreeNode

	if neighbour_index != -1 {
		if !n.IsLeaf {
			n.Pointers[n.NumKeys+1] = n.Pointers[n.NumKeys]
		}
		for i = n.NumKeys; i > 0; i-- {
			n.Keys[i] = n.Keys[i-1]
			n.Pointers[i] = n.Pointers[i-1]
		}
		if !n.IsLeaf { // why the second if !n.IsLeaf
			n.Pointers[0] = neighbour.Pointers[neighbour.NumKeys]
			tmp, _ = n.Pointers[0].(*BPlugTreeNode)
			tmp.Parent = n
			neighbour.Pointers[neighbour.NumKeys] = nil
			n.Keys[0] = k_prime
			n.Parent.Keys[k_prime_index] = neighbour.Keys[neighbour.NumKeys-1]
		} else {
			n.Pointers[0] = neighbour.Pointers[neighbour.NumKeys-1]
			neighbour.Pointers[neighbour.NumKeys-1] = nil
			n.Keys[0] = neighbour.Keys[neighbour.NumKeys-1]
			n.Parent.Keys[k_prime_index] = n.Keys[0]
		}
	} else {
		if n.IsLeaf {
			n.Keys[n.NumKeys] = neighbour.Keys[0]
			n.Pointers[n.NumKeys] = neighbour.Pointers[0]
			n.Parent.Keys[k_prime_index] = neighbour.Keys[1]
		} else {
			n.Keys[n.NumKeys] = k_prime
			n.Pointers[n.NumKeys+1] = neighbour.Pointers[0]
			tmp, _ = n.Pointers[n.NumKeys+1].(*BPlugTreeNode)
			tmp.Parent = n
			n.Parent.Keys[k_prime_index] = neighbour.Keys[0]
		}
		for i = 0; i < neighbour.NumKeys-1; i++ {
			neighbour.Keys[i] = neighbour.Keys[i+1]
			neighbour.Pointers[i] = neighbour.Pointers[i+1]
		}
		if !n.IsLeaf {
			neighbour.Pointers[i] = neighbour.Pointers[i+1]
		}
	}
	n.NumKeys += 1
	neighbour.NumKeys -= 1
	return
}
func (t *BPlugTree) coalesceNodes(n, neighbour *BPlugTreeNode, neighbour_index, k_prime int) {
	var i, j, neighbour_insertion_index, n_end int
	var tmp *BPlugTreeNode

	if neighbour_index == -1 {
		tmp = n
		n = neighbour
		neighbour = tmp
	}

	neighbour_insertion_index = neighbour.NumKeys

	if !n.IsLeaf {
		neighbour.Keys[neighbour_insertion_index] = k_prime
		neighbour.NumKeys += 1

		n_end = n.NumKeys
		i = neighbour_insertion_index + 1
		for j = 0; j < n_end; j++ {
			neighbour.Keys[i] = n.Keys[j]
			neighbour.Pointers[i] = n.Pointers[j]
			neighbour.NumKeys += 1
			n.NumKeys -= 1
			i += 1
		}
		neighbour.Pointers[i] = n.Pointers[j]

		for i = 0; i < neighbour.NumKeys+1; i++ {
			tmp, _ = neighbour.Pointers[i].(*BPlugTreeNode)
			tmp.Parent = neighbour
		}
	} else {
		i = neighbour_insertion_index
		for j = 0; j < n.NumKeys; j++ {
			neighbour.Keys[i] = n.Keys[j]
			n.Pointers[i] = n.Pointers[j]
			neighbour.NumKeys += 1
		}
		neighbour.Pointers[order-1] = n.Pointers[order-1]
	}

	t.deleteEntry(n.Parent, k_prime, n)
}
func getNeighbourIndex(n *BPlugTreeNode) int {
	var i int

	for i = 0; i <= n.Parent.NumKeys; i++ {
		if reflect.DeepEqual(n.Parent.Pointers[i], n) {
			return i - 1
		}
	}

	return i
}
func (t *BPlugTree) adjustRoot() {
	var new_root *BPlugTreeNode

	if t.Root.NumKeys > 0 {
		return
	}

	if !t.Root.IsLeaf {
		new_root, _ = t.Root.Pointers[0].(*BPlugTreeNode)
		new_root.Parent = nil
	} else {
		new_root = nil
	}
	t.Root = new_root

	return
}

func removeEntryFromNode(n *BPlugTreeNode, key int, pointer interface{}) *BPlugTreeNode {
	var i, num_pointers int

	for n.Keys[i] != key {
		i += 1
	}

	for i += 1; i < n.NumKeys; i++ {
		n.Keys[i-1] = n.Keys[i]
	}

	if n.IsLeaf {
		num_pointers = n.NumKeys
	} else {
		num_pointers = n.NumKeys + 1
	}

	i = 0
	for n.Pointers[i] != pointer {
		i += 1
	}
	for i += 1; i < num_pointers; i++ {
		n.Pointers[i-1] = n.Pointers[i]
	}
	n.NumKeys -= 1

	if n.IsLeaf {
		for i = n.NumKeys; i < order-1; i++ {
			n.Pointers[i] = nil
		}
	} else {
		for i = n.NumKeys + 1; i < order; i++ {
			n.Pointers[i] = nil
		}
	}

	return n
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
	var new_leaf *BPlugTreeNode //新的叶子节点
	var insertion_index, split, new_key, i, j int
	var err error

	new_leaf, err = makeLeaf()
	if err != nil {
		return nil
	}

	temp_keys := make([]int, order) //定义切片
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

//构建叶子节点
func makeLeaf() (node *BPlugTreeNode, e error) {
	leaf, err := makeNode()
	if err != nil {
		return nil, err
	}
	leaf.IsLeaf = true
	return leaf, nil
}

//创建一个节点数据
func makeNode() (node *BPlugTreeNode, e error) {
	new_node := new(BPlugTreeNode)
	if new_node == nil {
		return nil, errors.New("Error: Node creation.")
	}
	new_node.Keys = make([]int, order) //创建的数组不是和深度有关系 ？
	new_node.Pointers = make([]interface{}, order)
	new_node.IsLeaf = false
	new_node.NumKeys = 0
	new_node.Parent = nil
	new_node.Next = nil
	return new_node, nil
}

//创建值对象
func makeRecord(value interface{}) (*Record, error) {
	new_record := new(Record)
	if new_record == nil {
		return nil, errors.New("Error: Record creation.")
	} else {
		new_record.Value = value
	}
	return new_record, nil
}

//查找数据
func (t *BPlugTree) Find(key int) (*Record, error) {
	i := 0
	c := t.findLeaf(key) //从叶子节点开始查找
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
func (t *BPlugTree) findLeaf(key int) *BPlugTreeNode {
	i := 0
	node := t.Root
	if node == nil { //如果根节点为空，没有叶子节点
		return node
	}
	for !node.IsLeaf { //如果c不是叶子节点，查找key索引
		i = node.findIndex(key)
		node, _ = node.Pointers[i].(*BPlugTreeNode) //查找节点信息
	}
	return node
}

func (node *BPlugTreeNode) findIndex(Key int) int {
	 low := 0
	 high := node.NumKeys - 1
	// the cached index minus one, so that
	// for the first time (when cachedCompare is 0),
	// the default value is used
	 x := 0
	if x < 0 || x > high {
		x = high >> 1
	}
	for low <= high {
		if Key > node.Keys[x] {
			low = x + 1
		} else if Key < node.Keys[x] {
			high = x - 1
		} else {
			return x
		}
		x = (low + high) >> 1
	}
	return -1
}
