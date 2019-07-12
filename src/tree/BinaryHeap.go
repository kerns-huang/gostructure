package tree
/**
 *堆（英语：Heap）是计算机科学中的一种特别的树状数据结构。若是满足以下特性，即可称为堆：“给定堆中任意节点P和C，若P是C的母节点，那么P的值会小于等于（或大于等于）C的值”。
 *若母节点的值恒小于等于子节点的值，此堆称为最小堆（min heap）；
 *反之，若母节点的值恒大于等于子节点的值，此堆称为最大堆（max heap）。在堆中最顶端的那一个节点，称作根节点（root node），根节点本身没有母节点（parent node）。
 *
 *
 */

 type BHNode struct {
 	item Item
 	parent,left,right *BHNode
 }
 type BinaryHeap struct {
 	items []Item
 }

func(heap *BinaryHeap) Insert(item Item){
	heap.items= append(heap.items, item)

}