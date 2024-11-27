package parser

type Grammar struct{
	start string
	nonTerminals []string
	terminals []string
	rules []Rule
}


// Epsilon -> " "
type Rule struct{
	nonTerminal string
	production []string
}

func makeStandardGrammar() *Grammar{
	newGrammar := new(Grammar)
	newGrammar.start = "S"
	newGrammar.nonTerminals = append(newGrammar.nonTerminals, "S")
	return newGrammar
}

/*
func (grammar *Grammar) addRule(rules []Rule){
	for _, r := range rules{
		
		}
}
*/