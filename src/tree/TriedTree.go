package tree

/**
 * 字典树 ，前缀树查询，flashtext的结构树，
 * 以空间换时间，利用字符串的公共前缀来减少字符串的比较
 * 目前实现的是简单的26个字母的字典树，如果是中文的字典树，该如何实现，长度就是不固定的一个状态了，用切面的格式 ？
 */

//26个字母
const AlphabetSize = 26

//前缀树节点。
type trieNode struct {
	count    int                     //
	isKey    bool                    //是否是一个完整的关键词
	children [AlphabetSize]*trieNode //  包含的子节点
}

type TrieTree struct {
	root *trieNode //包含的root节点
}

func createTrieNode() *trieNode {
	node := &trieNode{count: 0}
	return node
}

// 新创建一个字典树
func NewTrieTree() *TrieTree {
	return &TrieTree{root: createTrieNode()}
}

// 插入数据
func (tree *TrieTree) Insert(key string) {
	//按照字典树规则插入节点，
	cur := tree.root
	for i, _ := range key {
		if cur.children[key[i]-'a'] == nil {
			cur.children[key[i]-'a'] = createTrieNode()
		}
		cur = cur.children[key[i]-'a']
	}
	if !cur.isKey {
		//设置为是完整的词
		cur.isKey = true
	}
	//访问次数+1
	cur.count++
}

//查询数据，不存在，返回0 ，存在返回出现的次数
func (tree *TrieTree) Search(key string) bool {
	cur := tree.root
	for i, _ := range key {
		if cur.children[key[i]-'a'] == nil {
			return false
		}
		cur = cur.children[key[i]-'a']
	}
	if !cur.isKey {
		return false
	}
	return true
}


