package parser

import (
	"compiler/lexer"
	"fmt"
	"strconv"
)

func Parse(path string) {
	tokenChannel := make(chan lexer.Token)
	go lexer.Lex(path, tokenChannel)

	throwAwayNext := false
	fmt.Println("Started Parsing...")
	linecount := 0

	for true {
		token, err := lexer.GetNext(tokenChannel)
		if err != nil {
			fmt.Println()
			fmt.Println("Lexer finished")
			fmt.Println("Parser finished")
			fmt.Println()
			break
		}

		if throwAwayNext {
			if token.Identifier != "NAME" {
				fmt.Println(token)
				lineErr := string(linecount)
				fmt.Println(lineErr)
				panic("PARSE ERROR: Keyword/ Symbol used as variable name. Token: " + token.Identifier)
			}
			throwAwayNext = false
			continue
		}

		switch token.Identifier {
		case "LINE":
			linecount, _ = strconv.Atoi(token.Value.(string))
		default:
			fmt.Print(token.Identifier + " ")
			if token.Value != nil {
				fmt.Print(token.Value)
			}
			fmt.Println()
		}
		// Do something with the token
	}
}
