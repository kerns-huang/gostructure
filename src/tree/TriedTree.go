package tree

/**
 * 字典树 ，前缀树查询，flashtext的结构树，
 * 以空间换时间，利用字符串的公共前缀来减少字符串的比较
 */

//26个字母
const AlphabetSize = 26
//前缀树节点。
type trieNode struct {
	count int                        //
	children *[AlphabetSize]trieNode //  包含的子节点
}


type TrieTree struct {
	root *trieNode //包含的root节点
}

func createTrieNode() *trieNode  {
      node:=&trieNode{count: 0}
      return node
}
// 插入数据
func(tree *TrieTree) Insert(key string)  {

}

//查询数据，不存在，返回0 ，存在返回出现的次数
func(tree *TrieTree) Search(key string)  {

}