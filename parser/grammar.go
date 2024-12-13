package parser

import (
	"fmt"
	"unicode"
)

type Grammar struct {
	start        string
	nonTerminals []string
	terminals    []string
	rules        []Rule
	follow       map[string][]string
	closure      map[string][]Rule
}

/*
	type GrammarFollow struct{
		nullable map[string]bool
		first map[string][]string
		follow map[string][]string
	}
*/
type Rule struct {
	nonTerminal string
	production  []string
}

func (grammar *Grammar) addSymbol(s string) {
	grammar.nonTerminals = append(grammar.nonTerminals, s)
}

func MakeRule(nonTerminal string, production []string) Rule {
	newRule := new(Rule)
	newRule.nonTerminal = nonTerminal
	newRule.production = append(newRule.production, production...)
	return *newRule
}

func (grammar *Grammar) AddRule(nonTerminal string, production []string) Rule {
	newRule := MakeRule(nonTerminal, production)
	grammar.rules = append(grammar.rules, newRule)
	return newRule
}

func MakeGrammar(rules []Rule, start string) *Grammar {
	newGrammar := new(Grammar)
	newGrammar.start = start
	newGrammar.rules = rules
	// Add Terminals/ Non terminals
	for _, r := range rules {
		if contains(newGrammar.nonTerminals, r.nonTerminal) == -1 {
			newGrammar.nonTerminals = append(newGrammar.nonTerminals, r.nonTerminal)
		}
		for _, s := range r.production {
			if isNT(s) {
				if contains(newGrammar.nonTerminals, s) == -1 {
					newGrammar.nonTerminals = append(newGrammar.nonTerminals, s)
				}
			} else {
				if contains(newGrammar.terminals, s) == -1 {
					newGrammar.terminals = append(newGrammar.terminals, s)
				}
			}
		}
	}
	return newGrammar
}

func isNT(input string) bool {
	for _, r := range []rune(input) {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func (grammar *Grammar) FIRST(nonTerminal string) []string {
	// TODO: CALC FIRST
	return []string{"id", "("}
}

/*
// Does NOT deal with epsilon Transitions
func (grammar *Grammar) SetFirst(){
	firstMap := *new(map[string][]string)
	nonTerminalMap := grammar.makeNonTerminalMap()
	grammar.recursiveFirst(grammar.start, firstMap, nonTerminalMap)
	grammar.first = firstMap
}

// This function does not take into account non terminals that cannot be reached from the start symbol
func (grammar *Grammar)recursiveFirst(nt string, firstMap map[string][]string){
	nonTerminalMap := grammar.nonTerminals
	dependancies := []string{}
	// Create the array for this non terminal
	firstMap[nt] = []string{}

	// For every rule associated with the starting non terminal...
	for _, rule := range nonTerminalMap[nt]{
		// for every symbol of the start...
		for _, prod := range rule.production{

			retString := ""
			// Is said symbol terminal??
			for _, t := range grammar.terminals{
				if t == prod{
					retString = prod
					// Breaks the loop over the terminals
					break
				}
			}

			// The symbol is a terminal!
			if retString != ""{
				firstMap[nt] = append(firstMap[nt], retString)
				// This breaks the loop over the production parts ~ jumps to the next rule
				break
			}

			//If we have reached this point, the symbol has to be a non terminal

			firstRec, ok := firstMap[prod]
			if ok{
				firstMap[nt] = firstRec
			}else{
				// We have not searched the non terminal yet
				dependancies = append(dependancies, prod)
				// We have to check the
				grammar.recursiveFirst(prod, firstMap, nonTerminalMap)
			}
			break
		}
	}
	// Append the first map of the dependancies
	for _, d := range dependancies{
		firstMap[nt] = append(firstMap[nt], firstMap[d]...)
	}
}

*/

func (grammar *Grammar) FOLLOW(nonTerminal string, first map[string][]string) []string {
	// TODO: CALC FOLLOW

	fmt.Println("FOLLOW not done")
	switch nonTerminal {
	case "S":
		return []string{"$"}
	case "E":
		return []string{")", "+", "$"}
	case "T":
		return []string{"+", "*", ")", "$"}
	case "F":
		return []string{"+", "*", ")", "$"}
	default:
		panic("Error Calculating Follow")
	}
}
