package list

import (
	"sync"
)

type Item int

type Node struct {
	pre  *Node
	data Item
	next *Node
}

type LinkedList struct {
	size int  //链表长度
	last *Node
	first *Node //跟节点
	lock sync.RWMutex //读写锁对象
}

func (list *LinkedList) Init() {
	(*list).size = 0
	(list).first = nil
}

func(list *LinkedList) addObj(t Item){
	node :=&Node{nil,obj,nil}
	(*list).add(node)
}

func (list *LinkedList) isEmpty() bool{
	return (*list).size==0
}

func (list *LinkedList) getSize() int  {
	return (*list).size
}

func (list *LinkedList) add(node *Node) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if (*list).size == 0 {
		(*list).first = node //如果没有元素，跟节点赋值
		(*list).last=node
	}else{
		(*list).last.next=node
		node.pre=(*list).last
		(*list).last=node;
		(*list).size++
	}
}

func(list *LinkedList) remove(node *Node) {

}



