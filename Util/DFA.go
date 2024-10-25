package main

type DFANode struct{
	// PLS do capital letter for Nodes and non capital for input alphabet
	Name rune
	Transitions map[rune] DFANode
	isFinal bool
}

/* func makeDFA (transitions [][]rune, beginning rune, end []rune) *DFANode{
	inputHead := new(DFANode)
	inputHead.Name = beginning

}

func (head DFANode) canRead ()
*/