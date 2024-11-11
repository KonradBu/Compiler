package util

import "errors"

type Automata struct {
	beginning    node
	
	// Name -> Pointer
	nodes        map[rune]node
	
	finishStates map[rune]bool
}

type node struct {
	Name        rune
	
	// Input -> NodeArray
	Transitions map[rune] []node

	Final       bool
}


func MakeAutomata(transitions [][3]rune, beginning rune, finishStates []rune) (*Automata, error) {
	// Check for not valid input
	if len(transitions) == 0 {
		return nil, errors.New("No Transitions in Input")
	}

	// Create new Automata and add the finish states 
	Automata := new(Automata)
	for _, state := range finishStates {
		Automata.finishStates[state] = true
	}

	// Add the Transitions to the Automata
	for _, newTransition := range transitions {
		Automata.AddTransition(newTransition)
	}

	// Add the beginning node and return
	beginningNode, ok := Automata.nodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *new(node)
		beginningNode.Name = beginning
		Automata.nodes[beginning] = beginningNode
	}

	Automata.beginning = beginningNode
	return Automata, nil
}

func (Automata Automata) AddTransition(newTransition [3]rune) {
	startNode, containsStart := Automata.nodes[newTransition[0]]
	endNode, containsEnd := Automata.nodes[newTransition[2]]

	if !containsStart {
		newNode := new(node)
		newNode.Name = newTransition[0]
		if Automata.finishStates[newNode.Name] {
			newNode.Final = true
		}
		// Adds node to Hashmap
		Automata.nodes[newTransition[0]] = *newNode
		startNode = *newNode
	}

	if !containsEnd {
		newNode := new(node)
		newNode.Name = newTransition[2]
		if Automata.finishStates[newNode.Name] {
			newNode.Final = true
		}
		// Adds node to Hashmap
		Automata.nodes[newTransition[2]] = *newNode
		endNode = *newNode
	}

	startNode.Transitions[newTransition[1]] = endNode
}

func (head node) Accepts(input string) bool {
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
	return nextNode.Accepts(input[1:])
}

func (head node) GetNext(a rune) (node, error) {
	nextNode, ok := head.Transitions[a]
	if !ok{
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (Automata Automata) GetStart() node{
	return Automata.beginning
}

func (head node) IsFinal() bool {
	return head.Final
}

func (head node) GetName() rune {
	return head.Name
}

func (head node) GetEdges() map[rune] []node{
	return head.Transitions
}