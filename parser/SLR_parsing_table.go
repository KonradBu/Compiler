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

func makeSlrParsingTable() *SLR_parsing_Table{
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
	if table.gotoToTable[item][symbol] != nil{
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
	if table.actionTable[state][terminal] != nil {
		panic("Grammar does not seem to be SLR Parsable")
	}
	table.actionTable[state][terminal] = &newAction
}

func MakeGoto(val int) *GoTo{
	newGoto := new(GoTo)
	newGoto.val = val
	return newGoto
}

func (table *SLR_parsing_Table) AddGoTo(state int, symbol string, newstate int) {
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

func (table SLR_parsing_Table) getNextExpectedTokens(item int) string{
	retString := ""
	for i, _ := range table.actionTable[item] {
		retString += " " + i
	}
	for i, _ := range table.gotoToTable[item]{
		retString += " " + i
	}
	return retString
}