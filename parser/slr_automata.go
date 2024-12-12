package parser

import (
	"fmt"
)

type SLR_automata struct {
	states []State
}
type State struct {
	id          int
	rules       []ItemRule
	transitions map[string]int
}

type ItemRule struct {
	rule Rule
	dot  int
}

type GrammarClosure struct {
	closure map[string][]ItemRule
}

var stateIndex int

func (grammar *Grammar) CreateSLRAutomata() *SLR_automata {
	stateIndex = 0
	automata := new(SLR_automata)
	var startRule Rule
	for _, r := range grammar.rules {
		if r.nonTerminal == "S" {
			startRule = r
			break
		}
	}
	grammarClosure := grammar.makeGrammarClosure()
	startItemRule := new(ItemRule)
	startItemRule.dot = 0
	startItemRule.rule = startRule
	startState := makeState([]ItemRule{*startItemRule}, *grammarClosure)
	automata.states = append(automata.states, *startState)
	startState.id = 0
	stateIndex++
	startState.GoTo(automata, *grammarClosure)
	return automata
}

func (grammar *Grammar) makeGrammarClosure() *GrammarClosure {
	newClosure := new(GrammarClosure)
	newClosure.closure = make(map[string][]ItemRule)
	for symbol, closure := range grammar.closure {
		newClosure.closure[symbol] = []ItemRule{}
		for _, rule := range closure {
			newItemRule := new(ItemRule)
			newItemRule.rule = rule
			newItemRule.dot = 0
			newClosure.closure[symbol] = append(newClosure.closure[symbol], *newItemRule)
		}
	}
	return newClosure
}

func makeState(itemRules []ItemRule, closure GrammarClosure) *State {
	newState := new(State)
	newState.transitions = make(map[string]int)
	newState.rules = itemRules
	newState.addClosure(&closure)
	return newState
}

func (state *State) addClosure(closure *GrammarClosure) {
	done := make(map[string]bool)
	state.addClosureRecursive(*closure, done)
}

func (state *State) addClosureRecursive(closure GrammarClosure, done map[string]bool) {
	changed := false
	for _, itemrule := range state.rules {
		var nt string
		if itemrule.dot < len(itemrule.rule.production) {
			nt = itemrule.rule.production[itemrule.dot]
		} else {
			continue
		}
		if !done[nt] {
			// Add the entrie closure to the state
			for _, newrule := range closure.closure[nt] {
				// Check if the rule exists already
				rulesAreTheSame := false
				for _, existingrule := range state.rules {
					if areTheRulesTheSame(newrule, existingrule) {
						rulesAreTheSame = true
						break
					}
				}
				if rulesAreTheSame {
					continue
				}
				// Is a new rule
				state.rules = append(state.rules, newrule)
				changed = true
			}
			done[nt] = true
		}
	}
	if changed {
		state.addClosureRecursive(closure, done)
	}
}

func (oldState *State) GoTo(automata *SLR_automata, closure GrammarClosure) {
	rulesPerSymbol := make(map[string][]ItemRule)

	for _, r := range oldState.rules {
		if r.dot < len(r.rule.production) {
			if rulesPerSymbol[r.rule.production[r.dot]] == nil {
				rulesPerSymbol[r.rule.production[r.dot]] = []ItemRule{}
			}
			rulesPerSymbol[r.rule.production[r.dot]] = append(rulesPerSymbol[r.rule.production[r.dot]], r)
		}
	}

	for symbol, rules := range rulesPerSymbol {
		newRules := []ItemRule{}
		for _, rule := range rules {
			newItemRule := new(ItemRule)
			newItemRule.rule = rule.rule
			newItemRule.dot = rule.dot + 1
			newRules = append(newRules, *newItemRule)
		}
		newState := makeState(newRules, closure)
		newState.addClosure(&closure)

		existingState, doesNotExist := automata.stateDoesNotExist(newState)
		if doesNotExist {
			newState.id = stateIndex
			stateIndex++
			automata.states = append(automata.states, *newState)
			oldState.transitions[symbol] = newState.id
			newState.GoTo(automata, closure)
		} else {
			oldState.transitions[symbol] = existingState.id
			fmt.Println("Deleted:")
			fmt.Println(newState)
		}
	}
}

func (automata *SLR_automata) stateDoesNotExist(newState *State) (*State, bool) {
	for _, existingState := range automata.states {
		if len(existingState.rules) != len(newState.rules) {
			continue
		}
		thisIsTheItem := true
		for _, existingRule := range existingState.rules {
			ruleFound := false
			for _, newRule := range newState.rules {
				if areTheRulesTheSame(newRule, existingRule) {
					// Found the rule, breaks loop over the rules of the new automata
					ruleFound = true
					break
				}
			}
			if !ruleFound {
				thisIsTheItem = false
				break
			}
		}
		if thisIsTheItem {
			return &existingState, false
		}
	}
	// Have not found a valid State
	return &State{}, true
}

func areTheRulesTheSame(existingRule ItemRule, newRule ItemRule) bool {
	// Fucking hours OMG FUCK
	if existingRule.dot != newRule.dot || existingRule.rule.nonTerminal != newRule.rule.nonTerminal || len(newRule.rule.production) != len(existingRule.rule.production) {
		return false
	}

	for i, s := range existingRule.rule.production {
		if newRule.rule.production[i] != s {
			return false
		}
	}
	return true
}

func (automata *SLR_automata) Print() {
	fmt.Println()
	fmt.Println()
	for _, state := range automata.states {
		fmt.Println()
		fmt.Println()
		fmt.Print("State ")
		fmt.Println(state.id)
		for _, r := range state.rules {
			fmt.Print(r.rule.nonTerminal + " -> ")
			fmt.Print(r.rule.production)
			fmt.Println(r.dot)
		}
		fmt.Println("Transitions:")
		for input, endStateId := range state.transitions {
			fmt.Print("-> ")
			b := endStateId
			fmt.Print(b)
			fmt.Println(" with " + input)
		}
	}
}
