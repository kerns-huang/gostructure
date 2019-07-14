package tree

/**
 *堆（英语：Heap）是计算机科学中的一种特别的树状数据结构。若是满足以下特性，即可称为堆：“给定堆中任意节点P和C，若P是C的母节点，那么P的值会小于等于（或大于等于）C的值”。
 *若母节点的值恒小于等于子节点的值，此堆称为最小堆（min heap）；
 *反之，若母节点的值恒大于等于子节点的值，此堆称为最大堆（max heap）。在堆中最顶端的那一个节点，称作根节点（root node），根节点本身没有母节点（parent node）。
 *                          11
 *			            /        \
 *		               9           10
 *		            /     \      /     \
 *	              5      6     7      8
 *               / \     / \
 *              1   2   3   4
 * 最大二叉堆一般使用数组来存储数据
 *  [11,10,9,8,7,6,5,4,3,2,1] ,
 *  对于一个很大的堆，这种存储是低效的。因为节点的子节点很可能在另外一个内存页
 *  extra: linux 内存页大小获取  getconf PAGE_SIZE
 *  如果存储数组的下标基于0，那么下标为i的节点的子节点是2i + 1与2i + 2；其父节点的下标是⌊floor((i − 1) ∕ 2)⌋
 *  函数floor(x)的功能是“向下取整”，或者说“向下舍入”
 */

type MinBinaryHeap struct {
	items []Item
}

func (heap *MinBinaryHeap) Push(item Item) *MinBinaryHeap{
	heap.items = append(heap.items, item)
	datas:=heap.heapup(heap.items, len(heap.items))
	heap.items=datas
	return heap

}

func (heap *MinBinaryHeap) Pop() *Item {
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
