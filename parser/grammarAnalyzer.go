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

func (grammar *Grammar) GetNullable() []string{
	checked := new(map[string]bool)
	start := grammar.start
	nullables := make(chan string)
	
	search := func (channel chan any){

	} 
	
	grammar.grammarSearch(start, *checked, nullables, search)
	/* 
	Nimmt alle Regeln vom Startsymbol
	-> Ist Regel == null -> s nullable
	-> Fals terminal enthält -> nicht nullable

	-> geht recursivly durch alle non Terminals in den start regeln durch 
	*/

	// add from channel
}

// Does not check unnreachables!!
func (grammar *Grammar) grammarSearch(start string, checked map[string]bool, results chan string, search func(channel chan any)){
	if checked[start] {
		return
	}

	// Go over all the relevant production rules
	// If epislon -> is nullable
	// If terminal -> not nullable
	// If nt -> recursivly call
	for _, r := range grammar.rules{
		if r.nonTerminal == start{
			for _, t := range r.production{
				//Nullable
				if r.production == ""{
					nullables <- start
				}


				/*if r.production == 
				go grammar.nullableRecursive(checked, )*/
			}
		} 
	}
}


/* Nullable:
	Geht durch jedes nt und checkt ob es 



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