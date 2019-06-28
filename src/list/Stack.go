package list

import (
	"errors"
)

type Stack []interface {
}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Empty() bool {
	return len(stack) == 0
}

func (stack *Stack) push(value interface{}){
	*stack=append(*stack,value)
}

func (stack Stack) top() (interface{},error){
	if stack.Empty(){
		return nil,errors.New("out of index,len is 0")
	}
	return stack[len(stack)-1],nil
}

func (stack Stack) pop() (interface{},error){

}