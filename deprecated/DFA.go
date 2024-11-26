package util

import "errors"

type DFA struct {
	beginning    DNode
	dnodes        map[string]DNode
}

// DNode = Determinitstic Node
type DNode struct {
	Name        string
	Transitions map[string] DNode
	Final       bool
}

func MakeDFA(transitions [][3]string, beginning string, finishStates []string) (*DFA) {
	// Create new DFA 
	DFA := new(DFA)

	// Add the Transitions to the DFA
	for _, newTransition := range transitions {
		DFA.addTransition(newTransition)
	}

	// Finish States Map
	finishMap := make(map[string]bool)
	for _, f := range finishStates{
		finishMap[f] = true
	}

	// Iterate over the Nodes and make them Final
	for name, isFinish := range finishMap{
		if isFinish{
			finishNode := DFA.dnodes[name]
			finishNode.Final = true	
		}
	}

	// Add the beginning node and return
	beginningNode, ok := DFA.dnodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *new(DNode)
		beginningNode.Name = beginning
		DFA.dnodes[beginning] = beginningNode
	}

	DFA.beginning = beginningNode
	return DFA
}

func (DFA *DFA) addTransition(newTransition [3]string) {
	startNode, containsStart := DFA.dnodes[newTransition[0]]
	endNode, containsEnd := DFA.dnodes[newTransition[2]]

	if !containsStart {
		startNode = *DFA.createNode(newTransition[0])
	}

	if !containsEnd {
		endNode = *DFA.createNode(newTransition[2])
	}

	startNode.Transitions[newTransition[1]] = endNode
}

func (DFA *DFA) createNode (name string) *DNode{
	newNode := new(DNode)
	newNode.Name = name

	// Adds DNode to Hashmap
	DFA.dnodes[name] = *newNode
	return newNode
}

func (head *DNode) accepts(input []string) bool {
	if len(input) == 0 {
		return head.Final
	}
	// Get first string of input
	nextLiteral := input[0]

	nextNode, err := head.getNext(nextLiteral)
	if err != nil {
		return false
	}
	// Slice the string without the first string
	return nextNode.accepts(input[1:])
}

func (head *DNode) getNext(a string) (DNode, error) {
	nextNode, ok := head.Transitions[a]
	if !ok{
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (DFA *DFA) getStart() DNode{
	return DFA.beginning
}

func (node *DNode) isFinal() bool {
	return node.Final
}

func (node *DNode) getName() string {
	return node.Name
}

func (node *DNode) getEdges() map[string] DNode{
	return node.Transitions
}