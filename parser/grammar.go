package parser

import (
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

func (grammar *Grammar) FIRST() map[string][]string {
	firstMap := make(map[string][]string)
	grammar.firstrecursive(grammar.start, firstMap)
	return firstMap
}

// Only works without epsilon and only specific kinds of left recursion ~ only with loop length == 1 :))
func (grammar *Grammar) firstrecursive(input string, firstMap map[string][]string) {
	_, ok := firstMap[input]
	if !ok {
		firstMap[input] = []string{}
	} else {
		return
	}
	for _, r := range grammar.rules {
		if r.nonTerminal == input {

			if isNT(r.production[0]) {
				grammar.firstrecursive(r.production[0], firstMap)
				for _, s := range firstMap[r.production[0]] {
					if contains(firstMap[input], s) == -1 {
						firstMap[input] = append(firstMap[input], s)
					}
				}
			} else {
				if contains(firstMap[input], r.production[0]) == -1 {
					firstMap[input] = append(firstMap[input], r.production[0])
				}
			}
		}
		for _, p := range r.production {
			if isNT(p) {
				if p != r.nonTerminal {
					grammar.firstrecursive(p, firstMap)
				}
			}
		}
	}
}

func (grammar *Grammar) FOLLOW(first map[string][]string) map[string][]string {
	followMap := make(map[string][]string)
	grammar.recursiveFollow("S", followMap, first)
	return followMap
}

func (grammar *Grammar) recursiveFollow(input string, followMap map[string][]string, first map[string][]string) {
	if followMap[input] == nil {
		followMap[input] = []string{}
	} else {
		return
	}
	if input == "S" {
		followMap[input] = []string{"$"}
	}
	for _, rule := range grammar.rules {
		for i, symbol := range rule.production {
			if symbol == input {
				if i == len(rule.production)-1 {
					grammar.recursiveFollow(rule.nonTerminal, followMap, first)
					for _, newEntry := range grammar.follow[rule.nonTerminal] {
						if contains(followMap[input], newEntry) == -1 {
							followMap[input] = append(followMap[input], newEntry)
						}
					}
				} else {
					next := rule.production[i+1]
					if isNT(next) {
						for _, newEntry := range first[next] {
							if contains(followMap[input], newEntry) == -1 {
								followMap[input] = append(followMap[input], newEntry)
							}
						}
					} else {
						if contains(followMap[input], next) == -1 {
							followMap[input] = append(followMap[input], next)
						}
					}
				}
				if isNT(symbol) {
					grammar.recursiveFollow(symbol, followMap, first)
				}
			}
		}
	}
}
