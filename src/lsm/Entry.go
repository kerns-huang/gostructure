package lsm

type Entry struct {
	Checksum uint32 //签名串
	Key []byte
	Value []byte
	Offset int64  //文件中的偏移量
}
