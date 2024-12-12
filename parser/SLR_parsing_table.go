package parser

import (
	"errors"
	"fmt"
)

type SLR_parsing_Table struct {
	actionTable map[int]map[string]*Action
	gotoToTable map[int]map[string]*GoTo
}

type Action struct {
	// Accept / Shift / Reduce
	actionType string
	value      int
}

type GoTo struct {
	val int
}

func (automata *SLR_automata) CreateSLRTable(grammar *Grammar) *SLR_parsing_Table {
	table := makeSlrParsingTable()

	for _, state := range automata.states {
		for _, itemrule := range state.rules {
			var afterdot string
			if itemrule.dot < len(itemrule.rule.production) {
				afterdot = itemrule.rule.production[itemrule.dot]
			} else {
				// The dot is at the end of the production
				for _, terminal := range grammar.follow[itemrule.rule.nonTerminal] {
					if terminal == "$" {
						table.AddAction(state.id, "$", "Accept", 0)
					} else {
						table.AddAction(state.id, terminal, "Reduce", state.transitions[itemrule.rule.nonTerminal])
					}
				}
			}
			switch {
			case contains(grammar.nonTerminals, afterdot) != -1:
				// The dot is before a non terminal
				// Goto from the current state with the non terminal into the state consuming the current non terminal
				table.AddGoTo(state.id, itemrule.rule.nonTerminal, state.transitions[afterdot])
			case contains(grammar.terminals, afterdot) != -1:
				// The dot is before a terminal
				table.AddAction(state.id, afterdot, "Shift", state.transitions[afterdot])
			}
		}
	}
	return table
}

func makeSlrParsingTable() *SLR_parsing_Table {
	newTable := new(SLR_parsing_Table)
	newTable.actionTable = make(map[int]map[string]*Action)
	newTable.gotoToTable = make(map[int]map[string]*GoTo)
	return newTable
}

func (table *SLR_parsing_Table) GetAction(state int, symbol string) (Action, error) {
	if table.actionTable[state][symbol] != nil {
		return *table.actionTable[state][symbol], nil
	}
	return Action{}, errors.New("Could not get next Action")
}

func (table *SLR_parsing_Table) GetGoto(state int, symbol string) (GoTo, error) {
	if table.gotoToTable[state][symbol] != nil {
		return *table.gotoToTable[state][symbol], nil
	}
	return GoTo{}, errors.New("Could not get next Goto")
}

func MakeAction(ty string, value int) Action {
	newAction := new(Action)
	newAction.actionType = ty
	newAction.value = value
	return *newAction
}

func (table *SLR_parsing_Table) AddAction(state int, terminal string, actionType string, ActionValue int) {
	newAction := MakeAction(actionType, ActionValue)
	if table.actionTable[state] == nil {
		table.actionTable[state] = make(map[string]*Action)
	}
	if table.actionTable[state][terminal] != nil {
		panic("Grammar does not seem to be SLR Parsable, Action Table Error")
	}
	table.actionTable[state][terminal] = &newAction
}

func MakeGoto(val int) *GoTo {
	newGoto := new(GoTo)
	newGoto.val = val
	return newGoto
}

func (table *SLR_parsing_Table) AddGoTo(state int, symbol string, newstate int) {
	if table.gotoToTable[state] == nil {
		table.gotoToTable[state] = make(map[string]*GoTo)
	}
	if table.gotoToTable[state][symbol] != nil {
		// TODO Panic Here
		fmt.Println("Grammar does not seem to be SLR Parsable, GoTo Table error")
	}
	table.gotoToTable[state][symbol] = MakeGoto(newstate)
}

func (table *SLR_parsing_Table) PrintTable() {
	fmt.Println("Action: ")
	for i, m := range table.actionTable {
		fmt.Print(i)
		for str, action := range m {
			if str == "$" {
				fmt.Print(str + " ")
				fmt.Print(action.actionType + " ")
				fmt.Print(action.value)
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("GoTo: ")
	for i, m := range table.gotoToTable {
		fmt.Print(i)
		for str, state := range m {
			fmt.Print(str + " ")
			fmt.Print(state)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (table SLR_parsing_Table) getNextExpectedTokens(state int) string {
	retString := ""
	for i := range table.actionTable[state] {
		retString += " " + i
	}
	for i := range table.gotoToTable[state] {
		retString += " " + i
	}
	return retString
}
