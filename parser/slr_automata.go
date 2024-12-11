package parser

import "fmt"

type SLR_automata struct {
	items        []Item
}

type Item struct {
	id          int
	rules       []Rule
	dots        map[*Rule]int
	transitions map[string]Item
}


func (grammar *Grammar) CreateSLRAutomata() *SLR_automata {
	automata := new(SLR_automata)
	var startRule Rule
	for _, r := range grammar.rules {
		if r.nonTerminal == "S" {
			startRule = r
			break
		}
	}
	startItem := automata.makeItem([]Rule{startRule}, []int{0}, grammar)
	automata.addGotoRecursive(startItem, grammar)
	return automata
}

func (automata *SLR_automata) makeItem(rules []Rule, dots []int, grammar *Grammar) *Item {
	newItem := new(Item)
	newItem.dots = make(map[*Rule]int)
	newItem.transitions = make(map[string]Item)
	newItem.rules = rules
	for i, rule := range rules {
		newItem.dots[&rule] = dots[i]
	}
	newItem.addClosure(grammar)
	automata.items = append(automata.items, *newItem)
	return newItem
}

func (item *Item) addClosure(grammar *Grammar) {
	done := make(map[string]bool)
	grammar.addClosureRecursive(item, done)
}

func (grammar *Grammar) addClosureRecursive(item *Item, done map[string]bool) {
	changed := false
	for _, rule := range item.rules {
		var nt string
		if item.dots[&rule] < len(rule.production){
			nt = rule.production[item.dots[&rule]]
		}else{
			continue
		}
		if !done[nt] {
			item.rules = append(item.rules, grammar.closure[nt]...)
			for _, rule := range grammar.closure[nt] {
				item.dots[&rule] = 0
			}
			changed = true
			done[nt] = true
		}
	}
	if changed {
		grammar.addClosureRecursive(item, done)
	}
}

func (automata *SLR_automata) addGotoRecursive(item *Item, grammar *Grammar) {
	for _, r := range item.rules {
		if item.dots[&r] < len(r.production) {
			automata.Goto(item, grammar, r.production[item.dots[&r]])
		}
	}
}

func (automata *SLR_automata) Goto(item *Item, grammar *Grammar, symbol string) {
	var newItem Item
	rules := []Rule{}
	dots := []int{}
	for _, r := range item.rules {
		if item.dots[&r] < len(r.production) && r.production[item.dots[&r]] == symbol {
			newdot := item.dots[&r] + 1
			rules = append(rules, r)
			dots = append(dots, newdot)
		}
	}
	newItem = *automata.makeItem(rules, dots, grammar)
	otherItem, itemDoesNotExist := automata.itemDoesNotExist(item)
	if itemDoesNotExist {
		automata.items = append(automata.items, newItem)
		item.transitions[symbol] = newItem
		automata.addGotoRecursive(item, grammar)
	}else{
		item.transitions[symbol] = otherItem
	}
}

func (automata *SLR_automata) itemDoesNotExist(newitem *Item) (Item, bool) {
	// Go through every item already in the automata.
	// Go through every rule in that automata
	// If that rule is not in the other automata skip this automata
	// If every rule is in that state -> return false
	// If the no state in the automata is the same, return true

	for _, existingItem := range automata.items {
		
		itemsAreTheSame := true
		
		for _, existingRule := range existingItem.rules {
			ruleIsInNewAutomata := false
			for _, newRule := range newitem.rules {
				if areTheRulesTheSame(existingRule, existingItem.dots[&existingRule], newRule, newitem.dots[&newRule]) {
					ruleIsInNewAutomata = true
					// Breaks the check over all the rules for one specific rule in the old old automata
					break
				}
			}
			if !ruleIsInNewAutomata {
				itemsAreTheSame = false
				break
			}
		}
		
		if itemsAreTheSame {
			return existingItem, false
		}
	}
	return Item{}, true
}

func areTheRulesTheSame(existingRule Rule, existingDot int, newRule Rule, newDot int) bool {
	if existingDot != newDot || existingRule.nonTerminal != existingRule.nonTerminal || len(newRule.production) != len(existingRule.production) {
		return false
	}

	for i, s := range existingRule.production {
		if newRule.production[i] != s {
			return false
		}
	}
	return true
}

func (automata *SLR_automata) Print() {
	fmt.Println()
	fmt.Println()
	for i, item := range automata.items {
		fmt.Print("Item ")
		fmt.Println(i)
		for _, r := range item.rules {
			fmt.Print(r.nonTerminal + " -> ")
			fmt.Print(r.production)
			fmt.Println(item.dots[&r])
		}
	}
}
