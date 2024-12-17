package parser

import (
	"compiler/lexer"
	"slices"

	"github.com/pterm/pterm"
)

type parseTree struct {
	leaf     parseLeaf
	branches []parseTree
}

type parseLeaf struct {
	name  string
	value any
}

func createParseTree(parseChan chan any) {
	Trees := []parseTree{}
	for true {
		newItem := <-parseChan
		switch newItem.(type) {
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
			for i := range rule.production {
				len := len(Trees)
				if len == 0 {
					panic("Parse Tree Error, no new node possible")
				}

				newBranches = append(newBranches, Trees[len-i-1])
			}
			Trees = Trees[:len(Trees)-len(rule.production)]
			slices.Reverse(newBranches)
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

func PrintTree(tree parseTree) {
	ptree := makePTree(tree)
	renderTree := pterm.DefaultTree.WithRoot(ptree)
	renderTree.Render()
}

func makePTree(tree parseTree) pterm.TreeNode {
	root := pterm.TreeNode{Text: tree.leaf.name, Children: []pterm.TreeNode{}}
	for _, t := range tree.branches {
		root.Children = append(root.Children, makePTree(t))
	}
	return root
}
