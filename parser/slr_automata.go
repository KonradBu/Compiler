package parser

type SLR_automata struct {
	items       []Item
	start_state  int
	nonTerminals map[string]bool
	terminals    map[string]bool
}

type Item struct {
	id          int
	rules       []Rule
	dots        map[*Rule]int
	transitions map[string]Item
}

type DotRule struct {
	dot         int
	nonTerminal string
	production  []string
}

func (automata *SLR_automata) CreateSLRTable(grammar *Grammar) *SLR_parsing_Table {
	table := makeSlrParsingTable()

	for _, item := range automata.items {
		for _, rule := range item.rules {
			dot := item.dots[&rule]
			switch {
			case automata.nonTerminals[rule.production[dot]]:
				// The dot is before a non terminal
				// Goto from the current item with the non terminal into the item consuming the current non terminal
				table.AddGoTo(item.id, rule.nonTerminal, item.transitions[rule.production[dot]].id)
			case automata.terminals[rule.production[dot]]:
				// The dot is before a terminal
				next := rule.production[dot]
				table.AddAction(item.id, rule.production[dot], "Shift", item.transitions[next].id)
			case len(rule.production) == dot:
				// The dot is at the end of the production
				for _, terminal := range grammar.follow[rule.nonTerminal] {
					if terminal == "$" {
						table.AddAction(item.id, "$", "Accept", 0)
					} else {
						table.AddAction(item.id, terminal, "Reduce", item.transitions[rule.nonTerminal].id)
					}
				}
			}
		}
	}
	return table
}

func (grammar *Grammar) CreateSLRAutomata() *SLR_automata {
	automata := new(SLR_automata)
	var startRule Rule
	for _, r := range grammar.rules {
		if r.nonTerminal == "S" {
			startRule = r
		}
	}
	startState := automata.makeState([]Rule{startRule}, []int{0}, grammar)
	automata.start_state = 0
	automata.addGotoRecursive(startState, grammar)
	return automata
}

func (automata *SLR_automata) makeState(rules []Rule, dots []int, grammar *Grammar) *Item{
	newItem := new(Item)
	newItem.dots = make(map[*Rule]int)
	newItem.transitions = make(map[string]Item)
	newItem.rules = rules
	for i, rule := range rules{
		newItem.dots[&rule] = dots[i]
	}
	newItem.addClosure(grammar, )
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
		nt := rule.production[item.dots[&rule]]
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

func (automata *SLR_automata) addGotoRecursive(item *Item, grammar *Grammar){
	for _, r := range item.rules{
		if item.dots[&r] < len(r.production){
			automata.Goto(item, grammar, r.production[item.dots[&r]])
		}
	}
}

func (automata *SLR_automata) Goto(item *Item, grammar *Grammar, symbol string) {
	var newItem Item
	rules := []Rule{}
	dots := []int{}
	for _, r := range item.rules {
		if item.dots[&r] < len(r.production) && r.production[item.dots[&r]] == symbol{
			dot := item.dots[&r] + 1
			rules = append(rules, r)
			dots = append(dots, dot)
		}
	}
	newItem = *automata.makeState(rules, dots, grammar)
	refItem, itemDoesNotExist := automata.itemDoesNotExist(item)
	if itemDoesNotExist{
		automata.items = append(automata.items, newItem)
		refItem = newItem
		automata.addGotoRecursive(item, grammar)
	}
	item.transitions[symbol] = refItem
}



func (automata *SLR_automata) itemDoesNotExist(item *Item) (Item, bool) {
	errorItem := new(Item)
	for _, existingItem := range automata.items {
		itemsAreTheSame := true
		for _, existingRule := range existingItem.rules {
			ruleIsInNewAutomata := false
			for _, newRule := range item.rules {
				if areTheRulesTheSame(existingRule, existingItem.dots[&existingRule], newRule, item.dots[&newRule]) {
					ruleIsInNewAutomata = true
					// Breaks the check over all the rules
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
	return *errorItem, true
}

func areTheRulesTheSame(existingRule Rule, existingDot int, newRule Rule, newDot int) bool {
	if existingDot != newDot || existingRule.nonTerminal != existingRule.nonTerminal {
		return false
	}

	for i, s := range existingRule.production {
		if newRule.production[i] != s {
			return false
		}
	}
	return true
}