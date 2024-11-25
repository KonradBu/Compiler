package parser

func (grammar *Grammar) getNullable() []string{
	var nullable []string
	for _, rule := range grammar.rules{
		// Ist die Regel eine Epsilon Transition?
		if rule.production[0] == " "{
			nullable = append(nullable, rule.nonTerminal)
		}
	}
	
	return nullable
}

func (grammar *Grammar) getFirst() map[string][]string{

}

func (grammar *Grammar) getFollow()map[string][]string{

}
