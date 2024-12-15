package parser

import "fmt"

/*
Step 1:
	Create Augmented Grammar
Step 2:
	Calculate the follow of all the rules in the Grammar
Step 3:
	Canonical LR(0) collection ~ Create Automata
		closure(non terminal, grammar)
			Starting With S'
			If "."" is followed by a non terminals:
			-> Write all the Rules from that non terminal into the state
			-> "." At the left side of the rule
			-> Recursivly ad closure of all the new rules
		goto(Current State, Symbol, Grammar)
			For all Rules:
				If "." is not at the end ->
					Shift "." by 1
					-> Make new state
					-> If the Symbol before the "." was a non Terminal:
					-> If there are other rules with the non terminal before the dot:
						Get included in the new state
					-> Closure of all the rules in the new state
					We need to check if we have created this state already
						~ If another state has all the same Grammar rules, with . at the same position
Step 4:
	Construct SLR Parsing Table
		Columns Split into 2: Action and Goto
		Lines are the different states/ items
			Action: Terminal Symbols in the Grammar + $
			Goto: Non Terminals in the Grammar
			In the Table we will fill out the different actions, that parser can take.
			Shift, Reduce, Accept and Error

		Shift:
			Go through all of our states:
				If the rule is followed by a terminal symbol:
					Shift into the state from goto + the terminal symbol

		Goto:
			For all of out states:
				If . is followed by a non terminal
				-> Note the state that goto (current state, non terminal) goes to into the goto table

		Reduce:
			If "." is at the end of a prediction:
				Write into table: Reduce + The ID of the rule with the . at the end
				At the terminals, which are in the FOLLOW of the Non terminal that starts this rule

		Accept:
			If "." follows the Rule S' -> S: write accept into the table for $

Step 5:
	Check if there are any states, where multiple rules could be used -> The grammar fails

Step 6:
	Parsing the input usung the SLR Parsing Table


	New Table:
		Stack: Starts with 0
		Input: Starts with Input and String + Dollar
		Action:

	Parsing:
		Look at the symbol in the table: State on top of stack + the first symbol in the input string
		If shift:
			Ad symbol and state on top of stack
			Remove Symbol from input
		If reduce:
			Replace the Items on top of the stack that correspond to the rule we reduce by,
			by the left side of the rule
			If the items dont correspond -> Error
			We then need to look at the goto table for that input symbol on the left side + the state before that on the stack
			-> Put the goto state on top of the stack
			Dont change the input buffer
		If empty entry:
			Error
		If accept + stack is emtpy:
			Accept :)
		If accept + stack is not empty:
			Error
*/

func (grammar *Grammar) Augment() {
	oldStart := grammar.start
	newStart := "S"
	grammar.start = newStart
	grammar.nonTerminals = append(grammar.nonTerminals, "S")
	grammar.AddRule(newStart, []string{oldStart})
}

func (grammar *Grammar) CalcFollow() {
	first := grammar.FIRST()
	PrintFirst(first)
	follow := make(map[string][]string)
	for _, nt := range grammar.nonTerminals {
		follow[nt] = grammar.FOLLOW(nt, first)
	}
	grammar.follow = follow
}

func PrintFirst(first map[string][]string){
	fmt.Println("FIRST:")	
	for nt, t := range first{
		fmt.Print(nt)
		fmt.Print(": ")
		for _, n := range t{
			fmt.Print(n)
			fmt.Print(" ")
		}
		fmt.Println()
	} 
}

func (grammar *Grammar) CalcClosure() {
	closure := make(map[string][]Rule)
	for _, nt := range grammar.nonTerminals {
		for _, rule := range grammar.rules {
			if rule.nonTerminal == nt {
				closure[nt] = append(closure[nt], rule)
			}
		}
	}
	grammar.closure = closure
}

func contains(arr []string, target string) int {
	for i, a := range arr {
		if a == target {
			return i
		}
	}
	return -1
}
