package parser

import (
	"fmt"
	"strings"
)

type Stack struct {
	Top *StackObject
}

type StackObject struct {
	Val  any
	Next *StackObject
}

func makeStack(item any) *Stack {
	newStack := new(Stack)
	newStackObject := new(StackObject)
	newStack.Top = newStackObject
	newStackObject.Val = &item
	newStackObject.Next = nil
	return newStack
}

func (stack *Stack) pop() any {
	if stack.Top == nil {
		return nil
	}
	tmp := stack.Top.Val
	stack.Top = stack.Top.Next
	return tmp
}

func (stack *Stack) add(item any) {
	newObject := new(StackObject)
	newObject.Val = &item
	newObject.Next = stack.Top
	stack.Top = newObject
}

func (stack *Stack) peek() any {
	if stack.Top == nil {
		return nil
	}
	return stack.Top.Val
}

func (stack *Stack) Print() string {
	var builder strings.Builder
	current := stack.Top
	for current != nil {
		fmt.Fprintf(&builder, "%v\n", current.Val)
		current = current.Next
	}
	return builder.String()
}
