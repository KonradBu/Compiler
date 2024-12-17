package parser

import (
	"compiler/lexer"
	"fmt"
	"github.com/m1gwings/treedrawer/tree"
)

type parseTree struct {
	leaf parseLeaf
	branches []parseTree
}

type parseLeaf struct{
	name string
	value any
} 

func createParseTree(parseChan chan any)  {
	Trees := []parseTree{}
	for true{
		newItem := <- parseChan 
		switch newItem.(type){
		case lexer.Token:
			token := newItem.(lexer.Token)
			newLeaf := parseLeaf{name: token.Identifier, value: token.Value}
			newTree := parseTree{leaf: newLeaf, branches: []parseTree{}}
			Trees = append(Trees, newTree)
		case Rule:
			rule := newItem.(Rule)
			newTree := parseTree{}
			newBranches := []parseTree{}
			newTree.leaf = parseLeaf{name: rule.nonTerminal, value: 0}
			for i := range rule.production{
				len := len(Trees)
				if  len == 0{
					panic("Parse Tree Error, no new node possible")
				}
				
				newBranches = append(newBranches, Trees[len - i -1])
			}
			Trees = Trees[:len(Trees) - len(rule.production)]
			newTree.branches = newBranches
			Trees = append(Trees, newTree)
		case bool:
			parseChan <- Trees[0]
			close(parseChan)
		}
	}
	parseChan <- Trees[0]
	close(parseChan)
}

func PrintTree(tree parseTree){
	fmt.Print(tree.leaf.name)
	fmt.Print(" (")
	fmt.Print(tree.leaf.value)
	fmt.Print("): ")
	fmt.Print(tree.branches)
	for _, t := range tree.branches{
		PrintTree(t)
	}
}