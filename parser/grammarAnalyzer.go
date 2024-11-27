package parser

func (grammar *Grammar) GetNullable() []string{
	var nullable []string
	nullableMap := make(map[string]bool)
	grammar.recursiveNullable(nullableMap, nullable)
	return nullable
}

func (grammar *Grammar) recursiveNullable(nullableMap map[string]bool, nullable []string){
	changed := false
	for _, rule := range grammar.rules{
		// Ist die Regel eine Epsilon Transition?
		if rule.production[0] == " "{
			nullable = append(nullable, rule.nonTerminal)
			nullableMap[rule.nonTerminal] = true
			changed = true
		}else{
			for _, s := range rule.production{
				if nullableMap[s]{
					nullable = append(nullable, rule.nonTerminal)
					nullableMap[rule.nonTerminal] = true
					changed = true
				}
			}
		}
	}
	if changed{
		grammar.recursiveNullable(nullableMap, nullable)
	}
}

/*func (grammar *Grammar)recursiveSearch(inputMap map[string]any,closure func(*Grammar)()(bool)){

}
*/

/*
func (grammar *Grammar) GetFirst(nullable []string) map[string][]string{
	var resMap map[string][]string
	for _, nt := range grammar.nonTerminals{
		var first []string
		var firstMap map[string]bool
		grammar.recursiveFirst(nt, first, firstMap)
		resMap[nt] = first
	}
	
}

func (grammar *Grammar)recursiveFirst(nt string, first []string, firstMap map[string]bool){

}

func (grammar *Grammar) GetFollow()map[string][]string{

}
*/