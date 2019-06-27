package list

import (
	"fmt"
	"sync"
)

type Item int

type Node struct {
	pre  *Node
	data Item
	next *Node
}

type LinkedList struct {
	size  int //链表长度
	last  *Node
	first *Node        //跟节点
	lock  sync.RWMutex //读写锁对象
}
//新生成一个LinkedList
func NewLinkedList() *LinkedList{
	list :=new(LinkedList)
	list.Init()
	return list
}

func (list *LinkedList) Init() {
	(*list).size = 0
	(list).first = nil
}

func (list *LinkedList) Add(t Item) {
	node := &Node{nil, t, nil}
	(*list).add(node)
}

func (list *LinkedList) Empty() bool {
	return (*list).size == 0
}

func (list *LinkedList) Size() int {
	return (*list).size
}


func (list *LinkedList) RemoveAt(i int) (*Item, error) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if i < 0 || i > list.size {
		return nil, fmt.Errorf("Index %d out of bonuds", i)
	}

	curNode := list.first
	preIndex := 0
	for preIndex < i {
		preIndex++
		curNode = curNode.next
	}
	item := curNode.data
	if list.first == curNode { //如果删除的头节点
		(*list).first = (*curNode).next
	} else if list.last == curNode { //如果是尾部节点
		curNode.pre.next = nil  //前节点的next节点设置为空
		list.last = curNode.pre //尾部节点设置为前一个节点
	} else { //中间节点移除
		curNode.pre.next = curNode.next //前节点的下一个节点是当前节点的子节点
		curNode.next.pre = curNode.pre
	}
	list.size--
	return &item, nil
}

func (list *LinkedList) Contains(vals ...Item) bool {
	ss := 0
	for _, v := range vals {
		curNode := list.first
		for {
			if curNode == nil {
				break
			}
			if (*curNode).data == v {
				ss++
				break
			}
			curNode = curNode.next
		}
	}
	return ss == len(vals)
}

func (list *LinkedList) ToString() {
	curNode := list.first
	for {
		if curNode != nil {
			println(curNode.data)
		}
		if curNode.next == nil {
			break
		}
		curNode = curNode.next
	}
}


func (list *LinkedList) add(node *Node) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if (*list).size == 0 {
		(*list).first = node //如果没有元素，跟节点赋值
		(*list).last = node
		(*list).size = 1
	} else {
		(*list).last.next = node
		node.pre = (*list).last
		(*list).last = node;
		(*list).size++
	}
}
