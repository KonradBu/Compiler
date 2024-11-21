package main

import (
	"errors"
	"sort"
	"sync"
)

type automata struct {
	beginning node
	nodes     map[string]*node
}

type node struct {
	Name        string
	Transitions map[string][]*node
	Final       bool
}

// Signature:
// Transitions as an Slice of 3 long arrays: Beginning Node-> Input String-> End Node
// Space: Epsilon transitions!
// Beginnign: Name of beginning Node (does not need to be defined by the transitions)
// finishStates: Name of nodes that are finishes (Need to be defined by the transitions)
// Returns: Pointer to an automata
func MakeAutomata(transitions [][3]string, beginning string, finishStates []string) *automata {
	// Create new automata
	newAutomata := new(automata)
	newAutomata.nodes = make(map[string]*node)

	// Add the Transitions to the automata
	for _, newTransition := range transitions {
		newAutomata.AddTransition(newTransition)
	}

	// Set final on final nodes
	for _, name := range finishStates{
		newAutomata.nodes[name].Final = true
	}
	// Add the beginning node and return
	beginningNode, ok := newAutomata.nodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = newAutomata.CreateNode(beginning)
	}

	newAutomata.beginning = *beginningNode
	return newAutomata
}

func (automata *automata) AddTransition(newTransition [3]string) *node {
	// newTransition [0] = Beginning Node; [1] = input; [2] = end node
	startNode, containsStart := automata.nodes[newTransition[0]]
	endNode, containsEnd := automata.nodes[newTransition[2]]

	if !containsStart {
		startNode = automata.CreateNode(newTransition[0])
	}

	if !containsEnd {
		endNode = automata.CreateNode(newTransition[2])
	}

	// Add node to map
	_, ok := startNode.Transitions[newTransition[1]]

	if !ok {
		startNode.Transitions[newTransition[1]] = []*node{endNode}
	} else {
		startNode.Transitions[newTransition[1]] = append(startNode.Transitions[newTransition[1]], endNode)
	}
	return endNode
}

// Only call if the automata is a DFA!!
func (head *node) DFAaccepts(input []string) bool {
	if len(input) == 0 {
		return head.Final
	}
	// Get first string of input
	nextLiteral := input[0]

	nextNode := head.GetNext(nextLiteral)
	if len(nextNode) == 0 {
		return false
	}
	// Slice the string without the first string
	return nextNode[0].DFAaccepts(input[1:])
}

// Pls dont have any inputs that can create each other if concatonated; e.g. aba, a, b
func (automata *automata) Accepts(input []string) bool {
	// Channel to check if the finish has been found already
	// Channels  ~ Message Passing (Saved in parent process memory)
	found := make(chan bool)

	// Has this combination of Node and Input Strings been checked already?
	// Map From Name of the State -> Another Map from a concatonated together input array to the bool value
	checked := make(map[string]map[string]bool)

	// Initialize wait group
	// ~ Thread safe counter
	var wg sync.WaitGroup
	wg.Add(1)

	// Create channel that waits for the end of the waitgroup
	// Useful for the select statement
	done := make(chan bool)
	go func() {
		wg.Wait()
		close(done)
	}()

	// Launch go routines from the head
	// Signature: input string, channel for early exit, check for checking if
	// we have checked the node + input before, waitgroup for concurrency
	//(Checking if every go routine has finished)
	go automata.beginning.acceptsRoutine(input, found, checked, &wg)

	// Wait until either: Every go routine finishes, or: A finish was found
	select {
	case <-done:
		return false
	case <-found:
		return true
	}
}

func (head *node) acceptsRoutine(input []string, found chan bool, checked map[string]map[string]bool, wg *sync.WaitGroup) {

	// Checks if channel exists or not, without blocking
	// If a select has a default, then it doesnt wait until finish, but instead
	// Continues on
	select {
	case _, ok := <-found:
		// Channel is closed -> Finish was found
		if !ok {
			return
		}
	default:
		// Do Nothing
	}

	// Check if the Input string is over
	if len(input) == 0 {
		// Can we reach a Finish using epsilon transitions?
		for _, reachableNodes := range head.EpsilonClosure() {
			if reachableNodes.Final {
				close(found)
			}
		}
		wg.Done()
		return
	}

	// Check if we have been here before:
	// Does the map exist?
	_, ok := checked[head.Name]
	if !ok {
		tmp := make(map[string]bool)
		checked[head.Name] = tmp
	}

	// Concatonate the input strings
	if checked[head.Name][concatonateStringArraySorted(input)] {
		wg.Done()
		return
	} else {
		checked[head.Name][concatonateStringArraySorted(input)] = true
	}

	// Get first string of input
	nextRune := input[0]

	// Check if there is a transition
	nextNodes := head.GetNext(nextRune)
	eTransitions := head.EpsilonClosure()
	// Remove Itself
	eTransitions = eTransitions[1:]

	if len(nextNodes) == 0 && len(eTransitions) == 0 {
		wg.Done()
		return
	}

	// Startup new go routines
	for _, newNode := range nextNodes {
		// Slice the string without the first string
		go newNode.acceptsRoutine(input[1:], found, checked, wg)
		wg.Add(1)
	}

	// Startup new go routines for the epsilon closure
	for _, newNode := range eTransitions {
		// Input the full string
		go newNode.acceptsRoutine(input, found, checked, wg)
		wg.Add(1)
	}

	wg.Done()
}

// Pls dont have names of states that combine to other names of states (e.g. no states like: a,bb,ab,ba)
func (NFA *automata) ToDFA() *automata {
	DFA := new(automata)
	DFA.beginning = *DFA.recursiveMerge(&NFA.beginning)
	return DFA
}

// Takes a node and recursivly merges all the states
func (DFA *automata) recursiveMerge(head *node) *node {
	// Epsilon Closure of itself
	toBeMergedNodes := head.EpsilonClosure()

	// Creates new node of itself + closure
	returnNode, err := DFA.makeCompositNode(toBeMergedNodes)

	// Have we created this node already?
	// Cant do this just over the name of the start node, as multiple start nodes may create the same merge
	if err != nil {
		return returnNode
	}

	// Goes through all the transitions of the set for every input
	// All the nodes that have just been merged into 1
	for _, mergedNode := range toBeMergedNodes {
		// All the transitions of said node
		for input, endNode := range mergedNode.Transitions {
			// Add the transitions to the new node
			_, exists := returnNode.Transitions[input]

			if !exists {
				returnNode.Transitions[input] = []*node{}
			}
			// Add the transitions of the new node to the old node
			returnNode.Transitions[input] = append(mergedNode.Transitions[input], endNode...)
		}
	}

	// Makes Final
	returnNode.Final = false
	for _, node := range toBeMergedNodes {
		if node.Final {
			returnNode.Final = true
		}
	}

	// Recursivly calls itself on the newly created nodes
	for input, newNode := range returnNode.Transitions {
		// The newnode has to be of length 1
		returnNode.Transitions[input] = []*node{DFA.recursiveMerge(newNode[0])}
	}
	return returnNode
}

// Creates a composit node out of a bunch of nodes and their epsilon closure
func (NFA *automata) makeCompositNode(startNodes []node) (*node, error) {
	var nodes []node
	// Add the epsilon closure
	for _, node := range startNodes {
		// The epsilon closure contains itself
		nodes = append(nodes, node.EpsilonClosure()...)
	}

	// Creates the composit node
	var newNameParts []string
	for _, node := range nodes {
		newNameParts = append(newNameParts, node.Name)
	}
	newName := concatonateStringArraySorted(newNameParts)

	// Check if we have made this node already
	alreadyExistingNode, existsAlready := NFA.nodes[newName]
	if existsAlready {
		return alreadyExistingNode, errors.New("Node already exists")
	}

	return NFA.CreateNode(newName), nil
}

// Creates a node and adds it to the automata
func (automata *automata) CreateNode(a string) *node {
	newNode := new(node)
	newNode.Name = a
	newNode.Transitions = make(map[string][]*node)

	// Adds node to Hashmap
	automata.nodes[a] = newNode
	return newNode
}

// Gets all the nodes reachable from a specific node using only one input a
func (head *node) GetNext(a string) []*node {
	nextNodes, ok := head.Transitions[a]

	// No Transitions for this input
	if !ok {
		return []*node{}
	}
	return nextNodes
}

func (inputNode *node) EpsilonClosure() []node {
	// Create map for easy lookup
	eTransitions := make(map[string]node)

	// Add all other nodes recursivly
	inputNode.epsilonRecursive(eTransitions)

	// Make slice to return
	var closure []node
	for _, node := range eTransitions {
		closure = append(closure, node)
	}
	return closure
}

func (inputNode *node) epsilonRecursive(eTransitions map[string]node) {
	// Add itself
	eTransitions[inputNode.Name] = *inputNode

	// Add all the current epsilon transitions
	for _, node := range inputNode.Transitions[" "] {
		_, checked := eTransitions[node.Name]
		if !checked {
			node.epsilonRecursive(eTransitions)
		}
	}
}

// Composits the names of a bunch of strings so that they are always the same
func concatonateStringArraySorted(names []string) string {
	// So that the name of the States in not dependand on node order
	sort.Strings(names)
	newName := ""
	for _, s := range names {
		newName += s
	}
	return newName
}

func (automata *automata) GetStart() node {
	return automata.beginning
}

func (node *node) IsFinal() bool {
	return node.Final
}

func (node *node) GetName() string {
	return node.Name
}

func (node *node) GetEdges() map[string][]*node {
	return node.Transitions
}
