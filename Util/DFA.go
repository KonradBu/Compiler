package main

import "errors"

type Automata struct {
	beginning    Node
	nodes        map[rune]Node
}

type Node struct {
	Name        rune
	Transitions map[rune] []Node
	Final       bool
}

func makeAutomata(transitions [][3]rune, beginning rune, finishStates []rune) (*Automata) {
	// Create new Automata 
	Automata := new(Automata)

	// Add the Transitions to the Automata
	for _, newTransition := range transitions {
		Automata.addTransition(newTransition)
	}


	// Finish States Map
	var finishMap map[rune]bool
	for _, f := range finishStates{
		finishMap[f] = true
	}

	// Iterate over the Nodes and make them Final
	for name, isFinish := range finishMap{
		if isFinish{
			finishNode := Automata.nodes[name]
			finishNode.Final = true	
		}
	}

	// Add the beginning node and return
	beginningNode, ok := Automata.nodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *new(Node)
		beginningNode.Name = beginning
		Automata.nodes[beginning] = beginningNode
	}

	Automata.beginning = beginningNode
	return Automata
}

func (Automata Automata) addTransition(newTransition [3]rune) {
	startNode, containsStart := Automata.nodes[newTransition[0]]
	endNode, containsEnd := Automata.nodes[newTransition[2]]

	if !containsStart {
		startNode = *Automata.createNode(newTransition[0])
	}

	if !containsEnd {
		endNode = *Automata.createNode(newTransition[2])
	}

	end, ok := startNode.Transitions[newTransition[1]]

	if !ok{
		end = []Node{endNode}
	}else{
		end = append(end, endNode)
	}
}

func (Automata Automata) createNode (a rune) *Node{
	newNode := new(Node)
	newNode.Name = a

	// Adds Node to Hashmap
	Automata.nodes[a] = *newNode
	return newNode
}

func (head Node) accepts(input string) bool {
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
	for _, newNode := range nextNode{
		if newNode.accepts(input[1:]){
			return true
		}
	}
	return false
}

func (head Node) getNext(a rune) ([]Node, error) {
	nextNode, ok := head.Transitions[a]
	if !ok{
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (Automata Automata) getStart() Node{
	return Automata.beginning
}

func (head Node) isFinal() bool {
	return head.Final
}

func (head Node) getName() rune {
	return head.Name
}