package lexer

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"unicode"
)

type Token struct {
	Identifier string
	Value      any
}

func GetNext(tokenChannel chan Token) (*Token, error) {
	// Wait until the channel with tokens has a value inside
	select {
	case newToken, ok := <-tokenChannel:
		// If channel is closed -> File is done
		if !ok {
			return nil, errors.New("Lexer Error: File Ended")
		}
		return &newToken, nil
	}
}

func Lex(path string, tokenChannel chan Token) {
	// Open File
	file, err := os.Open(path)
	if err != nil {
		panic("Lexer Error: File not able to be opened. Likely to be the wrong path. Path given: " + path)
	}
	defer file.Close()

	// Scan over the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lineNumber := 1

	for scanner.Scan() {
		// Save the line into "line"
		line := ""
		line += scanner.Text()

		// Split String, removing whitespace etc.
		tokens := []string{"\n"}
		buffer := ""
		isString := false
		isSingleLineComment := false
		isMultiLineComment := false
		for _, c := range line {
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
			
			
			// Attach symbol as is to the tokens, except if it can be 
			// concatonated with the symbol before -> check for comment begin
			case isSymbol(c):
				if buffer != "" {
					tmparr := []rune(buffer)[0]
					
					// The string in the buffer is not a symbol
					if !isSymbol(tmparr){
						tokens = append(tokens, buffer)
						buffer = ""
					}else{
						// Is symbol -> Can be concatonated to // /* etc.
						concSymbol := concatonateSymbols(tmparr, c)
						// Symbols cannot be concatonated
						if concSymbol == ""{
							tokens = append(tokens, buffer)
							buffer = ""
						}else{
							// Symbols can be concatonated -> check for if comment
							if concSymbol == "//"{
								isSingleLineComment = true
							}
							if concSymbol == "/*"{
								isMultiLineComment = true
							}
						}
					}
				}
				tokens = append(tokens, string(c))
			
			
			
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
		for _, token := range tokens {
			identifier := ""
			var tokenVal any

			tmpdigit, intConvErr := strconv.Atoi(token)
			
			// If could be converted to int
			if intConvErr == nil {
				identifier = "INTEGER_LITERAL"
				tokenVal = tmpdigit
				continue
			}

			tmpbool, boolConvErr := strconv.ParseBool(token)
			// If could be converted to bool
			if boolConvErr == nil {
				identifier = "INTEGER_LITERAL"
				tokenVal = tmpbool
				continue
			}

			// Check for the different symbols
			switch token {
			case "\n":
				identifier = "LINE "
				tokenVal = strconv.Itoa(lineNumber)
				lineNumber++
			case "namespace":
				identifier = "NAMESPACE"
			case "using":
				identifier = "USING"
			case "class":
				identifier = "CLASS"
			case "void":
				identifier = "VOID"
			case "static":
				identifier = "STATIC"
			case "Main":
				identifier = "MAIN"
			case "int":
				identifier = "INT"
			case "bool":
				identifier = "BOOL"
			case "string":
				identifier = "STRING"
			case "double":
				identifier = "DOUBLE"
			case "if":
				identifier = "IF"
			case "else":
				identifier = "ELSE"
			case "while":
				identifier = "WHILE"
			case "return":
				identifier = "RETURN"
			case ".":
				identifier = "DOT"
			case ",":
				identifier = "COMMA"
			case "|":
				identifier = "OR"
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
			isSingleLineComment = false

			// Make return token and add to channel
			returnToken := new(Token)
			returnToken.Identifier = identifier
			returnToken.Value = tokenVal
			tokenChannel <- *returnToken
		}
	}
	close(tokenChannel)
}

func isSymbol(r rune) bool {
	symbols := []rune{'\n', ';', '.', '-', '+', '*', '>', '<', '=', '{', '}', '(', ')', '[', ']', '|', ',', '/'}
	for _, symbol := range symbols {
		if r == symbol {
			return true
		}
	}
	return false
}

func concatonateSymbols(s1 rune, s2 rune) string{
	if s2 == '='{
		if s1 == '>' ||s1 == '<' || s1 == '!' || s1 == '='{
			return string(s1) + string(s2)
		}
	}
	if s1 == '/'{
		if s2 == '/' || s2 == '*'{
			return string(s1) + string(s2)
		}
	}
	if s1 == '*' && s2 == '/'{
		return "*/"
	}
	if s1 == '|' && s2 == '|'{
		return "||"
	}
	if s1 == '&' && s2 == '&'{
		return "&&"
	}
	return ""
}