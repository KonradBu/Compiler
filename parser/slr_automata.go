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
	dots map[int]int
	transitions map[string] SLR_automata_state
}

func(automata *SLR_automata) CreateSLRTable(grammar *Grammar, grammarFollow *GrammarFollow) *SLR_parsing_Table{
	table := MakeSLRTable()

	for _, state := range automata.states {
		for ruleid, rule := range state.rules{
			dot := state.dots[ruleid]
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
				for _, terminal := range grammarFollow.follow[rule.nonTerminal]{
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

func (grammar *Grammar) CreateSLRAutomata(follow GrammarFollow) *SLR_automata{

}
