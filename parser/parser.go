package parser

import (
	"compiler/lexer"
	"fmt"
	"strconv"
)

type ParseTree struct {
	identifier string
	leaves     *ParseTree
}

func Parse(path string, test bool) {

	tokenChannel := make(chan lexer.Token)
	go lexer.Lex(path, tokenChannel)

	//fmt.Println("Started Parsing...")
	linecount := 0

	slrTable, grammar := createParser(test)

	stack := makeStack(0)

	accepts := false
	for true {
		token := lexer.GetNext(tokenChannel)

		if token.Identifier == "" {
			if accepts {
				break
			} else {
				_, val := stack.pop()
				next := slrTable.getNextExpectedTokens(val.(int))
				panic("File ended unnexpectedly. Still waiting for tokens. Possible Token:" + next)
			}
		}
		if token.Identifier == "LINE" {
			linecount++
			continue
		}

		// Do only once, unless reduce is found
		for i := 0; i < 1; i++ {
			stackVal := *stack.Val
			res, err := slrTable.GetAction(stackVal.(int), token.Value.(string))
			if err != nil {
				lineString := strconv.Itoa(linecount)
				panic("Parsing Error. Cannot work with the symbol: " + token.Identifier + " at line " + lineString)
			}
			switch res.actionType {
			case "Shift":
				stack = stack.add(token)
				stack = stack.add(res.value)
			case "Reduce":
				// Redo the loop ~ Dont get another input symbol
				i--
				// Get the Rule that we reduce by
				reductionRule := grammar.rules[res.value]
				for range reductionRule.production {
					stack, _ = stack.pop()
					stack, _ = stack.pop()
				}
				stack = stack.add(reductionRule.production)
				stateBefore := *stack.Next.Val
				gotoVal, err := slrTable.GetGoto(stateBefore.(int), reductionRule.nonTerminal)
				if err != nil {
					lineString := strconv.Itoa(linecount)
					panic("Parsing Error. Cannot work with the symbol:  " + token.Identifier + " at line " + lineString)
				}
				stack = stack.add(gotoVal.val)
			case "Accept":
				accepts = true
			}
		}
	}
	if accepts {
		//fmt.Println("Parser finished")
		accept()
	} else {
		_, val := stack.pop()
		next := slrTable.getNextExpectedTokens(val.(int))
		panic("File ended unnexpectedly. Still waiting for tokens. Possible Token:" + next)
	}
}

func createParser(test bool) (*SLR_parsing_Table, *Grammar) {
	// Only done for test case
	rules := defGrammar(test)
	grammar := MakeGrammar(rules, "START")
	// Done
	grammar.Augment()
	// Not done??
	grammar.CalcFollow()
	// Done
	grammar.CalcClosure()
	// Done
	//fmt.Println(grammar.closure)
	automata := grammar.CreateSLRAutomata()
	//automata.Print()
	// Done
	table := automata.CreateSLRTable(grammar)
	//fmt.Println("Table: ")
	//table.PrintTable(grammar)
	return table, grammar
}

func accept() {
	fmt.Println()
	fmt.Println()
	fmt.Println("###############")
	fmt.Println()
	fmt.Println("ACCEPTED")
	fmt.Println()
	fmt.Println("###############")
}
