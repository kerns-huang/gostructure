package list

/**
 * 1:是数组对象，2：可以动态扩展，go里面的数组是切片，所以不需要在程序代码里面提现动态扩展的逻辑
 */
type ArrayList struct {
	elementData []interface{}
}
//添加元素
func (list *ArrayList) Add(obj interface{}) ()  {
	list.elementData=append(list.elementData, obj)
}
//删除元素
func (list *ArrayList) Remove(obj interface{}) {


}

//是否为空
func (list *ArrayList) IsEmpty() bool{
	return len(list.elementData)==0
}
//是否包含
func(list *ArrayList) contains(obj interface{}) bool{
	return len(list.elementData)==0
}

