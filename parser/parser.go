package parser

import (
	"compiler/lexer"
)

func Parse(path string) {
	go lexer.Lex(path)

	for true {
		token, err := lexer.GetNext()
		if err != nil{
			break
		}

		// Do something with the token
	}
}
