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
		token, err := lexer.GetNext(tokenChannel)
		if err != nil {
			break
		}

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
	grammar := AddGrammar()
	grammar.Augment()
	follow := grammar.CalcFollow()
	automata := grammar.CreateSLRAutomata(follow)
	table := automata.CreateSLRTable(grammar, &follow)
	return table
}

func AddGrammar() *Grammar{

}