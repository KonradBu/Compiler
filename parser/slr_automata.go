package parser

type SLR_automata struct{
	states []SLR_automata_state
	start_state int
	nonTerminals map[string]bool
	terminals map[string]bool
}

type SLR_automata_state struct{
	id int
	rules []Rule
	dots map[*Rule]int
	transitions map[string] SLR_automata_state
}

type DotRule struct{
	dot int
	nonTerminal string
	production []string
}

func(automata *SLR_automata) CreateSLRTable(grammar *Grammar) *SLR_parsing_Table{
	table := MakeSLRTable()

	for _, state := range automata.states {
		for _, rule := range state.rules{
			dot := state.dots[&rule]
			switch{
			case automata.nonTerminals[rule.production[dot]]:
				// The dot is before a non terminal
				// Goto from the current state with the non terminal into the state consuming the current non terminal 
				table.AddGoTo(state.id, rule.nonTerminal, state.transitions[rule.production[dot]].id)
			case automata.terminals[rule.production[dot]]:
				// The dot is before a terminal
				next := rule.production[dot] 
				table.AddAction(state.id, rule.production[dot], "Shift", state.transitions[next].id)
			case len(rule.production) == dot:
				// The dot is at the end of the production
				for _, terminal := range grammar.follow[rule.nonTerminal]{
					if terminal == "$"{
						table.AddAction(state.id, "$", "Accept", 0)
					} else {
						table.AddAction(state.id, terminal, "Reduce", state.transitions[rule.nonTerminal].id)
					}
				}
			}
		}
	}
	return table
}

func (grammar *Grammar) CreateSLRAutomata() *SLR_automata{
	automata := new(SLR_automata)
	startState := new(SLR_automata_state)
	automata.states = append(automata.states, *startState)
	automata.start_state = 0
	grammar.addClosure(startState)
	for _, rule := range grammar.rules {
		dot := startState.dots[&rule] 
		if  dot < len(rule.production){
			grammar.Goto(startState, rule.production[dot])
		}
	}
	return automata
}

func (automata *SLR_automata) addStates(){
	newState := new(SLR_automata_state)


	if automata.stateDoesNotExist(newState){
		automata.states = append(automata.states, *newState)
	}else{
		
	}
} 

func (grammar *Grammar) addStateRecurivly(map[*SLR_automata_state]bool){

}

func (automata *SLR_automata) stateDoesNotExist (state *SLR_automata_state) bool {
	for _, existingState := range automata.states{
		statesAreTheSame := true
		for _, existingRule := range existingState.rules{
			ruleIsInNewAutomata := false
			for _, newRule := range state.rules{
				if areTheRulesTheSame (existingRule, existingState.dots[&existingRule], newRule, state.dots[&newRule]){
					ruleIsInNewAutomata = true
					// Breaks the check over all the rules
					break
				}
			}
			if !ruleIsInNewAutomata {
				statesAreTheSame = false
				break
			}
		}
		if statesAreTheSame{
			return true
		}
	}
	return true
}

func areTheRulesTheSame (existingRule Rule, existingDot int, newRule Rule,newDot int) bool{
	//TODO
	return false
}


func (grammar *Grammar) addClosure(state *SLR_automata_state){
	done := make(map[string]bool)
	grammar.addClosureRecursive(state, done)
}

func (grammar *Grammar) addClosureRecursive (state *SLR_automata_state, done map[string]bool){
	changed := false
	for _, rule := range state.rules{
		nt := rule.production[state.dots[&rule]]
		if !done[nt]{
			state.rules = append(state.rules, grammar.closure[nt]...)
			for _, rule := range grammar.closure[nt]{
				state.dots[&rule] = 0
			}
			changed = true
			done[nt] = true
		}
	}
	if changed {
		grammar.addClosureRecursive(state, done)
	}
}

func (grammar *Grammar) Goto(state *SLR_automata_state, symbol string) {

}