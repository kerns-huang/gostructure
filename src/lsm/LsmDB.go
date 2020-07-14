package lsm

import "sync"

//lsm 算法 ，目前levelDB 和RocksDb 的常用算法，适合的场景，写多，读少的场景，
//但按照目前的优化，其实读多也无所谓，只是类似数据库索引的场景，没建立一个索引，会额外需要多用一份空间，这也是需要考虑的事情
//lsm 算法比较适合的是机械硬盘场景的读写优化，固态硬盘下有效果，但是优势没有那么明显
// 参考 git@github.com:prologic/bitcask.git

type lsmDb struct {
	lock sync.RWMutex //读写锁
	path string
	curr *DataFile  // 当前文件地址
	datafiles map[int]DataFile	// 历史的写入文件地址
					            // 树结构，保存索引路径

}