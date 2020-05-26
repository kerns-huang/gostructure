package list

/**
 * 1:是数组对象，2：可以动态扩展，go里面的数组是切片，所以不需要在程序代码里面提现动态扩展的逻辑
 */
type ArrayList struct {
	elementData []interface{}
}

//添加元素
func (list *ArrayList) Add(obj interface{}) () {
	list.elementData = append(list.elementData, obj)
}

//删除元素，先查找index，然后通过index 删除数据
func (list *ArrayList) Remove(obj interface{}) {
	i := list.searchIndex(obj)
	// 重新赋值数值，这个比java的copy清晰多了
	if i > 0 {
		list.elementData = append(list.elementData[:i], list.elementData[i+1:])
	}
}
//计算array size的长度
func (list *ArrayList) Size() int {
	return len(list.elementData)
}

//查找搜索的对象在数组中的位置，如果不在数组中，直接返回-1,建议使用二分查找算法，数据量多的时候效果会快很多
func (list *ArrayList) searchIndex(obj interface{}) int {
	for i := 0; i < len(list.elementData); i++ {
		if list.elementData[i] == obj {
			return i
		}
	}
	return -1
}

//是否为空
func (list *ArrayList) IsEmpty() bool {
	return len(list.elementData) == 0
}

//是否包含其实是一个查找函数，其实可以通过二分查找去实现数据的查找
func (list *ArrayList) contains(obj interface{}) bool {
	if list.searchIndex(obj) > 0 {
		return true
	}
	return false
}
