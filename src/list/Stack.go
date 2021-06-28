package list

/**
 * 栈逻辑，支持先进先出的逻辑
 */
import (
	"errors"
)

type Stack []interface {
}

func (stack Stack) Len() int {
	// 计算数组的长度
	return len(stack)
}

func (stack Stack) Empty() bool {
	// 计算数组是否是空的
	return len(stack) == 0
}

func (stack *Stack) push(value interface{}) {

	*stack = append(*stack, value) //插入到最后一个数组
}

func (stack Stack) top() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("out of index,len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) pop() (interface{}, error) {
	theStack := *stack
	if stack.Empty() {
		return nil, errors.New("stack is empty,len is 0")
	}
	item := theStack[len(theStack)-1]      //返回最后一个数组
	*stack = theStack[0 : len(theStack)-1] //重新生成切片
	return item, nil
}
