package Tokenizer

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Tokenize(path string) string {
	// Open File
	file, err := os.Open(path)
	if err != nil {
		panic("file not able to be read")
	}
	defer file.Close()

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
	returnCode := ""
	lineNumber := 1
	for _, token := range tokens {
		identifier := ""
		if isDigit(token[0]) {
			returnCode += "INTEGER_LITERAL: " + token + "\n"
			continue
		}
		// Check for the different symbols
		switch token {
		case "\n":
			identifier = "LINE " + strconv.Itoa(lineNumber)
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
			identifier = "IDENTIFIER: " + token
		}
		returnCode += identifier + "\n"
	}
	return returnCode
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
