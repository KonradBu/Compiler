package automata

import (
	"errors"
	"sort"
	"sync"
)

type automata struct {
	beginning node
	nodes     map[string]node
}

type node struct {
	Name        string
	Transitions map[string][]node
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
	automata := new(automata)

	// Add the Transitions to the automata
	for _, newTransition := range transitions {
		automata.AddTransition(newTransition)
	}

	// Finish States Map
	finishMap := make(map[string]bool)
	for _, f := range finishStates {
		finishMap[f] = true
	}

	// Iterate over the Nodes and make them Final
	for name, isFinish := range finishMap {
		if isFinish {
			finishNode := automata.nodes[name]
			finishNode.Final = true
		}
	}

	// Add the beginning node and return
	beginningNode, ok := automata.nodes[beginning]

	// If the beginning node hasnt been generated yet
	if !ok {
		beginningNode = *automata.CreateNode(beginning)
	}

	automata.beginning = beginningNode
	return automata
}

func (automata *automata) AddTransition(newTransition [3]string) *node {
	// newTransition [0] = Beginning Node; [1] = input; [2] = end node

	startNode, containsStart := automata.nodes[newTransition[0]]
	endNode, containsEnd := automata.nodes[newTransition[2]]

	if !containsStart {
		startNode = *automata.CreateNode(newTransition[0])
	}

	if !containsEnd {
		endNode = *automata.CreateNode(newTransition[2])
	}

	// Add node to map
	end, ok := startNode.Transitions[newTransition[1]]

	if !ok {
		end = []node{endNode}
	} else {
		end = append(end, endNode)
	}
	return &endNode
}

// Only call if the automata is a DFA!!
func (head *node) DFAaccepts(input []string) bool {
	if len(input) == 0 {
		return head.Final
	}
	// Get first string of input
	nextLiteral := input[0]

	nextNode := head.GetNext(nextLiteral)
	if len(nextNode) == 0{
		return false
	} 
	// Slice the string without the first string
	return nextNode[0].DFAaccepts(input[1:])
}


func (head *node) Accepts(input []string) bool {
	// Channel to check if the finish has been found already
	found := make(chan bool)

	// Has this combination of Node and Input Strings been checked already?
	// Map From Name of the State -> Another Map from a string array to the bool value
	checked := make(map[string]map[[]string]bool)

	// Initialize wait group
	var wg sync.WaitGroup
	wg.Add(1)

	// Create channel that waits for the end of the waitgroup
	done := make(chan bool)
	go func() {
		wg.Wait()
		close(done)
	}()

	// Launch go routines from the head
	// Signature: input string, channel for early exit, check for checking if
	// we have checked the node + input before, waitgroup for concurrency
	//(Checking if every go routine has finished)
	go head.acceptsRoutine(input, found, checked, &wg)

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
		// Channel is found -> Finish was found
		if !ok {
			return
		}
	default:
		// Do Nothing
	}

	// Check if the Input string is over
	if len(input) == 0 {
		if head.Final {
			close(found)
		}
		wg.Done()
	}

	// Check if we have been here before:
	// Suprisingly hashing arrays compares content, not identity
	if checked[head.Name][input] {
		wg.Done()
		return
	} else {
		checked[head.Name][input] = true
	}

	// Get first string of input
	nextRune := input[0]

	// Check if there is a transition
	nextNodes, err := head.GetNext(nextRune)
	eTransition := head.EpsilonTransition

	if err != nil && len(eTransition) == 0 {
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
	for _, newNode := range head.EpsilonTransition {
		// Input the full string
		go newNode.acceptsRoutine(input, found, checked, wg)
		wg.Add(1)
	}

	wg.Done()
}

// Pls dont have names of states that combine to other names of states (e.g. no states like: a,b,ab)
func (NFA *automata) ToDFA() *automata {
	DFA := new(automata)
	DFA.beginning = DFA.recursiveMerge(NFA.beginning)
	return DFA
}

// Takes a node and recursivly merges all the states
func (DFA *automata) recursiveMerge(head node) node {
	// Epsilon Closure of itself
	toBeMergedNodes := head.EpsilonClosure()

	// Creates new node of itself + closure
	returnNode, mergedNodes, err := DFA.makeCompositNode(toBeMergedNodes)

	// Have we created this node already?
	if err != nil {
		return *returnNode
	}

	// Goes through all the transitions of the set for every input
	// All the nodes that have just been merged into 1
	for _, mergedNode := range mergedNodes {
		// All the transitions of said node
		for input, endNode := range mergedNode.Transitions {
			// Add the transitions to the new node
			_, exists := returnNode.Transitions[input]

			if !exists {
				returnNode.Transitions[input] = []node{}
			}
			// Add the transitions of the new node to the old node
			returnNode.Transitions[input] = append(mergedNode.Transitions[input], endNode...)
		}
	}

	// Makes Final
	returnNode.Final = false
	for _, node := range mergedNodes {
		if node.Final {
			returnNode.Final = true
		}
	}

	// Recursivly calls itself on the newly created nodes
	for input, newNode := range returnNode.Transitions {
		// The newnode has to be of length 1
		returnNode.Transitions[input] = []node{DFA.recursiveMerge(newNode[0])}
	}

	return *returnNode
}

// Creates a composit node out of a bunch of nodes and their epsilon closure
func (NFA *automata) makeCompositNode(startNodes []node) (*node, []node, error) {
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
	newName := compositNodeName(newNameParts)

	// Check if we have made this node already
	alreadyExistingNode, existsAlready := NFA.nodes[newName]
	if existsAlready {
		return &alreadyExistingNode, nodes, errors.New("Node already exists")
	}

	return NFA.CreateNode(newName), nodes, nil
}

// Composits the names of a bunch of nodes so that they are always the same
func compositNodeName(names []string) string {
	// So that the name of the States in not dependand on node order
	sort.Strings(names)
	newName := ""
	for _, s := range names {
		newName += s
	}
	return newName
}

// Creates a node and adds it to the automata
func (automata *automata) CreateNode(a string) *node {
	newNode := new(node)
	newNode.Name = a
	newNode.Transitions = make(map[string][]node)

	// Adds node to Hashmap
	automata.nodes[a] = *newNode
	return newNode
}

// Gets all the nodes reachable from a specific node using only one input a and epsilon transitions
func (head *node) GetNext(a string) []node {
	nextNodes, ok := head.Transitions[a]

	// Append epsilon transitions
	nextNodes = append(nextNodes, head.EpsilonClosure()...)

	// No Transitions for this input
	if !ok {
		return []node{}
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

func (automata *automata) GetStart() node {
	return automata.beginning
}

func (node *node) IsFinal() bool {
	return node.Final
}

func (node *node) GetName() string {
	return node.Name
}

func (node *node) GetEdges() map[string][]node {
	return node.Transitions
}