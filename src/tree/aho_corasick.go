package tree

import (
	"container/list"
)

/**
 * ac 自动机是tiredTree的进阶版本，引入fail 节点解决失败之后从上层节点查找数据的逻辑
 */
const childNodeCount = 16

// AC自动机节点结构定义
type AcAutoNode struct {
	endCount int            // 结束模式串个数
	prefixCount int            // 前缀模式串个数
	failNode *AcAutoNode    // fail指针节点
	childNode [childNodeCount]*AcAutoNode    // 子节点
}

// AC自动机初始化
var GAcAuto *AcAutoNode
func init(){
	GAcAuto = new(AcAutoNode)
}

// 添加配置
func BuildTree(s []string) {
	// 遍历模式串列表
	for uli := 0; uli < len(s); uli++ {
		node := GAcAuto
		// 遍历模式串字符
		for _, runeCh := range s[uli] {
			// 分高低位判断
			runeStr := string(runeCh)
			for ulj := 0; ulj < len(runeStr); ulj++ {
				// 取 整
				indexHigh := runeStr[ulj] / childNodeCount
				if node.childNode[indexHigh] == nil {
					node.childNode[indexHigh] = &AcAutoNode{}
				}
				node = node.childNode[indexHigh]
				// 取 余
				indexLow := runeStr[ulj] % childNodeCount
				if node.childNode[indexLow] == nil {
					node.childNode[indexLow] = &AcAutoNode{}
				}
				node = node.childNode[indexLow]
			}
			node.prefixCount++
		}

		node.endCount++
	}
}

//设置查询失败的指向节点
func SetNodeFailPoint() {
	GAcAuto.failNode = nil
    //初始化一个链表
	nodeList := list.New()
	nodeList.PushBack(GAcAuto)

	// 逐层遍历trie树节点，为节点设置fail指针
	for {
		length := nodeList.Len()
		if length <= 0 {
			break
		}

		for uli := 0; uli < length; uli++ {
			ele := nodeList.Front()
			node, ok := ele.Value.(*AcAutoNode)
			if ok {
				if node == GAcAuto {
					// 根节点的子节点的fail指针都指向根节点
					for ulj := 0; ulj < childNodeCount; ulj++ {
						if node.childNode[ulj] != nil {
							node.childNode[ulj].failNode = GAcAuto
						}
					}
				} else {
					// 其他节点的子节点的fail指针就看它父节点fail指针指向的节点的子节点情况
					for ulj := 0; ulj < childNodeCount; ulj++ {
						// 遍历所有非空的子节点，为其设置fail指针
						if node.childNode[ulj] != nil {
							// fail指针设置原则是：
							// 1）查看father->failNode下有没有和自己一样的子节点，有则fail指针取该子节点
							// 2）否则，沿father->failNode->failNode继续查询下，如果一直没有，fail指针就取根节点
							nextNode := node.failNode
							for {
								if nextNode == nil {
									node.childNode[ulj].failNode = GAcAuto
									break
								} else {
									if nextNode.childNode[ulj] != nil {
										node.childNode[ulj].failNode = nextNode.childNode[ulj]
										break
									} else {
										nextNode = nextNode.failNode
									}
								}
							}
						}
					}
				}
			}

			nodeList.Remove(ele)
		}
	}
}
//字符串是否匹配
func AcAutoMatch(input string) bool {
	node := GAcAuto
	for _, runeCh := range input {
		count := 0
		runeStr := string(runeCh)
		for uli := 0; uli < len(runeStr); uli ++ {
			for ulj := 0; ulj < 2; ulj++ {
				index := runeStr[uli] / childNodeCount
				if ulj != 0 {
					index = runeStr[uli] % childNodeCount
				}

			Match:
				if node != nil && node.childNode[index] != nil {
					// 找到即退出，没有结束继续查找
					count++
					if node.childNode[index].endCount > 0 && count == 2 * len(runeStr) {
						return true
					}  else {
						node = node.childNode[index]
					}
				} else {
					if node == nil {
						// 当前字符一直查不到，则换下个字符，故重置node，取值根节点
						node = GAcAuto
					} else {
						// 当前节点没有目标字符，则去下个节点看下，即查看fail指针
						node = node.failNode
						goto Match
					}
				}
			}
		}
	}
	return false
}
