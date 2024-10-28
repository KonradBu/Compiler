package main

import "errors"

type DFA struct {
	beginning    Node
	nodes        map[rune]Node
	finishStates map[rune]bool
}

type Node struct {
	Name        rune
	Transitions map[rune]Node
	Final       bool
}

func makeDFA(transitions [][3]rune, beginning rune, finishStates []rune) (*DFA, error) {
	// Check for not valid input
	if len(transitions) == 0 {
		return nil, errors.New("No Transitions in Input")
	}

	// Create new DFA and add the finish states 
	DFA := new(DFA)
	for _, state := range finishStates {
		DFA.finishStates[state] = true
	}

	// Add the Transitions to the DFA
	for _, newTransition := range transitions {
		DFA.addTransition(newTransition)
	}

	// Add the beginning node and return
	beginningNode, ok := DFA.nodes[beginning]
	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *new(Node)
		beginningNode.Name = beginning
		DFA.nodes[beginning] = beginningNode
		DFA.beginning = beginningNode
	}

	DFA.beginning = beginningNode
	return DFA, nil
}

func (DFA DFA) addTransition(newTransition [3]rune) {
	startNode, containsStart := DFA.nodes[newTransition[0]]
	endNode, containsEnd := DFA.nodes[newTransition[2]]

	if !containsStart {
		newNode := new(Node)
		newNode.Name = newTransition[0]
		if DFA.finishStates[newNode.Name] {
			newNode.Final = true
		}
		startNode = *newNode
	}

	if !containsEnd {
		newNode := new(Node)
		newNode.Name = newTransition[2]
		if DFA.finishStates[newNode.Name] {
			newNode.Final = true
		}
		endNode = *newNode
	}

	startNode.Transitions[newTransition[1]] = endNode
}

func (head Node) accepts(input string) bool {
	if len(input) == 0 {
		return head.Final
	}
	nextLiteral := []rune(input)[0]
	nextNode, err := head.getNext(nextLiteral)
	if err != nil {
		return false
	}
	// Slice the string without the first rune
	return nextNode.accepts(input[1:])
}

func (head Node) getNext(a rune) (Node, error) {
	nextNode, ok := head.Transitions[a]
	if !ok{
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (head Node) isFinal() bool {
	return head.Final
}

func (head Node) getName() rune {
	return head.Name
}