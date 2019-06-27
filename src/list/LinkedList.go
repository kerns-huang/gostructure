package list

import (
	"sync"
	"fmt"
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

func (list *LinkedList) Init() {
	(*list).size = 0
	(list).first = nil
}

func (list *LinkedList) addObj(t Item) {
	node := &Node{nil, t, nil}
	(*list).add(node)
}

func (list *LinkedList) isEmpty() bool {
	return (*list).size == 0
}

func (list *LinkedList) getSize() int {
	return (*list).size
}

func (list *LinkedList) add(node *Node) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if (*list).size == 0 {
		(*list).first = node //如果没有元素，跟节点赋值
		(*list).last = node
		(*list).size=1
	} else {
		(*list).last.next = node
		node.pre = (*list).last
		(*list).last = node;
		(*list).size++
	}
}

func (list *LinkedList) RemoveAt(i int) (*Item, error) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if i < 0 || i > list.size {
		return nil, fmt.Errorf("Index %d out of bonuds", i)
	}

	curNode := list.first
	preIndex := 0
	for preIndex < i-1 {
		preIndex++
		curNode = curNode.next
	}
	item := curNode.data
	curNode.next = curNode.next.next
	list.size--
	return &item, nil
}

func (list *LinkedList) toString(){
	curNode := list.first
	for{
		if curNode.next ==nil{
			break
		}
		println(curNode.data)
		curNode=curNode.next
	}
}
