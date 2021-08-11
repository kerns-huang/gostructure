package lsm

import "gostructure/src/tree"

type memorySegment struct {
   tree *tree.RBTree
}

func newMemorySegment()  segment{
	segment:= new(memorySegment)
	segment.tree=tree.NewRBTree()
	return segment
}

// Put 保存数据
func(segment *memorySegment) Put(key []byte, value []byte) error{
	return nil
}
// Get 获取数据
func(segment *memorySegment) Get(key []byte) ([]byte, error){
	return nil,nil
}
// Remove 删除数据
func(segment *memorySegment) Remove(key []byte) error{
	return nil
}
//Lookup 查找返回数据
func(segment *memorySegment) Lookup(lower []byte, upper []byte) (LookupIterator, error){
	return nil,nil
}
// Close 关闭片段
func(segment *memorySegment) Close() error{
	return nil
}

