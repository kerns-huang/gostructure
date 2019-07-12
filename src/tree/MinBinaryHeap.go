package tree

/**
 *堆（英语：Heap）是计算机科学中的一种特别的树状数据结构。若是满足以下特性，即可称为堆：“给定堆中任意节点P和C，若P是C的母节点，那么P的值会小于等于（或大于等于）C的值”。
 *若母节点的值恒小于等于子节点的值，此堆称为最小堆（min heap）；
 *反之，若母节点的值恒大于等于子节点的值，此堆称为最大堆（max heap）。在堆中最顶端的那一个节点，称作根节点（root node），根节点本身没有母节点（parent node）。
 *
 *
 */

type MinBinaryHeap struct {
	items []Item
}

func (heap *MinBinaryHeap) push(item Item) *MinBinaryHeap{
	heap.items = append(heap.items, item)
	datas:=heap.heapup(heap.items, len(heap.items))
	heap.items=datas
	return heap

}

func (heap *MinBinaryHeap) pop() *Item {
	item := heap.items[len(heap.items)-1]
	heap.items = heap.items[:len(heap.items)-1]
	return &item

}

/**
 * 上浮调整
 */
func (heap *MinBinaryHeap) heapup(datas []Item, index int) []Item {
	if index > 1 {
		parent := index / 2
		parentValue := datas[parent-1]
		indexValue := datas[index-1]
		if parentValue.Compare(indexValue) == 1 {
			tmp := datas[parent-1]
			datas[parent-1] = datas[index-1]
			datas[index-1] = tmp
			heap.heapup(datas, parent)
		} else {
			//没有发生交换，说明新增的数据已经找到它的位置了
			return datas
		}
	}
	return datas
}
