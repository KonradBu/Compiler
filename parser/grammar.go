package parser

import "unicode"

type Grammar struct{
	start string
	nonTerminals []string
	terminals []string
	rules []Rule
	follow map[string][]string
	closure map[string][]Rule
}

/*type GrammarFollow struct{
	nullable map[string]bool
	first map[string][]string
	follow map[string][]string
}

*/

// Epsilon -> " "
type Rule struct{
	nonTerminal string
	production []string
}

func (grammar *Grammar) addSymbol(s string){
	grammar.nonTerminals = append(grammar.nonTerminals, s)
}

func MakeRule (nonTerminal string, production []string) Rule{
	newRule := new(Rule)
	newRule.nonTerminal = nonTerminal
	newRule.production = append(newRule.production, production...)
	return *newRule
}

func (grammar *Grammar) AddRule (nonTerminal string, production []string) Rule{
	newRule := MakeRule(nonTerminal, production)
	grammar.rules = append(grammar.rules, newRule)
	return newRule
}

func MakeGrammar(rules []Rule, start string) *Grammar{
	newGrammar := new(Grammar)
	newGrammar.start = start
	newGrammar.rules = rules
	// Add Terminals/ Non terminals
	for _, r := range rules{
		newGrammar.nonTerminals = append(newGrammar.nonTerminals, r.nonTerminal)
		for _, s:= range r.production{
			if isNT(s){
				newGrammar.nonTerminals = append(newGrammar.nonTerminals, s)
			}else{
				newGrammar.terminals = append(newGrammar.terminals, s)
			}
		}
	}
	return newGrammar
}

func isNT (input string) bool{
	for _, r := range []rune(input){
		if !unicode.IsUpper(r){
			return false
		}
	}
	return true
}

func (grammar *Grammar) NULLABLE() *map[string]bool{
	// TODO: CALC NULLABLE
}

func (grammar *Grammar) FIRST(nonTerminal string, nullable map[string]bool) []string{
	// TODO: CALC FIRST
}

func (grammar *Grammar) FOLLOW(nonTerminal string, nullable map[string]bool, first map[string][]string) []string{
	// TODO: CALC FOLLOW
}


func makeStandardGrammar() *Grammar{
	newGrammar := new(Grammar)
	newGrammar.start = "S"
	newGrammar.nonTerminals = append(newGrammar.nonTerminals, "S")
	return newGrammar
}


