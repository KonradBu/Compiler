package parser

type Stack struct{
	Val *any
	Next *Stack
}

func makeStack(item any) *Stack{
	newStack := new(Stack)
	newStack.Val = &item
	newStack.Next = nil
	return newStack
}

func (stack *Stack) pop() any{
	if stack.Val == nil{
		return nil
	}
	tmp := stack.Val
	stack = stack.Next
	return tmp
}

func (stack *Stack) add(item any){
	newTop := new(Stack)
	newTop.Val = &item
	newTop.Next = stack
	stack = newTop
}
