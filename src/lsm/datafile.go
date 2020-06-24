package lsm

import (
	"bufio"
	"encoding/binary"
	"os"
)

// 文件系统的读写，go语言里面的接口，不需要明显的implement，只需要实现方法就是继承。
type DataFile interface {
	FileId() int64     //文件ID
	Name() string      //文件名字
	Close() error      //关闭文件流
	Write(entry Entry) int64 // 写入对象,返回写入的字节流长度
}
type datafile struct {
	id     int64
	name   string   //文件名字
	file   *os.File //操作为文件
	offset int64    //文件的偏移量
}

const (
	keySize      = 4
	valueSize    = 8
	checksumSize = 4
)

//把对象写入文件系统
func (df *datafile) Write(e Entry) int64 {
	w := bufio.NewWriter(df.file)
	var bufKeyValue = make([]byte, keySize+valueSize)
	//写入key和value的长度
	binary.BigEndian.PutUint32(bufKeyValue[:keySize], uint32(len(e.Key)))
	binary.BigEndian.PutUint64(bufKeyValue[keySize:keySize+valueSize], uint64(len(e.Value)))
	w.Write(bufKeyValue)
	w.Write(e.Key)
	w.Write(e.Value)
	//分配多少字节
	bufChecksumSize := bufKeyValue[:checksumSize]
	//给字节赋值
	binary.BigEndian.PutUint32(bufChecksumSize, e.Checksum)
	w.Write(bufKeyValue)
	w.Flush()
	return int64(keySize+valueSize+len(e.Key)+len(e.Value)+checksumSize);
}

func (df *datafile) Close() error {
	//TODO 注意多线程的问题
	return df.file.Close()
}

func (df *datafile) Name() string {
	return df.name
}

func (df *datafile) FileId() int64 {
	return df.id
}
