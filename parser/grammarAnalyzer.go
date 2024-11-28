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

// Requires NULLABLE to be done!
func (grammar *Grammar) GetFirst(){
	firstMap := *new(map[string][]string)
	nonTerminalMap := grammar.makeNonTerminalMap()
	grammar.recursiveFirst(grammar.start, firstMap, nonTerminalMap)
	grammar.first = firstMap
}

// This function does not take into account non terminals that cannot be reached from the start symbol
func (grammar *Grammar)recursiveFirst(nt string, firstMap map[string][]string, nonTerminalMap map[string][]Rule){
	dependancies := []string{}
	// Create the array for this non terminal
	firstMap[nt] = []string{}
	
	// For every rule associated with the starting non terminal...
	for _, rule := range nonTerminalMap[nt]{
		// for every symbol of the start...
		for _, prod := range rule.production{

			
			retString := ""
			// Is said symbol terminal??
			for _, t := range grammar.terminals{
				if t == prod{
					retString = prod
					// Breaks the loop over the terminals
					break
				}
			}

			// The symbol is a terminal!
			if retString != ""{
				firstMap[nt] = append(firstMap[nt], retString)
				// This breaks the loop over the production parts ~ jumps to the next rule
				break
			}
			
			//If we have reached this point, the symbol has to be a non terminal
			
			firstRec, ok := firstMap[prod]
			if ok{
				firstMap[nt] = firstRec
			}else{
				// We have not searched the non terminal yet
				dependancies = append(dependancies, prod)
				// We have to check the 
				grammar.recursiveFirst(prod, firstMap, nonTerminalMap)
			}
			
			// Check if the first symbol is nullable
			if !grammar.nullable[prod]{		
				// This breaks the loop over the production parts ~ jumps to the next rule
				break
			}
			// It isnt -> we have to keep looping over the symbols of this production rule until we find a non nullable / there are no more symbols left
		}
	}
	// Append the first map of the dependancies
	for _, d := range dependancies{
		firstMap[nt] = append(firstMap[nt], firstMap[d]...)
	}
}


// Requires FIRST to be done
func (grammar *Grammar) GetFollow(){
	followMap := *new(map[string][]string)
	followMap[grammar.start] = []string{"$"}
	nonTerminalMap := grammar.makeNonTerminalMap()
	for _, nt := range grammar.nonTerminals{
		grammar.calcFollow(nt, followMap, nonTerminalMap)
	}
	grammar.follow = followMap
}


func (grammar *Grammar)calcFollow(nt string, followMap map[string][]string, nonTerminal map[string][]Rule){
	followMap[nt] = []string{}
	// For every rule associated with the starting non terminal...
	for _, rule := range grammar.rules{
		// for every symbol of the start...
		for i, prod := range rule.production{
			if prod == nt{
				next := rule.production[i + 1]
				for _, t := range grammar.terminals{
					if t == prod{
						followMap[nt] = append(followMap[nt], next)	
					}
					// Break the loop and skip to the next rule
					break
				}
				// -> it has to be a non terminal
				for j := i + 1; j < len(rule.production); j++{
					followMap[nt] = append(followMap[nt], grammar.first[next]...)
					if !grammar.nullable[rule.production[j]]{
						i = j
						break
					}
				}
			}
		}
	}
}


// Makes a map from non terminal -> all the possible productions
func (grammar *Grammar) makeNonTerminalMap() map[string][]Rule{
	nonTerminalMap := *new(map[string][]Rule)
	for _, r := range grammar.rules{
		_, ok := nonTerminalMap[r.nonTerminal]
		if !ok{
			nonTerminalMap[r.nonTerminal] = []Rule{}
		}
		nonTerminalMap[r.nonTerminal] = append(nonTerminalMap[r.nonTerminal], r) 
	}
	return nonTerminalMap
}