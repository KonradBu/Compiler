package util

import (
	"errors"
	"sync"
)

type NFA struct {
	beginning NNode
	nnodes    map[string]NNode
}

type NNode struct {
	Name              string
	Transitions       map[string][]NNode
	EpsilonTransition []NNode
	Final             bool
}

// Space: Epsilon transitions!
func makeNFA(transitions [][3]string, beginning string, finishStates []string) *NFA {
	// Create new NFA
	NFA := new(NFA)

	// Add the Transitions to the NFA
	for _, newTransition := range transitions {
		NFA.addTransition(newTransition)
	}

	// Finish States Map
	finishMap := make(map[string]bool)
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

func (NFA *NFA) addTransition(newTransition [3]string) {
	startNode, containsStart := NFA.nnodes[newTransition[0]]
	endNode, containsEnd := NFA.nnodes[newTransition[2]]

	if !containsStart {
		startNode = *NFA.createNode(newTransition[0])
	}

	if !containsEnd {
		endNode = *NFA.createNode(newTransition[2])
	}

	// Check for epsilon transitions
	if newTransition[1] == " " {
		startNode.EpsilonTransition = append(startNode.EpsilonTransition, endNode)
		return
	}

	// Is not an epsilon
	// Add node to map
	end, ok := startNode.Transitions[newTransition[1]]

	if !ok {
		end = []NNode{endNode}
	} else {
		end = append(end, endNode)
	}
}

func (NFA *NFA) createNode(a string) *NNode {
	newNode := new(NNode)
	newNode.Name = a

	// Adds NNode to Hashmap
	NFA.nnodes[a] = *newNode
	return newNode
}

func (head *NNode) accepts(input []string) bool {
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

func (head *NNode) acceptsRoutine(input []string, found chan bool, checked map[string]map[string]bool, wg *sync.WaitGroup) {

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
	nextNodes, err := head.getNext(nextRune)
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

func (NFA *NFA) toDFA() *DFA {
	var transitions [][3]string
	var finishState []string
	for _, n := range NFA.nnodes {
		if n.Final == true {
			finishState = append(finishState, n.Name)
		}
	}

	// Create new transitions

	//

	return MakeDFA(transitions, NFA.beginning.Name, finishState)
}

func (head *NNode) getNext(a string) ([]NNode, error) {
	nextNode, ok := head.Transitions[a]
	if !ok {
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (node *NNode) epsilonClosure() []NNode {
	return node.EpsilonTransition
}

func (NFA *NFA) getStart() NNode {
	return NFA.beginning
}

func (node *NNode) isFinal() bool {
	return node.Final
}

func (node *NNode) getName() string {
	return node.Name
}

func (node *NNode) getEdges() map[string][]NNode {
	return node.Transitions
}
