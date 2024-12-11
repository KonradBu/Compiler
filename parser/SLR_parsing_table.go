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

	for _, item := range automata.items {
		for _, rule := range item.rules {
			dot := item.dots[&rule]
			switch {
			case contains(grammar.nonTerminals, rule.production[dot]) != -1:
				// The dot is before a non terminal
				// Goto from the current item with the non terminal into the item consuming the current non terminal
				table.AddGoTo(item.id, rule.nonTerminal, item.transitions[rule.production[dot]].id)
			case contains(grammar.terminals, rule.production[dot]) != -1:
				// The dot is before a terminal
				next := rule.production[dot]
				table.AddAction(item.id, rule.production[dot], "Shift", item.transitions[next].id)
			case len(rule.production) == dot:
				// The dot is at the end of the production
				for _, terminal := range grammar.follow[rule.nonTerminal] {
					if terminal == "$" {
						table.AddAction(item.id, "$", "Accept", 0)
					} else {
						table.AddAction(item.id, terminal, "Reduce", item.transitions[rule.nonTerminal].id)
					}
				}
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

func (table *SLR_parsing_Table) GetAction(item int, symbol string) (Action, error) {
	if table.actionTable[item][symbol] != nil {
		return *table.actionTable[item][symbol], nil
	}
	return Action{}, errors.New("Could not get next Action")
}

func (table *SLR_parsing_Table) GetGoto(item int, symbol string) (GoTo, error) {
	if table.gotoToTable[item][symbol] != nil {
		return *table.gotoToTable[item][symbol], nil
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
	if table.actionTable[state] == nil{
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
	if table.gotoToTable[state][symbol] != nil{
		panic("Grammar does not seem to be SLR Parsable, GoTo Table error")
	}
	table.gotoToTable[state][symbol] = MakeGoto(newstate)
}

func (table *SLR_parsing_Table) PrintTable() {
	fmt.Println("Action: ")
	for i, m := range table.actionTable {
		fmt.Print(i)
		for str, action := range m {
			fmt.Print(str + " ")
			fmt.Print(action.actionType + " ")
			fmt.Print(action.value)
			fmt.Print(" ")
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

func (table SLR_parsing_Table) getNextExpectedTokens(item int) string {
	retString := ""
	for i := range table.actionTable[item] {
		retString += " " + i
	}
	for i := range table.gotoToTable[item] {
		retString += " " + i
	}
	return retString
}
