package parser

// Value 1: Non terminal; Value 2: terminal
func parseLL1Table(table *parsetable, input []string) bool {
	for i := 0; i < len(input); i += 2 {
		nonterminal := input[i]
		terminal := input[i+1]

		// Index of the values
		ntVal := table.nonterminals[nonterminal]
		tVal := table.terminals[terminal]

		if table.table[ntVal][tVal] != "" {
			return true
		}
		return false
	}
}
