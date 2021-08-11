package lsm

// 片段数据
type segment interface {
	// Put 保存数据
	Put(key []byte, value []byte) error
	// Get 获取数据
	Get(key []byte) ([]byte, error)
	// Remove 删除数据
	Remove(key []byte) error
	//Lookup 查找返回数据
	Lookup(lower []byte, upper []byte) (LookupIterator, error)
	// Close 关闭片段
	Close() error
}

//LookupIterator 迭代器
type LookupIterator interface {

}
