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

func (automata *automata) AddTransition(newTransition [3]string) *node{
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

func (head *node) accepts(input []string) bool {
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

func (automata *automata) toDFA(){
	// Iterate over all the Nodes if the automata
	for _, currentNode := range automata.nodes{
		// Iterate over all the transition of every node
		for input, endNodes := range currentNode.Transitions{
			// Add epsilon closure of all the nodes
			for _, epsilonCheckNode := range endNodes{
				endNodes = append(endNodes, epsilonCheckNode.EpsilonClosure()...)
			}  
			// Check if the nodes has any nondeterminism, and if yes create a new composit node
			if len(endNodes) > 1{
				automata.addCompositNode(&currentNode, input, endNodes)
			}
		}
	}
} 


func (automata *automata) addCompositNode(oldnode *node, input string, endNodes[]node) *node{
	newName := "" 
	// create 
	sort.Slice(endNodes, func(i, j int) bool {
		
	})
	for _, endNode := range endNodes{
		newName += endNode.Name
	}
	node, ok := automata.nodes[] 
	compositNode := automata.AddTransition([3]string{oldnode.Name, input, newName})
	for 
	
}

func (automata *automata) CreateNode(a string) *node {
	newNode := new(node)
	newNode.Name = a

	// Adds node to Hashmap
	automata.nodes[a] = *newNode
	return newNode
}

func (head *node) GetNext(a string) ([]node, error) {
	nextNode, ok := head.Transitions[a]
	if !ok {
		return nextNode, errors.New("No transition found")
	}
	return nextNode, nil
}

func (inputNode *node) EpsilonClosure() []node {
	var returnSlice []node
	for input, endNode := range inputNode.Transitions {
		if input == " " {
			returnSlice = append(returnSlice, endNode...)
		}
	}
	return returnSlice
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
