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

func (stack *Stack) pop() (*Stack, any){
	if stack.Val == nil{
		return &Stack{}, nil
	}
	tmp := stack.Val
	return stack.Next, tmp
}

func (stack *Stack) add(item any) *Stack{
	newTop := new(Stack)
	newTop.Val = &item
	newTop.Next = stack
	return newTop
}
