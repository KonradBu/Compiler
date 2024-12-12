package parser

func defGrammar(test bool) []Rule{
	grammar := []Rule{}
	if test{
		grammar = append(grammar, testGrammar()...)
	}else{
		grammar = append(grammar, compilerGrammar()...)
	}
	return grammar
}

func testGrammar() []Rule{
	rules := []Rule{}
	rules = append(rules, MakeRule("E", []string{"E", "+", "T"}))
	rules = append(rules, MakeRule("E", []string{"T"}))
	rules = append(rules, MakeRule("T", []string{"T", "*", "F"}))
	rules = append(rules, MakeRule("T", []string{"F"}))
	rules = append(rules, MakeRule("F", []string{"(", "E", ")"}))
	rules = append(rules, MakeRule("F", []string{"id"}))
	return rules
}

func compilerGrammar() []Rule{
	return []Rule{}
}