package parser

func defGrammar(test bool) []Rule {
	grammar := []Rule{}
	if test {
		grammar = append(grammar, testGrammar()...)
	} else {
		grammar = append(grammar, compilerGrammar()...)
	}
	return grammar
}

func testGrammar() []Rule {
	rules := []Rule{

		// NAME, LITERALS
		// INPUT BLOCK; ARGBLOCK

		//
		MakeRule("TYPE", []string{"double"}),
		MakeRule("TYPE", []string{"int"}),
		MakeRule("TYPE", []string{"string"}),
		MakeRule("TYPE", []string{"bool"}),
		//
		MakeRule("RETURNTYPE", []string{"void"}),
		MakeRule("RETURNTYPE", []string{"TYPE"}),
		//
		MakeRule("LITERAL", []string{"stringliteral"}),
		MakeRule("LITERAL", []string{"boolliteral"}),
		MakeRule("LITERAL", []string{"NUMLITERAL"}),
		MakeRule("NUMLITERAL", []string{"intliteral"}),
		MakeRule("NUMLITERAL", []string{"intliteral", ".", "intliteral"}),

		MakeRule("START", []string{"USINGBLOCK"}),
		MakeRule("USINGBLOCK", []string{"using", "name", ";", "USINGBLOCK"}),
		MakeRule("USINGBLOCK", []string{"NAMESPACE"}),
		MakeRule("NAMESPACE", []string{"namespace", "name", "{", "CLASS", "}"}),
		MakeRule("CLASS", []string{"class", "name", "{", "FUNCBLOCK"}),
		MakeRule("FUNCBLOCK", []string{"FUNC", "FUNCBLOCK"}),
		MakeRule("FUNCBLOCK", []string{"}"}),

		MakeRule("FUNC", []string{"static", "RETURNTYPE", "name", "(", "INPUTBLOCK", "{", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"}"}),
		MakeRule("STATEMENTBLOCK", []string{"FUNCCALL", ";", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"RETURN", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"VARASSIGN", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"VARIABLEDECLARATION", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"IF", "STATEMENTBLOCK"}),
		MakeRule("STATEMENTBLOCK", []string{"WHILE", "STATEMENTBLOCK"}),

		MakeRule("INPUTBLOCK", []string{")"}),
		MakeRule("INPUTBLOCK", []string{"string", "[", "]", "name", ")"}),
		MakeRule("INPUTBLOCK", []string{"INPUTSTART"}),
		MakeRule("INPUTSTART", []string{"TYPE", "name", "INPUTCONTINUED"}),
		MakeRule("INPUTCONTINUED", []string{",", "TYPE", "name"}),
		MakeRule("INPUTCONTINUED", []string{")"}),

		MakeRule("ARGBLOCK", []string{")"}),
		MakeRule("ARGBLOCK", []string{"ARGSSTART"}),
		MakeRule("ARGSSTART", []string{"EXPRESSION", "ARGCONTINUED"}),
		MakeRule("ARGCONTINUED", []string{")"}),
		MakeRule("ARGCONTINUED", []string{",", "EXPRESSION", "ARGCONTINUED"}),

		MakeRule("FUNCCALL", []string{"name", "(", "ARGBLOCK"}),
		MakeRule("FUNCCALL", []string{"name", ".", "name", "(", "ARGBLOCK"}),
		MakeRule("RETURN", []string{"return", "EXPRESSION", ";"}),
		MakeRule("RETURN", []string{"return", ";"}),
		MakeRule("VARIABLEDECLARATION", []string{"EMPTYVARIABLEDECLARATION"}),
		MakeRule("VARIABLEDECLARATION", []string{"SETVARIABLEDECLARATION"}),
		MakeRule("EMPTYVARIABLEDECLARATION", []string{"TYPE", "name", ";"}),
		MakeRule("SETVARIABLEDECLARATION", []string{"TYPE", "name", "=", "EXPRESSION", ";"}),
		MakeRule("VARASSIGN", []string{"name", "=", "EXPRESSION", ";"}),

		MakeRule("IF", []string{"if", "(", "EXPRESSION", ")", "{", "STATEMENTBLOCK"}),
		MakeRule("IF", []string{"if", "(", "EXPRESSION", ")", "{", "STATEMENTBLOCK", "ELSE"}),
		MakeRule("ELSE", []string{"else", "{", "STATEMENTBLOCK"}),
		MakeRule("WHILE", []string{"while", "(", "EXPRESSION", ")", "{", "STATEMENTBLOCK"}),

		MakeRule("EXPRESSION", []string{"EXPRESSION", "booloperator", "TERM"}),
		MakeRule("EXPRESSION", []string{"TERM"}),
		MakeRule("TERM", []string{"TERM", "plusoperator", "FACTOR"}),
		MakeRule("TERM", []string{"FACTOR"}),
		MakeRule("FACTOR", []string{"FACTOR", "timesoperator", "PRIMARY"}),
		MakeRule("FACTOR", []string{"PRIMARY"}),
		MakeRule("PRIMARY", []string{"FUNCCALL"}),
		MakeRule("PRIMARY", []string{"LITERAL"}),
		MakeRule("PRIMARY", []string{"name"}),
		MakeRule("PRIMARY", []string{"(", "EXPRESSION", ")"}),
	}
	return rules
}

func compilerGrammar() []Rule {
	return []Rule{}
}
