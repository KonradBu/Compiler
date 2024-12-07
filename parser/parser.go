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


	/*
	# How to SLR Parse
	Init Grammar  + Augmentation
	
	*/ 

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
		
		// End of file
		if token.Identifier == "END"{
			break
		}
	}

	
	fmt.Println("Parser finished")
	fmt.Println()
}
