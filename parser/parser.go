package parser

import (
	"compiler/lexer"
	"fmt"
	"strconv"
)

func Parse(path string) {
	tokenChannel := make(chan lexer.Token)
	go lexer.Lex(path, tokenChannel)

	fmt.Println("Started Parsing...")
	linecount := 0

	createParser()

	for true {
		token := lexer.GetNext(tokenChannel)

		switch token.Identifier {
		case "LINE":
			linecount, _ = strconv.Atoi(token.Value.(string))
		case "ERROR":
			panic("hihi " + string(linecount))
		default:
			fmt.Print(token.Identifier + " ")
			if token.Value != nil {
				fmt.Print(token.Value)
			}
			fmt.Println()
		}
		// Do something with the token
		
		
		fmt.Println("Parser finished")
		fmt.Println()
	}
}

func createParser() *SLR_parsing_Table{
	// Not done
	grammar := AddGrammar()
	// Done
	grammar.Augment()
	// Not done??
	grammar.CalcFollow()
	// Done
	grammar.CalcClosure()
	// Done
	automata := grammar.CreateSLRAutomata()
	// Done
	table := automata.CreateSLRTable(grammar)
	return table
	// Not done
	// parse()
}

func AddGrammar() *Grammar{
	
}