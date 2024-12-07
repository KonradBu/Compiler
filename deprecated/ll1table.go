package parser

type parsetable struct{
	startSymbol string
	nonterminals map[string]int
	terminals map[string]int
	
	// First nonterminals then terminals then rule
	table [][]string
}

func (table *parsetable) getRules(nonterminal string,terminal string) string{
	ntVal := table.nonterminals[nonterminal]
	tVal := table.terminals[terminal]

	return table.table[ntVal][tVal]
} 

func (table *parsetable) getStartSymbol() string{
	return table.startSymbol
}

func (table *parsetable) getNonTerminals() []string{
	var retArray []string
	for key, _ := range table.nonterminals{
		retArray = append(retArray, key)
	}
	return retArray
}

func (table *parsetable) getTerminals() []string{
	var retArray []string
	for key, _ := range table.terminals{
		retArray = append(retArray, key)
	}
	return retArray
}

func (table *parsetable) getEpsilon() []string{
	panic("")
}
