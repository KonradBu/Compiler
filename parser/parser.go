package parser

import (
	"compiler/lexer"
	"fmt"
)

func Parse(path string) {
	tokenChannel := make(chan lexer.Token)
	go lexer.Lex(path, tokenChannel)
	for true {
		token, err := lexer.GetNext(tokenChannel)
		if err != nil {
			break
		}

		linecount := 0
		switch token.Identifier {
		case "LINE":
			linecount = int(token.Value.(int))
		case "NAMESPACE":
			//
		case "STATCIC":
			//
		default:
			fmt.Print(token.Identifier + " ")
			if token.Value != nil {
				fmt.Print(token.Value)
			}
			fmt.Println(" " + string(linecount))
		}
		// Do something with the token
	}
}
