package lexer

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"unicode"
)

type Token struct {
	identifier  string
	value any
}

var tokenChannel chan Token

func GetNext() (*Token, error) {
	// Wait until the channel with tokens has a value inside
	select {
	case newToken, ok := <-tokenChannel:
		// If channel is close -> File is empty
		if !ok {
			return nil, errors.New("Lexer Error: File Ended")
		}
		return &newToken, nil
	}
}

func Lex(path string) {
	// Open File
	file, err := os.Open(path)
	if err != nil {
		panic("Lexer Error: File not able to be opened. Likely to be the wrong path")
	}
	defer file.Close()

	tokenChannel = make(chan Token)

	// Convert File into string
	code := ""
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		code += scanner.Text() + "\n"
	}

	// Convert String into string array
	tokens := []string{"\n"}
	buffer := ""
	isString := false
	for _, c := range code {
		switch {
		case c == '"':
			if !isString {
				isString = true
			} else {
				isString = false
				tokens = append(tokens, buffer)
				buffer = ""
			}
		case isString:
			buffer = buffer + string(c)
		case isSymbol(c):
			if buffer != "" {
				tokens = append(tokens, buffer)
				buffer = ""
			}
			tokens = append(tokens, string(c))
			continue
		case unicode.IsSpace(c):
			if buffer != "" {
				tokens = append(tokens, buffer)
			}
			buffer = ""
		default:
			buffer = buffer + string(c)
		}
	}

	// Determine Identifier
	lineNumber := 1
	for _, token := range tokens {
		identifier := ""
		var tokenVal any		
		
		if isDigit(token[0]) {
			identifier = "INTEGER_LITERAL"
			tokenVal = token
			continue
		}
		
		// Check for the different symbols
		switch token {
		case "\n":
			identifier = "LINE "
			tokenVal = strconv.Itoa(lineNumber)
			lineNumber++
		case "public":
			identifier = "PUBLIC_SYMBOL"
		case "class":
			identifier = "CLASS_SYMBOL"
		case "void":
			identifier = "VOID_SYMBOL"
		case "static":
			identifier = "STATIC_SYMBOL"
		case "int":
			identifier = "INT_SYMBOL"
		case ".":
			identifier = "DOT"
		case "=":
			identifier = "EQUALS"
		case ">":
			identifier = "BIGGER_THAN"
		case "<":
			identifier = "SMALLER_THAN"
		case "+":
			identifier = "PLUS"
		case "*":
			identifier = "STAR"
		case ";":
			identifier = "SEMICOLON"
		case "{":
			identifier = "LEFT_BRACE"
		case "}":
			identifier = "RIGHT_BRACE"
		case "(":
			identifier = "LEFT_PARENTHESIS"
		case ")":
			identifier = "RIGHT_PARENTHESIS"
		case "[":
			identifier = "LEFT_BRACKET"
		case "]":
			identifier = "RIGHT_BRACKET"
		default:
			identifier = "IDENTIFIER"
			tokenVal = token
		}

		// Make return token and add to channel
		returnToken := new(Token)
		returnToken.identifier = identifier
		returnToken.value = tokenVal
		tokenChannel <- *returnToken 
	}
}

func isSymbol(r rune) bool {
	symbols := []rune{'\n', ';', '.', '-', '+', '*', '>', '<', '=', '{', '}', '(', ')', '[', ']'}
	for _, symbol := range symbols {
		if r == symbol {
			return true
		}
	}
	return false
}

func isDigit(b byte) bool {
	if b-48 >= 0 && b-48 <= 9 {
		return true
	}
	return false
}
