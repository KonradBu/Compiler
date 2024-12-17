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

func Parse(path string, test bool) bool {
	tokenChannel := make(chan lexer.Token)
	slrTable, grammar := createParser(test)

	go lexer.Lex(path, tokenChannel)

	//fmt.Println("Started Parsing...")
	linecount := 0

	stack := makeStack(0)

	accepts := false
	for true {
		token := lexer.GetNext(tokenChannel)

		if token.Identifier == "" {
			if accepts {
				break
			} else {
				parseError(*token, linecount, *stack, slrTable)
			}
		}
		if token.Identifier == "LINE" {
			var err error
			linecount, err = strconv.Atoi(token.Value.(string))
			if err != nil {
				fmt.Print("Parser Error, Linecount is not int")
				return false
			}
			continue
		}

		// Do only once, unless reduce is found
		for i := 0; i < 1; i++ {
			stackVal := stack.peek().(*any)
			res, err := slrTable.GetAction((*stackVal).(int), token.Identifier)
			if err != nil {
				parseError(*token, linecount, *stack, slrTable)
				return false
			}
			switch res.actionType {
			case "Shift":
				stack.add(token)
				stack.add(res.value)
			case "Reduce":
				// Redo the loop ~ Dont get another input symbol
				i--
				// Get the Rule that we reduce by
				reductionRule := grammar.rules[res.value]
				for range reductionRule.production {
					stack.pop()
					stack.pop()
				}
				stateBefore := stack.peek().(*any)
				gotoVal, err := slrTable.GetGoto((*stateBefore).(int), reductionRule.nonTerminal)
				if err != nil {
					parseError(*token, linecount, *stack, slrTable)
					return false
				}
				stack.add(reductionRule.nonTerminal)
				stack.add(gotoVal.val)
			case "Accept":
				accepts = true
			}
		}
	}
	if accepts {
		//fmt.Println("Parser finished")
		accept()
	} else {
		parseError(lexer.Token{}, linecount, *stack, slrTable)
		return false
	}
	return true
}

func createParser(test bool) (*SLR_parsing_Table, *Grammar) {
	// Only done for test case
	rules := defGrammar(test)
	grammar := MakeGrammar(rules, "START")
	// Done
	grammar.Augment()
	// Not done??
	first := grammar.FIRST()

	// TODO
	// There is something wrong with first, no time to fix it yet
	first["S"] = append(first["S"], "namespace")
	first["START"] = append(first["START"], "namespace")
	//fmt.Println("First bodge still in place")
	//PrintFirst(first)
	follow := grammar.FOLLOW(first)
	//PrintFollow(follow)
	grammar.follow = follow
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

func parseError(token lexer.Token, linecount int, stack Stack, table *SLR_parsing_Table) {
	lineString := strconv.Itoa(linecount)
	state := stack.peek().(*any)
	next := table.getNextExpectedTokens((*state).(int))

	conv, err := token.Value.(int)
	if err && conv != 0 {
		next = append(next, fmt.Sprintf("%v", token.Value))
	}

	nextString := formatNext(next)

	if token.Identifier == "$" {
		fmt.Println("Unexpected end of file reached. At line: " + lineString + ".\nExpecting: " + nextString)
		return
	}
	unexpected := formatToken(token)

	fmt.Println("Syntax Error. Unexpected: \"" + unexpected + "\" at line " + lineString + ".\nExpecting: " + nextString)
}

func formatNext(next []string) string {
	returnstring := ""
	for _, n := range next {
		nextString := ""
		switch n {
		case "logicaloperator":
			nextString = "==, ||, &&, >=, <=, =="
		case "multoperator":
			nextString = "*, /"
		case "unaryoperator":
			nextString = "+, -"
		case "name":
			// Nothing :)
		case "intliteral":
			nextString = "num"
		case "boolliteral":
			nextString = "bool"
		default:
			nextString = n
		}
		returnstring += nextString + "," + " "
		nextString = ""
	}
	return returnstring[:len(returnstring)-2]
}

func formatToken(token lexer.Token) string {
	switch token.Identifier {
	case "name", "logicaloperator", "multoperator", "unaryoperator", "boolliteral":
		return token.Value.(string)
	case "intliteral":
		return strconv.Itoa(token.Value.(int))
	case "stringliteral":
		return "string"
	default:
		return token.Identifier
	}
}

func accept() {
	fmt.Println()
	fmt.Println()
	fmt.Println("###############")
	fmt.Println()
	fmt.Println("ACCEPTED")
	fmt.Println()
	fmt.Println("###############")
	fmt.Println()
	fmt.Println()
}
