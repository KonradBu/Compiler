package parser

import (
	"compiler/lexer"
	"fmt"
)


type ParseTree struct{
	identifier string
	leaves *ParseTree
}


func Parse(path string, test bool) {
	
	tokenChannel := make(chan lexer.Token)
	go lexer.Lex(path, tokenChannel)

	fmt.Println("Started Parsing...")
	linecount := 0

	slrTable, grammar := createParser(test)

	stack := makeStack("0")

	accepts := false
	for true {
		token := lexer.GetNext(tokenChannel)

		if token == nil{
			if accepts{
				break
			}else{
				panic("File Ended. Unnexpected Symbol " + token.Identifier)
			}
		}
		if token.Identifier == "LINE"{
			linecount++
			continue
		}

		// Do only once, unless reduce is found
		for i := 0; i < 1; i++ { 
			stackVal:= *stack.Val
			res, err := slrTable.GetAction(stackVal.(int), token.Identifier)
			if err != nil{
				panic("Parsing Error. Cannot work with the symbol:  " + string(stackVal.(int)) + " at line " + string(linecount))
			}
			switch res.actionType{
			case "Shift":
				stack.add(token)
				stack.add(res.value)
			case "Reduce":
				// Redo the loop ~ Dont get another input symbol
				i--
				// Get the Rule that we reduce by
				reductionRule := grammar.rules[res.value]
				for range reductionRule.production{
					stack.pop()
					stack.pop()
				}
				stack.add(reductionRule.production)
				stateBefore := *stack.Next.Val
				gotoVal, err := slrTable.GetGoto(stateBefore.(int), reductionRule.nonTerminal)
				if err != nil{				
					panic("Parsing Error. Cannot work with the symbol:  " + string(stackVal.(int)) + " at line " + string(linecount))
				}
				stack.add(gotoVal.val)
			case "Accept":
				accepts = true
			}
		}
	}
	if accepts{
		fmt.Println("Parser finished")
		fmt.Println()
	}else{
		next := slrTable.getNextExpectedTokens(stack.pop().(int))
		panic("File ended unnexpectedly. Still waiting for tokens. Possible Token:" + next)
	}
}

func createParser(test bool) (*SLR_parsing_Table, *Grammar){
	// Only done for test case
	rules:= defGrammar(test)
	grammar := MakeGrammar(rules, "S")
	// Done
	grammar.Augment()
	// Not done??
	grammar.CalcFollow()
	// Done
	grammar.CalcClosure()
	// Done
	// Done
	automata := grammar.CreateSLRAutomata()
	// Done
	table := automata.CreateSLRTable(grammar)
	return table, grammar
}

func AddGrammar() *Grammar{
	
}
