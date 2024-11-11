package util

import "errors"

type NFA struct {
	beginning NNode
	nnodes     map[rune]NNode
}

type NNode struct {
	Name        rune
	Transitions map[rune][]NNode
	EpsilonTransition []NNode 
	Final       bool
}

func makeNFA(transitions [][3]rune, beginning rune, finishStates []rune) *NFA {
	// Create new NFA
	NFA := new(NFA)

	// Add the Transitions to the NFA
	for _, newTransition := range transitions {
		NFA.addTransition(newTransition)
	}

	// Finish States Map
	finishMap := make(map[rune]bool)
	for _, f := range finishStates {
		finishMap[f] = true
	}

	// Iterate over the Nodes and make them Final
	for name, isFinish := range finishMap {
		if isFinish {
			finishNode := NFA.nnodes[name]
			finishNode.Final = true
		}
	}

	// Add the beginning node and return
	beginningNode, ok := NFA.nnodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *new(NNode)
		beginningNode.Name = beginning
		NFA.nnodes[beginning] = beginningNode
	}

	NFA.beginning = beginningNode
	return NFA
}

func (NFA NFA) addTransition(newTransition [3]rune) {
	startNode, containsStart := NFA.nnodes[newTransition[0]]
	endNode, containsEnd := NFA.nnodes[newTransition[2]]

	if !containsStart {
		startNode = *NFA.createNode(newTransition[0])
	}

	if !containsEnd {
		endNode = *NFA.createNode(newTransition[2])
	}

	end, ok := startNode.Transitions[newTransition[1]]

	if !ok {
		end = []NNode{endNode}
	} else {
		end = append(end, endNode)
	}
}

func (NFA NFA) createNode(a rune) *NNode {
	newNode := new(NNode)
	newNode.Name = a

	// Adds NNode to Hashmap
	NFA.nnodes[a] = *newNode
	return newNode
}

func (head NNode) accepts(input string) bool {
	if len(input) == 0 {
		return head.Final
	}
	// Get first rune of input
	nextLiteral := []rune(input)[0]

	nextNode, err := head.getNext(nextLiteral)
	if err != nil {
		return false
	}
	// Slice the string without the first rune
	for _, newNode := range nextNode {
		if newNode.accepts(input[1:]) {
			return true
		}
	}
	return false
}

func (node NNode) epsilonClosure() []NNode{

}

func (head NNode) getNext(a rune) ([]NNode, error) {
	nextNode, ok := head.Transitions[a]
	if !ok {
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (NFA NFA) getStart() NNode {
	return NFA.beginning
}

func (node NNode) isFinal() bool {
	return node.Final
}

func (node NNode) getName() rune {
	return node.Name
}

func (node NNode) getEdges() map[rune] []NNode{
	return node.Transitions
}